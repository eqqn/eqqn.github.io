# Hackrocks Cyber Summer Camp 2022

### intro
It's the summer, and there are many conferences and CTF's going on. This summer I was moving, changing jobs and had to 
miss out on the French side of scene, but managed to poke at couple CTF's for fun. Summer cyber camp was good, since over the month
there was plenty of time to come back to challenges after busy periods are over.

## Arith, 30 points, Easy


> The service you are to study asks you to guess a total of 23 numbers, all less than 150. Now, as a member of the purple team, can you help the investigator beat this service?
> 
> The game can be found at:
> 
> challenges.REDACTED.com:4747

Easy way to interact with a server is *nc* - netcat command.

![Image](https://eqqn.github.io/images/asym_wrong_guess.JPG)

23 numbers guessing game over TCP?  Time to write a bot... 

```python
import socket
import time
import array as arr

# Create a TCP/IP socket
#sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# Connect the socket to the port where the server is listening
#server_address = ('challenges.REDACTED.com', 4747)
#print('connecting to {} port {}'.format(*server_address))
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server_address = ('challenges.REDACTED.com', 4747)
sock.connect(server_address)
sock.settimeout(1)


cci=8
counter=0
#lst= [103, 109, 98, 104, 124, 37, 86, 67, 54, 56, 50, 85, 118, 85, 52, 96, 99, 122, 96, 49, 111, 52, 126] #completed list

lst = [103,109,98,104,124,37,86,67,54] # incomplete list if you want to run it yourself

def reconnect():
    global sock
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_address = ('challenges.REDACTED.com', 4747)
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

I converted the characters to hex

`67 6D 62 68 7C 25 56 43 36 38 32 55 76 55 34 60 63 7A 60 31 6F 34 7E`

`flag{$UB571TuT3_by_0n3}`

## amAPT, 60 points, medium

> The pcap file is all you needed.
> 
> In addition to the information of pcap, the attacker may have implemented a C2. 

You have a pcap file  [link](https://github.com/eqqn/eqqn.github.io/blob/master/uploads/log.pcap) 





