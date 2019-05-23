# European Cyber Security Challenge Incident 3615 Forensics

The annual European Cyber Security Challenge holds a competition for selection of candidates.

Thanks to my friends, I learned about the French competition and was motivated to do some challenges.

If you want to know how to use volatility to do memory dump forensics with a dash of ransomware, keep reading. We will dive deep into ransomware infested memory and try to track the encryption key to save the day.


## Part 1
Challenge states ( translated from french) :

>A victim has fallen prey to ransomware. **Paying the ransom amount is not an option**. You are called to restore the encrypted data.

>A set of elements is required to advance with the incident response investigation.

>To begin, What is the ransomware executable name, proccess identifier and what happened to the file *flag.docx* once encrypted ? Give the SHA1 of the filename with extension.
Respond in this format ECSC{name_of_ransomware.exe:PiD:sha1}. You are provided with a [memory dump](https://www.google.comhttps://www.ecsc-teamfrance.fr/files/c2565a7d58ad2cb7cc6870f5caf4efe0/mem.dmp.ce117720fa4126f57814b3a779a7eb4ba21570e3f5dfd44a6706771783a46f1b.zip?token=eyJ0ZWFtX2lkIjpudWxsLCJ1c2VyX2lkIjo3NDMsImZpbGVfaWQiOjgwfQ.XOa0SA.A4ZiPl7AMvjuI8XRiCsu1MR5Vpk) . 

Volatility is a well known command line tool for memory analysis, and I am going to show you some basic uses.

For it to work, you need to load the memory dump with correct profile.
`volatility -f mem.dmp imageinfo` will give you a suggestion to which profile you should use, however you should find which of the recommended profiles works best. 
![Image](https://eqqn.github.io/images/volatility_profile.JPG)

Next up is identifying which proccess is malicious. There is no magic way to do this, but we know the that the proccess is ransomware so it will be doing some encryption. Calls to crypto libraries and dll's can give it away. We also know that the file *flag.docx* was encrypted, so lets look for it.

`volatility volatility -f mem.dmp --profile=Win10x64_14393 pslist` or `pstree` will show running proccesses. 

To dump memory of a certain proccess use *memdump*. After dumping a few proccesses that looked unusual I finally found the ransomware.

`volatility -f mem.dmp --profile=Win10x64_10563 memdump -p 5208 -D` where 5208 is the proccess ID of "assistance.exe" ( sounds helpfully malicious now doesn't it?) 

Looking for *flag.docx*  helps you find what happened to the filename
![Image](https://eqqn.github.io/images/renaming_flag.JPG)

and what happened to the rest of the files...

![Image](https://eqqn.github.io/images/encrypting.jpg)
If we look for *ZmxhZy5kb2N4* (new name of flag.docx) we find another reference to it with a french sounding "chiffré" extension on top of it.  SHA1 of the filename is the last piece of the puzzle for the first flag.


## Part 2
Part 2 of the challenge tasks you to recover the encryption key, which is the next flag.

To see the network activity inside the memdump, use netscan command : 

`volatility -f mem.dmp --profile Win10x64_14393 netscan`

`0xe0001265ad10     TCPv4    192.168.248.133:49774(local add)          192.168.1.25:8080(remote add)    ESTABLISHED      5208     assistance.exe 2019-05-08 20:00:17 UTC+0000`

We see assistance.exe calling to 192.168.1.25:8080. 

Inside the 5208 proccess dump we see 

![Image](https://eqqn.github.io/images/ransom_endpoint.jpg)

No luck decoding the payload, so lets try looking elsewhere. But we can divine that the ransomware is sending the encryption keys out to the C&C server.

I stumbled upon the ransom note somewhere in the dump:

![Image](https://eqqn.github.io/images/ransom_note.JPG)
The identifier is a clue to finding the request sent to C&C server.

And there is another version with it's original source code! It will be **very** useful later on.

![Image](https://eqqn.github.io/images/ransom_note_source.JPG)

Looking for the identifier leads to the add keys api call in cleartext: 

![Image](https://eqqn.github.io/images/enckey.JPG)

This is the part that has tripped up some people. The key isn't clearly evident and the enckey string comes in a format as such:
>cd18c00bb476764220d05121867d62de64e0821c53c7d161099be2188b6cac24cd18c00bb476764220d05121867d62de64e0821c53c7d161099be2188b6cac2495511870061fb3a2899aa6b2dc9838aa422d81e7e1c2aa46aa51405c13fed15b95511870061fb3a2899aa6b2dc9838aa422d81e7e1c2aa46aa51405c13fed15b

Upon closer inspection you can notice that id is repeated across the string. If we look into original source of the ransomware at https://github.com/mauri870/ransomware/ and look for encryption key or decryption functions, you will see something similar 
```
		// Generate the id and encryption key
		keys["id"], _ = utils.GenerateRandomANString(32)
		keys["enckey"], _ = utils.GenerateRandomANString(32)
```

SO lets raise a hypothesis that key length is like in original, 32. If we remove repeating parts and user id from the string, we get 3 distinct 32char strings : 

- 64e0821c53c7d161099be2188b6cac24
- 95511870061fb3a2899aa6b2dc9838aa
- 422d81e7e1c2aa46aa51405c13fed15b

The 2nd one is the encryption key and validates the challenge.

## Part 3

The final task is to decrypt a file (flag.docx) given to us, that was encrypted by the ransomware.
As the ransom note and code snippets state, it is encrypted in AES256-CTR mode. Analyzing the code we learn that the first AES block is used as Initialization Vector (IV), or a salt. Without writing too much code, I found a way to plug-in the file and key into the golang test files.

The final result : 

![Image](https://eqqn.github.io/images/decrypted.JPG)

ECSC{M4udit3_C4mp4gn3_2_r4NC0nG1c13L} "The cursed ransomware campaign "


### References
I was using this blog series to learn more about volatility
https://www.andreafortuna.org/2017/06/25/volatility-my-own-cheatsheet-part-1-image-identification/


