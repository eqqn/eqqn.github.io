# European Cyber Security Challenge Incident 3615 Forensics

The annual European Cyber Security Challenge holds a competition for selection of candidates.

Thanks to my friends, I learned about the French competition and was motivated to do some challenges.

If you want to know how to use volatility to do memory dump forensics with a dash of ransomware, keep reading.


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



 ( each proccess has a PID-proccess ID, and PPID - parent proccess ID 

If you note the proccess tree, assistance.exe spawned a connhost.exe

To see the network activity inside the memdump, use netscan command : 

`volatility -f mem.dmp --profile Win10x64_14393 netscan`

> 0xe0001265ad10     TCPv4    192.168.248.133:49774(local add)          192.168.1.25:8080(remote add)    ESTABLISHED      5208     assistance.exe 2019-05-08 20:00:17 UTC+0000

We see assistance.exe calling to 192.168.1.25:8080. 

Inside the 5208 proccess dump we see 

![Image](https://eqqn.github.io/images/ransom_endpoint.jpg)



Part 3?



I was using this blog series to learn more about volatility
https://www.andreafortuna.org/2017/06/25/volatility-my-own-cheatsheet-part-1-image-identification/
