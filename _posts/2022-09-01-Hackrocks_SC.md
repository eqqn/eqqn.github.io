# Hackrocks Cyber Summer Camp 2022

### intro
It's the summer, and there are many conferences and CTF's going on. This summer I was moving, changing jobs and had to 
miss out on the French side of scene, but managed to poke at couple online CTF's for fun. Summer cyber camp was good, since over the month
there was plenty of time to come back to challenges after busy periods are over.

## Arith, 30 points, Easy


> The service you are to study asks you to guess a total of 23 numbers, all less than 150. Now, as a member of the purple team, can you help the investigator beat this service?
> 
> The game can be found at:
> 
> challenges.hackrocks.com:4747

Easy way to interact with a server is *nc* - netcat command.

![Image](https://eqqn.github.io/images/asym_wrong_guess.JPG)

23 numbers guessing game over TCP?  Time to write a bot... 

```python
import socket
import time
import array as arr

# Connect the socket to the port where the server is listening
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server_address = ('challenges.hackrocks.com', 4747)
sock.connect(server_address)
sock.settimeout(1)

cci=8  # current correct array index
counter=0
#lst= [103, 109, 98, 104, 124, 37, 86, 67, 54, 56, 50, 85, 118, 85, 52, 96, 99, 122, 96, 49, 111, 52, 126] #completed list

lst = [103,109,98,104,124,37,86,67,54] # incomplete list if you want to run it yourself

def reconnect():
    global sock
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_address = ('challenges.hackrocks.com', 4747)
    sock.connect(server_address)
    time.sleep(0.2)
    return sock

def send(number,sock):
    number=f"{number}\n"
    #print(number)
    sock.send(bytes(number, 'utf-8'))
    time.sleep(0.15)
    char = sock.recv(100)
    #print(char)

def guess(number,sock):
    cleannum=number
    number=f"{number}\n"
    #print(number)
    sock.send(bytes(number, 'utf-8'))
    time.sleep(0.15)
    char = sock.recv(32)
    global cci
    global counter
    global lst
    if char != None and b'correct' in char:
        lst.append(cleannum) 
        counter=0
        print(lst)
        cci=cci+1 
        guess(counter,sock)
    if char != None and b'wrong' in char:
        sock.close()
        reconnect()
        #print(char)
        char=""
        counter=counter+1


while True and cci<23:
    for i in lst:
        send(i,sock)
    guess(counter, sock)
```

I wrote a bot to do some guessing. At first I populated the list manually, but after the refactor and even including recursion(!) 
I was able to have it run silently and error free... Except that it exits when all the numbers are submitted, so I changed some indexes here and there.  
Uncomment the print statements if you like.

![Image](https://eqqn.github.io/images/asym_populating_lst.JPG)

![Image](https://eqqn.github.io/images/asym_all_correct_flag_text.JPG)

Cheeky, but doesn't look too hard. We have a bunch of numbers, 37 to 126. That's well within ASCII range. After trying a few ciphers and encodings, 
I found the one. 
https://www.dcode.fr/ascii-shift-cipher

I converted the characters to hex before passing them through decoders.

`67 6D 62 68 7C 25 56 43 36 38 32 55 76 55 34 60 63 7A 60 31 6F 34 7E`

`flag{$UB571TuT3_by_0n3}`

## amAPT, 60 points, medium

> The pcap file is all you needed.
> 
> In addition to the information of pcap, the attacker may have implemented a C2. 

You have a pcap file  [link](https://github.com/eqqn/eqqn.github.io/blob/master/uploads/log.pcap) 

Looking at HTTP and TCP streams, it is is clear something was downloaded. Wireshark allows you to export HTTP objects (files) that are not encrypted, and we recover "*stashed.bin*". 

![Image](https://eqqn.github.io/images/ace-tcp.png)

Opening the file, it looks like a zip archive with some "zzzzzzzzz" padded before regular file contents which begin with PK. We also have a sneak peak that there is `stolen/flag.txt`

https://users.cs.jmu.edu/buchhofp/forensics/formats/pkzip.html

![Image](https://eqqn.github.io/images/amAPTstolentxt.png)

You can probably guess that we saw the password in the TCP session because it was sent over plaintext. We use it to decrypt the file.

`flag{c0rrupt3d_zip_4nd_extr4cted_m4lware}` 

I've flagged it under 20 minutes and it was a slight confidence boost.

## Screamshot, 65 points, medium

> The web itself was exposed but no suspicious nor malicious activities...
> But the VA result point us to many vulnerabilities.
> The challenge can be found at:  challenges.hackrocks.com:9998

This was by far the best chall for me. You have a website with open registration, and you can register as a user. 
However, the feature panel is for admin only. You are also given partial application source code https://github.com/eqqn/eqqn.github.io/tree/master/uploads/screamshot-participant .

Initially I tried to execute some SSTI strings on the username/email/bio fields, but it wasn't working. I tried to crack the JWT token secret with john-the-ripper and wordlists, as well as some brute force cracking, without much results.

I put the challenge on hold. I used one of the hints and it was telling that I should enumerate the application some more.

>  Clue nº 1
> All you need is basic to advance directory knocking to your web infrastructure like “/.git” (NO FUZZERS ALLOWED)

The "NO FUZZERS" statement didn't get across to me and I wasted a time and CTF infra running scans trying to find some hidden directory.

It is apparent that my enumeration methods are a bit dated and I need a better wordlist.

I also tried some suspected python files but nothing sticked related to config, dependencies, routing. After some conversation with another player, it was clear I was very close.
I retried with different ideas sometime later and found `http://challenges.hackrocks.com:9998/.env` file containing a secret JWT key.

`JWT_KEY = S3crEt_t0K3n$$$$`  . I used it to change my JWT token cookie and sign it on jwt.io  I also incremented the validity time to last me for the rest of CTF :)

![Image](https://eqqn.github.io/images/screamshot_admin_jwt.png)

#### admin panel

#### exploitation


## ACE-ng, 55 points, medium

> For the moment, our digital forensic expert has informed us that the malware is coded in Python 3.7.0. After that, he has mysteriously disappeared...
> It's up to you to continue his mission... Are you ready?

You are provided with a file https://github.com/eqqn/eqqn.github.io/blob/master/uploads/extracted_malware

#### Going off the wrong path

Not entirely sure how to interact with the file, I decided to run it[1] . I noticed some libraries being written and loaded when following it's debug outputs, and seeing the message output. However, the files were prompty removed. I inserted some breakpoints to grab the files before they are deleted. 

< screenshot for btc address > 

< pic of folder struct>

decompiling the pyc files with uncompyle6

While I understood the application a bit better, none of the libraries contained something I could latch on. I put the challenge on hold, took the hints.

I also ran the file as root

<screenshot>
    
There goes my VM... 

#### Correcting the course

After some time, I decided to research more into pytohn decompilation, which is in line with the hint. Maybe there is a way to decompile it differently?



triton malware mention. f-secure blog link? 
    
https://blog.f-secure.com/how-to-decompile-any-python-binary/
    
https://www.mandiant.com/resources/blog/attackers-deploy-new-ics-attack-framework-triton
    





[1] In retrospect that is not very smart.
