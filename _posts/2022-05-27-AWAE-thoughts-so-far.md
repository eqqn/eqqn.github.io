I thought OSWE would be a logical progression to OSCP, but after a week into the course, I am not so sure.

Heavily focused on SQL injection.

OTOH there is nothing bleeding edge about this. Old fashioned bugs, windows hosts and whatnot, 
does not really reflect my current enterprise experience, having extensively worked in dockerized, cloud, kubernetes environments...
My main take-away will be the methodology, and expanding my approach to some tests.

### NO tools *
As we know SQLmap is not allowed in exams and discouraged, along with other tools that make for productive work.
Cynical me wants to tell Offsec to perhaps write their own ASM and instructions by hand if they are taking things from foundations,
but kind me appreciates the lessons learned on how to manually assess a system without overwhelming it with a full blown scan.

Magic hashes, php, type juggling is very 2010s.

### Exam expectations
During OSCP labs, the learning proccess was clear. You have labs, and with each lab you take in a lesson :

1. how to google better
2. how to recognize certain flaws like shellshock
3. how to use kernel exploits
4. not getting stuck on using single language for exploitation
5. things specific to some framework
6. privesc methods 

for OSWE, I am struggling to to understand the expectation. The exam FAQ and documentation is VERY POOR. Almost by design, to provide
least information possible.
Doing course materials, we are role-playing vulnerability discovery. I say role-playing because we make huge leaps in code review or conclusions
to which exploit will work. Then you are given an exercise, to find a vulnerability that is very similar in the code base. I suppose that is 
the expectation of exam, one which I'd have very hard time to fulfill, at least in my current track. 
It might become more apparent once I start poking at the lab... 
