## Introduction

Ever wondered what it is like to work in a high tech manufacturing company? Earlier in 2022 I was looking for career change and jumped on the opportunity to work for a high precision Swiss 🇨🇭 machine manufacturing company.

![5-axis-machine](https://eqqn.github.io/images/5-Axis-CNC-Milling-Machines.png)

They make all kinds of **very expensive** CNC, EDM and other machines for customers globally in all kinds of industries - automotive, watchmaking, luxury goods, aerospace, medical, IT, military. The parts are either used directly, or machined into molds for plastic and other materials.


P.S.: I will avoid naming the company I work for, so images and examples will be rather general, but accurately represent the industry (although shouldn't be hard to guess,find. Don't take it as a challenge). Opinions my own and so on.
[![Assembly-lines](https://eqqn.github.io/images/youtube_studer_thubmbnail.jpg)](https://www.youtube.com/watch?v=Dbrcj4qm6So)

We have an assembly line not unlike in the video above, and industrial secrecy is important (no photos allowed on the shop floor).

## What's inside a machine?

As you can probably guess, a 3-5 axis CNC machine will contain a lot of electromechanical parts. 
PLCs, motors, connection cables, and buses for industrial protocols.

From security and usability point of view, it makes sense that they are isolated and sit behind an internal switch. To control the CNC, a control unit is used, called TNC. This is familiar territory - a rugged terminal with screen, keyboard, USB peripherals and Windows or Linux OS running on top.

Additionally, turnkeys and emergency shutdown switches are there to ensure safety of the operator.

TNC host executes the design and jobs loaded onto the machine, displays status, allows to edit job instructions/code and so on.
![TNC](https://eqqn.github.io/images/siemens-neu_7.jpg)

## So what about the security?

When assessing such machines from security point of view, a series of questions arises:
- how the machine is connected to the rest of the network
- what ports and [protocols](https://en.wikipedia.org/wiki/Category:Industrial_Ethernet) does it expose? Proprietary or open? Fieldbus, modbus, etherCAT, profinet, OPC-UA, MTconnect?
- standard services running? VNC, ssh, SMB, ftp? 
- APIs, [automation protocols](https://en.wikipedia.org/wiki/List_of_automation_protocols)?

As you need to transport the job files to the machine ( a complex job or job sequences can be gigabytes!), you need connectivity. There are probably multiple CAD/CAM engineers who will need access. On the other hand, the job files are intellectual property allowing to replicate the part. Which can be a mold for a soda bottle, or turbine for an airplane. 

Depending on spec or requirements, additional software can be installed - easier calibration tools, clients, license managers, even [app stores](https://www.okuma.com/okuma-app-store) .

And is just scratching the surface.

## Industry 4.0

Whether you call it Cloud, Digital Transformation, Manufacturing 4.0... It's there. Across the industry, the untapped potential for better automation, big data analytics and ERP integration is there. But it's also hampered by lack of interoperability on many levels, as well as valid security concerns.

Try to find [top 10 CNC machine brands](https://www.stylecnc.com/blog/top-10-best-cnc-machine-manufacturers-brands.html). I personally don't know half of the list and could add a bunch of more. And that's not even covering high-performance segments. Or different workflows - it's not just milling, but also grinding, [wire EDM](https://en.wikipedia.org/wiki/Electrical_discharge_machining)(not dance music) and laser/additive(think of 3D printing with metal powders heated by laser!) manufacturing. These machines are even more specialized and made by smaller companies.
As there are no prevailing common standard, we have many different workflows, controls, APIs and protocols.

Because of a fragmented market it is probably not at the forefront of software industry[citation needed]

### Trendmicro industry 4.0 research

Recently published [research by Trendmicro](https://www.trendmicro.com/vinfo/us/security/news/internet-of-things/uncovering-security-weak-spots-in-industry-4-0-cnc-machines) highlights some of the vulnerabilities around the few machines they tried. They also make a very good point on the length of the supply chain.

![subbly_chain](https://eqqn.github.io/images/supply_chain_tm.JPG)

The machines subjected to the research seem to be chosen mostly due to their availability/access in collaboration with University of Milan. For example, the [Haas Super Mini Mill](https://www.haascnc.com/content/haascnc/en/build-and-price/choose-options.SMINIMILL.html) 2 CNC starts at 50k$, but their factory floor machines are more likely in 150-300k range if you go anything beyond standard.

The researchers claim there has not been (m)any (public)reports as theirs, that focus on assessing real world implementations and it's easy to see why. 

Do penetration testers test industrial networks regularly? How much of the machines are off limits? Considering the cost of equipment, scrap, downtime, might be logical to exclude those targets... 

However, for attackers these present another target for espionage or disrupting of operations.

## Industrial networks

The networks in which these machines operate are also rather diverse. Each of the segments mentioned earlier have their own caveats. SCADA? Air-gaps?

A tool shop with 5 machines from one brand will probably be different from one combining multiple machine vendors and processes into one megafactory. 

Different network models , trends , expectations, realities. 

Software - non software companies, IT departments, differences and challenges. All likelyhood OT is behind IT ( evidence in popular research? ) 

Players , cloud providers.

Security teams

## silver lining

standardisation is developing, some open protocols, big investments are coming from cloud providers.
Instead of every tool manufacturing coming with their own platform, some independent entities trying to connect them all are springing up.

### DevSecOps

BYOS - Bring your own security.

SAST

DAST

SCA

Secret scanning

Container scans




### Links , references
[GitHub](https://github.com/) 

[Cool PLC security talk](https://www.youtube.com/watch?v=JtsyyTfSP1I)

[Trend micro industry 4.0 CNC report PDF](https://documents.trendmicro.com/assets/white_papers/wp-the-security-risks-faced-by-cnc-machines-in-industry-4-0.pdf)
