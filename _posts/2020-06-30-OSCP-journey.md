![Image](https://eqqn.github.io/images/oscp-logo.png)

### Motivation for OSCP

Every since I learned about this certificate a couple years ago I would think of in very high regard. Something about 24 hour hacking exam that made it seem pretty hard and beyond my abilities at the time. This impression was made only stronger after reading peoples blogs on this topic. Their failures, successes, lessons learned. OffSec themselves advertise this image with their motto "Try harder". After all, the course materials for PWK are a given, but to earn the OSCP certification you will need to "Try harder".

A certificate that does not require yearly renewal, membership fee, re-certification is also a big plus in my book. Less time (and money) spent doing paperwork and more time for hacking, research, attending conferences, participating in CTF's, blogging. Mind you, OSCP is by no means simple. The rules and checklists are long, submission/examination procedures and guidelines will take many hours and revising to follow and memorize, but I reckon it is still simpler than than any of the alternatives.

### Certifications as a business model
I spent quite a while examining what certification I could start with, to enhance my cybersecurity credentials and expertise better. Some come with year and experience requirements in Cybersecurity,
Let's not kid ourselves and acknowledge that certifications is just another business model.
For example, we have a bunch of certifications that need renewal every 3 years. Between different training providers (I would rather call them vendors) they all have the same-ish recertification proccess : proving 120 credits of work.
Lets break down the list 
* 2 day seminar - 16 credits
* reading a book, blog, resource, magazine - 5 credits
* detect vulnerability in your organization - 10 credits (good luck with NDA's buddy, unless you publish a CVE which is much more work than 10 credits imo)
* hosting events - 1 credit/hr
* teaching cybersecurity - 3 credits/hr

and then they hit you with the best offer - a 60 credit "learning" material for a tiny annual price of 200$ a year , which is most likely of questionable quality and merit, just a box to tick. I've spent some 30 hours on various materials in pluralsight to teach myself certain things and let me tell you, diminishing returns and repetition sets in pretty quick when learning on superficial, surface level, without practical application of the subject matter.

Imagine the paperwork, looking for approval , documenting all the things you did and worrying that some of them may not get as much points as you want. Sounds a bit like OSCP exam experience, except they dont offer you to buy it away with money, you have to earn it. I find it worthwhile to pay it up front with work and not having to worry about it expring. I am not looking forward to going this path, but when it's all paid for and clocked in during work hours, I will have to do it as well. 

![there are easier certs out there... difference is only couple letters](https://eqqn.github.io/images/humor2.jpg)

there are easier certs out there... difference is only a couple of letters...
### Advocating for informed defence

Why put so much work into "pentesting certificate" if you are not a pentester? In software world the mantra is "Agile". It should be applied to security with more rigour too. If you work in security you will know it. The false positives, the "vulnerabilities", lack of understanding of threats. The bigger the company, the more policies. The more policies, the more "solutions" will be employed. The security management has to make policies to show progress. Some policies will be good, some not so much. It is valuable to understand such things to tell which one deserves your team effort , and which one is just a box to tick, worth only a fraction of effort, or not worth doing at all. A lot of waste is created, ironically which is good for the industry itself. Complexity costs. 

The more intimate you are with the technical aspects of hacking, the faster you can rule out irrelevant threats, have reasonable threat models. If you can whip up a proof-of-concept attack and demonstrate it to your team, it will have a much bigger impact and receive priority. If you know what an attack looks like, you can make a log metric. You can make actionable insights outside pentesting.


### Preparation
At the time of writing, I am a masters graduate in Cybersecurity of a well regarded French "Grande École", ENSIMAG, with 2 years of work experience in DevSecOps, compliance and testing. Suffice to say all my work revolves around new technologies and security, but not a whole lot of it is offensive. Most of it is done on my own time, or it's how I get my kicks during the slower days at work, by looking at the neighbours in our network or doing OSINT on our company.

Over past years I would crack a few challenges on Root-Me every now and then, also after purchasing Burpsuite I tried out their [Web Security Academy](https://portswigger.net/web-security) which does a good job at teaching you concepts such as command injection, SQL injection and many others. What I learned there, I tried to validate on root-me to reinforce my lessons, as the latter usually has more advanced and black-box type challenges that require not only the knowledge of the framework/technology but also a fair bit of trial and error and intuition.

I tried hack the box, mostly in the challenges category. I got user-level flag on one current machine, but I didn't know privesc well enough to carry through. I needed more background. 
Following some infosec twitter accounts I noticed "The Cyber Mentor" and his course [Practical Ethical Hacking](https://www.udemy.com/course/practical-ethical-hacking) . It is very condensed and amazing value for money. For 20-30$ you will learn so many things far above it's price point, such as Active directory. At the time, not even PWK had that in their curriculum, as it was only added in the 2020 PWK course.

It teaches you basic methodology and you can try it out on retired HackTheBox machines. Each machine also has their walkthrough which is very short and sweet, and have their lessons delivered timely. There are also Ippsec video walkthroughs, which I used to learn some very intricate things, but those videos are 3-4 times longer and includes things outside critical path to exploit a machine - i.e you will be sidetracked by watching him make custom wordlists for a bruteforce attack, when in the end it won't be necessary to get the root. You can learn many things from him, but it just isn't my cup of tea as I can only stay focused for so long.

You can also achieve similar results by reading the write-ups. I tried to solve all the machines myself before looking at video solution, but sometimes I would also check the write-ups whenever I got stuck and was being unproductive. 
All of TCM's course machines fall into TjNull's [OSCP like HTB machines](https://docs.google.com/spreadsheets/d/1dwSMIAPIam0PuRBkCiDI88pU3yzrqqHkDtBngUHNCw8/htmlview#), so you get a head start on the list by going through the course. Finally, inspired by the course, I rooted 3 of the current "easy" HTB machines ( those which dont have write-ups), although some forum crawling for hints helped.
At this point I was still not focused on OSCP as such, so I was using metasploit a plenty.

Let me just tell you, there is no point to shout RTFM at a person who simply hasn't learned a certain vulnerability before. Things such as shellshock have a few telling signs, true. But before you learn them, it is hard to guess what a certain machine is vulnerable to when your enumeration methods are still developing. I learned that this is trained over time, and by seeing certain exploits in action, rather than things being obvious or something that you resolve by reading `man` pages. 

I have also attended a number of conferences and CTF's in recent years, which keeps me motivated and fresh. Also a bit of experience with hacking through the night will go a long way. It is good to know your limits, strengths, weaknesses, prior to taking the exam.

### Covid-19

Watching the world around me slowly descend into lockdown, with my holiday plans and flights to far off lands cancelled. Conferences postponed, then cancelled. My work going remote and offering some voluntary furlough. My French language courses being almost suspended and ill prepared for remote world. I saw an opportunity. With 20% workload gone, and 2 weeks of holidays to take before July, I could give OSCP a shot. Finding another time slot of 90 days would be rather hard, so I paid upfront and began waiting for my course materials. (All with my own money, as you'd probably imagine, asking your competency manager to quickly approve a training that isn't vital to your performance in current position would be unlikely during recession, especially when company motto becomes "cash is king" for the foreseeable future)
![Image](https://eqqn.github.io/images/not-sans.JPG)


### Study regime
Take notes. Better yet if you have practiced note taking before in write-ups, or while doing courses. My notes went from mostly gibberish and only useful at the point of time, to a proper reference and walkthroughs for certain problems. As I was going through course materials, I took notes, even if there is a PDF available, you will want your notes when the time comes. The PDF is 800+ pages long. It's not easy to search it. Make your own notes for each chapter. It will be leaner and better, and something your own, with your own comments and experiences.

First few weeks I kept notes on how much and how fast I progressed. I thought to finish the course materials within the first month. in first 15 days, I had 9 study sessions, some on weekends. But then I felt like burning out and didn't want to study for days. And my furlough wasn't until the next month. I slowed down, learned mostly when I felt my mind open for it. Can't fill a cup that is already full. 

Just a bit before the end of study materials I finally started knocking on the lab network. Now I was having 8 days off each month + any national holidays and weekends to focus on the lab. I also started hacking some boxes during my work hours. Some days I woke up fresh at 6 am and wanted to hack. The other times, I was taking a few days breaks. I put the CTF's on hold for the time being.

### The meta
Before the course, and during it, immerse yourself into the community. Join the discord. Watch youtube videos, OSCP experiences and advice they can give. Have a point of reference. The more details you glean from others, the less time you will spend worrying and more time working on the things that matter.

Choose if you want to do the lab report or not. It's 5 extra points. Personally I didnt #READ-HARDER and didnt notice that lab report MUST include course exercises ( not just machines hacked write-ups). One month in, I already knew that it would take me at least a week of effort to re-trace and it isn't wort it. For me, the week or two is better spent on the labs. 
Again, learn to document. Make it your reflex to get proof.txt files with ipconfig. Even if you aren ot making a report. Document every step. Every command. Learn the lessons a machine provides. Use Joplin or similar for note taking.

### Buffer overflows
I have to give it to the OffSec staff, the buffer overflow materials were top notch. My classes on software security in university( which I aced by the way) could only go as far as 
`python -c 'print "A"*300'` and debugging with gdb and merely attempting to overwrite the EIP with something useful. GDB is a great tool, but each time I see  
"Smashing The Stack For Fun And Profit"  my mind rejects it ( I did use PEDA-GDB for some reverse engineering and stepping over instructions for CTF challenges, but that's it). 
Offsec teaches you BO's on both Windows and Linux, gives you a decent debugger to start with and explains not only how to crash a program, but also shellcode it and get execution on a remote machine. They also provide you with some binaries to practice on. It was good practice as I wasn't sure I internalized the learnings. I dedicated a couple days of my 14 day exam prep to creating exploits for them, re-writing my buffer overflow notes and really taking the lessons in. And now I had experience with 4 binaries, not just the 2 from the course material.

### Course materials
The video materials are voiced by the most fitting voice for the task. If you are paying attention, he will mention that he is the author of a certain service/tool at some point. Over time you will adopt the catchphrases such as "Nice" and "Excellent" in your reports, also whenever your elaborate exploit works on the first try.

The first half of the training is more passive. It's mostly enumeration. But it becomes a good reference for the thoroughness you will need to apply later on. The flaws are easy and something you've seen or done before and not too advanced. Then it ramps up in difficulty and new concepts ( at least for me) with buffer overflow, active directory, tunelling, privesc.

They also demonstrate a pentest which you can follow along. I started a bit late at 8pm thinking there will be some checkpoint in form of credential or such like HTB often does, but seeing it unfold I stayed up until 5AM just trying to finish it off in one go. Do dedicate a day for it. Start early. It's the grand slam of penetration test that one day we can hope to achieve.

In my opinion the lab network should be introduced earlier on. And introduced properly in the video materials. When you follow video materials or PDF, you know it's there and you log in to one of the remote desktop clients dedicated for classes, but they only introduce the sandbox rather late.

The course network is split into 3 categories

* Your clients - VMs to use while learning
* sandbox - this is where practical lessonsof video/pdf take place
* Lab network - very little mention of it, but thats where the most action and self-learning happens. You go into OffSec forums to learn more about it

The lab will teach you what is needed for the exam and what the course materials haven't quite covered (in detail). Every few boxes I learned a lesson from it. On enumeration, on exploitation, or privesc. I also opted to not use metasploit ( and in most cases it didn't seem necessary) to be able to perform same or similar exploits during exam if needed. Every now and then I get stuck, I look at the forums. I learned how to use them too, not going after spoiler posts, but rather something related to what I am also doing. Perhaps I am doing it wrong. I am here to learn after all. Not to rack up points, but get lessons that can only be aquired by experience. 

When people say the course materials is 20% and lab is 80%, thats what they meant. As I spent 50% of my time on materials I was getting worried. But I also learned from some OSCP's that they also spent a month on the materials ( and that was when they were less than half as long)
So stick to the course materials and learn the methodology. 
I personally rooted "only" 17 machines, 2 of which they provide write-ups for ( which I was trying to take in as much as I can to learn the methodology OffSec expects of us)
The Active Directory machines are also amazing and I was really enjoying my lessons on them.
Realistically, I only saw the first half of the lab network, and didn't pivot into others yet. I would like to, but there may be other resources for learning these things at a less steep price ( looking at HTB labs or VirtualHackingLabs which emerged recently).
I know I can perform far more advanced attacks in CTF's or learn form them. Then again, I haven't encountered most of the big 4, save for "Pain" which is probably the easiest of the few. Although the scale of difficulty has shifted with the addition of 2020 materials. Although the exam was the only driver for the labs to begin with, it feels like there is a lot left to learn.

### Exam experience and timeline

Before the exam I had 2 weeks off work to avoid any context switching. I know this is an extreme luxury and a bit of a sacrifice, a gamble, a privilege. I put a lot of faith into succeeding. At the very least I tried to think that I will learn something out of the experience if I fail, but with the exam approaching I was antsy. I continued knocking on lab machines, checked some resources on WIndows Privesc as I was not feeling confident in my abilities ( as much as I like TCM's Ethical hacking course, I didnt like the windows privesc one. Perhaps because I looked into it AFTER I was through with PWK materials, I didn't see the depth early on. I tried to look at one certain attack, but instead found more depth in Ippsec's video. Perhaps its for OSCP prep, I didn't look very deep into it)

The 2 days before exam I took a break. What's learned is learned, I can't cram more or afford to get exhausted before exam. I biked to the sea, went on a run, shopped for trainers. I avoided the screen. During the studies I learned more about myself. I knew I had to come hungry(to perform penetration testing. dont skip breakfast). I performed best when I took a few days off. Sometimes after 5 days of work I didn't look at materials, because I was not in shape to study or willing to take the lessons. Exam night I got some sleep ( whole 5-6 hours, which is better than none that some people reported)

Exam time. I read the proctoring guidelines in advance. I was ready. My starting time was the usual time I start my day, at 10AM. Before lunch at 12:30 I got down the buffer overflow. 14:30 - 10 pointer. Then there was a bit of a shock. Both of the 20s , very different, but same in that they were not responding as I thought they would. On one I had the webshell, but couldn't elevate. The other exploit wasn't working. And I was not familiar with the service to understand the exploit. At that point I was really suffering and thinking where I had gone wrong. It's dinner time and I call another break. I keep ruling out more and more possiblities. The unfamiliar service had a metasploit module, so I gave it a try, it performed same as non metasploit at first glance, didn't return anything.
I go back to the webshell, try things differently and it works! Its 22:50. At least I will know next time what can be done, I have low priv shell, but still severely lacking points. I check the other machine in the meantime... The metasploit module was working! Its 23:04.

Now the #tryharder really kicks in. If I escalate privileges on these 2 machines, I don't need to look at the 25 pointer. 
I was so tunnel visioned that I ignored the 25 point host besides nmap scan and checking the web service. It was also a very unfamiliar combination of ports to me. 

I knew that privilege escalation is just ruling out things until you have tried them all. So I dedicated the rest of exam time to that. 
00:04, first privesc complete. Now I am really on adrenaline. 00:44 the other one rooted. 15 hours in I have passing points. I am tired. my head hurts. But I am happy as hell. I get a glass of milk and some bread with jam, a quick shower. Now Offsec tells us they are expecting our life to still happen, to eat and sleep and I take that to heart. I spend the next 3 hours retracing my steps, taking every screenshot I can, copy pasting commands to my exam notes. 
its 4AM. I've been up for 20 hours. I make sure my notes are okay, and 4:30 I ask the proctor to end the exam early. 
I sleep, dead tired, but content. It was a long night ;) .Now normally I don't get so riled up in CTF's because I am not at the level to compete for the podium, and it's mostly a friendly event where I connect with my friends and network, rather than strictly hacking. This exam was totally different. I was invested.

### The exam report
The next day I am still a bit early, I am at my desk before 10AM with a cup of coffee. I used the template offsec provided, but edited it on my break days to fit my style a bit. I wanted a report that represents me and OffSec best, and lets face it - the template is a bit dated. At 12:30 I go out for lunch. Give my SO a call (we are stranded thanks to the current situation). Before 2pm I am back at my desk. It took me until 23:00s , then I finished my submission before midnight. Indeed, it took me close to 13 hours. I didn't anticipate that, but it was fulfilling work, knowing that I did everything right (at least form my point of view). Looking at the 25 pointer nmap scan outputs I believe I understood where it wanted you to go, so next time I see that port/service I will know where to start.

### Summary? 
A lot of people I follow or admire for their work in the domain are OSCP's. Lately I also saw some jokes saying its "Google, the cert". Some define themselves as OSCP's which I considered a bit of a stretch. I expected people to shout "try harder" whenever I need help, but the community and offsec staff are much warmer than that.

Dont use my example as a guide, but you may use it as a point of reference. Take care of yourself, know that if you fail, you will still have gained a valuable lesson out of the experience. Don't give up. Regardless if you earn the OSCP, you will learn more through this training than any other "entry level" certification and other certs at this price range. All of this will make you a better cybersecurity professional.  Some professional pentesters get their ass handed by the exam the first time around too. Persistence is key!
