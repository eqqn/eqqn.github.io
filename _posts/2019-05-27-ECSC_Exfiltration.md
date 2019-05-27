# ECSC France - Exfiltration Write-up
Challenge text:

>Our SoC has detected a confidential document being exfiltrated! Fortunately the employed method is not very advanced a timely network capture was made.
>Retrieve the document

We get a PCAP capture with 27k packets.
Not expecting anything too advanced we look for some suspicious HTTP traffic.
![Image](https://eqqn.github.io/images/packets.JPG)

![Image](https://eqqn.github.io/images/mal_proof.JPG)

Bingo, we are on the right track.

See POST requests with "data" and "uuid" parameters and the server is responding with [418 - I'm a teapot](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/418)

![Image](https://eqqn.github.io/images/exfil_payload.JPG)

If we apply Wireshark to filter only for forms with urlencoded POST requests, we get the whole exchange between victim and C&C server.

![Image](https://eqqn.github.io/images/exfil_filter.jpg)

We can then export the packets into a text file, remove the uuid bit, and have the full data.

The data however did not seem to translate well from hex to any readable format so lets look for more clues.

The best place to look for clues is before the exchange happened.

Filtering by the malware server IP address we see some ICMP (ping) requests. 

![Image](https://eqqn.github.io/images/exfil_icmp.JPG)

In a primitive way to evade our investigation, it was sent over different protocol. 

The ICMP packets transmit configuration for the malware: 

![Image](https://eqqn.github.io/images/exfil_config.JPG)
```
config : exfiltered_file_size=4193bytes
config : data_len_for_each_packet=random
config : file_type=DOCX
```

So now we know the "encryption" algorithm(XOR), file size and format.

File format can be useful for identifying the "Magic bytes" in document types, i.e DOCX, PNG and other file formats have fixed hexadecimal sequences at the start and the end of the file.

We first take and convert the cleaned hex payload (strip all spaces, newlines) and convert it to binary file using *xxd* 

`xxd -r -p input.hex output.bin`

```[sec@sec-pc ecsc]$ ls -l output.bin
-rw-r--r-- 1 sec sec 4193 27 mai   14:55 output.bin
```
The filesize is good.

Since we know it is encrypted with XOR we can bruteforce the key to get the file.

XORtool is popular among CTF players to do just that.

![Image](https://eqqn.github.io/images/xortool.JPG)

`python2 xortool/xortool output.bin -b` will bruteforce the key. ( the selection of right key is based on how many valid bytes for binary or plaintext there are after XOR)

We can see that "ecsc" is a possible key candidate. The script outputs all the combinations to a folder. The CSV file helps us locate which key belongs to which file, in this case , "000.out" is detected as archive, but upon opening we see [Content_Types].xml inside, which is typical document behaviour when opened as archive. We rename the extension to ".docx" and voilla!

![Image](https://eqqn.github.io/images/exfil_solution.JPG)

ECSC{v3ry_n01sy_3xf1ltr4t10n}





