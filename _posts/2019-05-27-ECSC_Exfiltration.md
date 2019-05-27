ECSC Network forensics exfiltration

>Our SoC has detected a confidential document being exfiltrated! Fortunately the employed method is not very advanced a timely network capture was made.
>Retrieve the document

We get a PCAP capture with 27k packets
Not expectign anythign too advanced we see some suspicious HTTP traffic.
POST requests with "data" and "uuid" parameters and the server is responding with 418 - I'm a teapot
If we apply Wireshark to filter only for forms with urlencoded POST requests, we get the whole exchange between victim and C&C server.

We can then export the packets into a text file, remove the uuid bit, and have the full data.

The data however did not seem to translate well from hex to any readable format so lets look for more clues.

The best place to look for clues is before the exchange happened, as we don't know what the encryption of the file is.

