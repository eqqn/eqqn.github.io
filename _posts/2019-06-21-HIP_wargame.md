# HackInParis 2019 Wargame write-up

### intro

I didn't attend HIP(only budgeted for LeHack :p), but they were kind to host a publicly accessible wargame. 
It was a good opportunity to evaluate Sysdream "Malice" CTF platform of on-demand CTF's for security training or other needs.

## Meet your doctor 1

We are tasked to retrieve the site administrators password. We have a website without any login forms or other vectors for injection.

However, there is a link "Try our API!", which gives us a nice [GraphQL playground](https://github.com/prisma/graphql-playground), 
which allows us to see the API schema and construct the queries inside the browser. 

![Image](https://eqqn.github.io/images/Doctor3.JPG)

We get a list of doctors and their passwords by constructing this query:
```js
{
    doctors{
    firstName, 
    lastName, 
    password}
}
```

And receive a response of many doctors, plus their passwords

```js
      {
        "firstName": "Admin",
        "lastName": "Admin",
        "password": "Now-_Let$|GetSeri0us"
      }
```

## Defcon 5
![Image](https://eqqn.github.io/images/defcon1.JPG)

In this challenge you have to go through levels 5 to 1 analyzing javascript code of a dashboard similar to NEXT GEN -1 DAYS SECURITY vendors.

The first password is hiding in sourcecode at *static/js/script.js*  in clear: **STRONGPASSWORD**
![Image](https://eqqn.github.io/images/defcon2.JPG)

## Defcon 4
Level 4 checking script is in the same file and looks like this:

```js
function check4(pwd) {
  var input = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
  var output = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm";
  var index = x => input.indexOf(x);
  var translate = x => (index(x) > -1 ? output[index(x)] : x);
  return (
    pwd
      .split("")
      .map(translate)
      .join("") === "ebgEBGeBgEbgEBGEBgEBGEBGEBGebgeBgEbGebg"
  );
}
```

We can see that var input and output are related - it is just a shift cipher - Caesar or ROT13 as commonly known. If we input the string it compares against back, it should be shifted the same amount - returning the plaintext

If we put "ebgEBGeBgEbgEBGEBgEBGEBGEBGebgeBgEbGebg" as the argument of check4() 
![Image](https://eqqn.github.io/images/defcon3.JPG)
We get "rotROTrOtRotROTROtROTROTROTrotrOtRoTrot"  which validates the challenge!


#### Outro
Although I only solved the 3 most popular (read easy) challenges, I had exciting time learning about graphQL, introspection, running my own playground server, retrieving schemas and trying different approaches I found online. I spent a lot of time trying "RootMyNeighbour" seeing if i can execute CVE-2018-1111 (because the only traffic I could sniff as a server was DHCP  requests, but I was misdirected, and french locale keyboard on a VM is worse than stepping on a lego) I hope to be able to learn from others' writeups on MeetMyDoctor #2 & #3.

