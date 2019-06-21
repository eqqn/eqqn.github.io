# HackInParis 2019 Wargame write-up

### intro

I didn't attend HIP(only budgeted for LeHack :p), but they were kind to host a publicly accessible wargame. 
It was a good opportunity to evaluate Sysdream "Malice" CTF platform of on-demand CTF's for security training or other needs.

## Meet your doctor 1

We are tasked to retrieve the site administrators password. We have a website without any login forms or other vectors for injection.

However, there is a link "Try our API!", which gives us a nice [GraphQL playground](https://github.com/prisma/graphql-playground), 
which allows us to see the API schema and construct the queries inside the browser. 
We get a list of doctors and their passwords by constructing this query:
``` 
{
    doctors{
    firstName, 
    lastName, 
    password}
}
```

And receive a response of many doctors, plus their passwords

```
      {
        "firstName": "Admin",
        "lastName": "Admin",
        "password": "Now-_Let$|GetSeri0us"
      }
```

## Defcon 5
