# vba01-baby
### The challenge: 
"Help, I've accidentally opened an email attachment, help me find out what happened" 

We are given a [XLS sheet](https://eqqn.github.io/uploads/vba01-baby_272038055eaa62ffe9042d38aff7b5bae1faa518.xls) with macro enabled.
![worskheet_content](https://eqqn.github.io/images/artificialbitcoin.jpg "Worskheet content")

Nothing suspicious in the Worksheet.

At this point it is wise to proceed in a Windows VM, because macros will run as soon as the document is opened, once they are enabled. Let's edit the macro:

### Macro
![macro script](https://eqqn.github.io/images/macro_content.jpg "Macro script")

Interestingly enough, the payload is generated from cells contents.
VBA language is used for macros and printing the output of the function is not as straightforward, but we can direct the payload content into a cell, to see what's happening. Looking online, you can find samples how it is done, if you are not familiar with VBA.
Adding `Cells(1, 1) = aaaaaaa(0) + aaaaaaa(1) + aaaaaaa(2)` after regwrite, we dump the string into the first cell.

**WScript.ShellHKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run\
flagINS{Do_n0t_Ena8le_M4cro}**

## Extra
* [Looking for an internship!](https://www.linkedin.com/in/marius-giedrius/) mariusgd@protonmail.com (shameless, I know )

* vba02-bitminer was a similar chall, except that the result of the payload was a "miner" executable. After disassembly it was difficult to see any magic strings, so I did the only reasonable thing - running it, and inspecting the packet sent on Wireshark. 
The recipient being bitminer.insomni.hack , which greeted requests not matching the format with "Go away hacker" text.
I was able to craft my own POST requests on the API with arbitrary parameters, and the error messages indicated it was an SQL DB, but with 15 minutes remaining I didn't go through with blind SQL injection. 

* My first Insomnihack was a success, I've completed 2 challenges(the other being backdoor) on my own right and the rest of the team contributed 1 flag. Being one of the first 15 to complete this chall, I enjoyed the limelight in such a big event on the scoreboard.

