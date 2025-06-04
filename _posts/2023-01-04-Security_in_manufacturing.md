## Introduction

Ever wondered what it is like to work in a high tech manufacturing company? Earlier in 2022 I was looking for career change and jumped on the opportunity to work for a high precision Swiss 🇨🇭 machine manufacturing company.

![5-axis-machine](https://eqqn.github.io/images/5-Axis-CNC-Milling-Machines.png)

They make all kinds of **very expensive** CNC, EDM and other machines for customers globally in all kinds of industries - automotive, watchmaking, luxury goods, aerospace, medical, IT, defense. The parts are either used directly, or machined into molds for plastic and other materials.


P.S.: I will avoid naming the company I work for, so images and examples will be rather general, but accurately represent the industry (although shouldn't be hard to guess,find. Don't take it as a challenge). Opinions my own and so on.
[![Assembly-lines-](https://eqqn.github.io/images/youtube_studer_thubmbnail.jpg)](https://www.youtube.com/watch?v=Dbrcj4qm6So)
Titans of CNC has made an excellent video of one such companies.

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

Authenticated sessions and security considerations seem to be a rather recent development for some control providers.

Depending on spec or requirements, additional software can be installed - easier calibration tools, clients, license managers, even [app stores](https://www.okuma.com/okuma-app-store) .

And it's just scratching the surface.

## Industry 4.0

Whether you call it Cloud, Digital Transformation, Manufacturing 4.0... It's there. Across the industry, the untapped potential for better automation, big data analytics and ERP integration. But it's also hampered by lack of interoperability on many levels, as well as valid security concerns.

Try to find [top 10 CNC machine brands](https://www.stylecnc.com/blog/top-10-best-cnc-machine-manufacturers-brands.html). I personally don't know half of the list and could add a bunch of more. And that's not even covering high-performance segments. Or different workflows - it's not just milling, but also grinding, [wire EDM](https://en.wikipedia.org/wiki/Electrical_discharge_machining)(not dance music) and laser/additive(think of 3D printing with metal powders heated by laser!) manufacturing. These machines are even more specialized and made by smaller companies.
As there are no prevailing common standard, we have many different workflows, controls, APIs and protocols.

### Trendmicro industry 4.0 research

Recently published [research by Trendmicro](https://www.trendmicro.com/vinfo/us/security/news/internet-of-things/uncovering-security-weak-spots-in-industry-4-0-cnc-machines) highlights some of the vulnerabilities around the few machines they tried. They also make a very good point on the length of the supply chain.

![supply_chain](https://eqqn.github.io/images/supply_chain_tm.JPG)

The machines subjected to the research seem to be chosen mostly due to their availability/access in collaboration with University of Milan. For example, the [Haas Super Mini Mill](https://www.haascnc.com/content/haascnc/en/build-and-price/choose-options.SMINIMILL.html) 2 CNC starts at 50k$, but their factory floor machines are more likely in 150-300k range if you go anything beyond standard.

The researchers claim there has not been (m)any (public)reports as theirs, that focus on assessing real world implementations and it's easy to see why. 

Do penetration testers test industrial networks regularly? How much of the machines are off limits? Considering the cost of equipment, scrap, downtime, might be logical to exclude those targets... 

However, for attackers these present another target for espionage or disrupting of operations.

## Industrial networks

The networks in which these machines operate are also rather diverse. Each of the segments mentioned earlier have their own caveats. 

A tool shop with 5 machines from one brand will probably be different from one combining multiple machine vendors and processes into one megafactory. 

Air gapped networks and strict network segregation between IT/OT is widely used in mature organizations with strict security requirements. This makes it harder for machine providers to support remote servicing, get machine data that would feed into planning, inventory management and business ERP systems. 

Conversely, small shops might opt for a plug-and-play experience and more on-line features.

### Lifecycle of a machine

Because machines are built to last and cost large amounts of money, it is not uncommon to see older models still operating, being leased or resold. With replacement parts and proper maintenance, it can rightfully outlast consumer hardware, but software does not age like wine... But any updates or instability could have issues on performance. While software updates are possible, replacing core computing components is unheard of, due to complex wiring, calibration and conditions that can only be provided in the OEM facility.

## Edge devices 

Luckily, there are some good solutions to counter some of the problems. Edge device in front of the machine to handle secure communication, extensibility and PKI based connectivity to cloud networks. This circumvents the limitations of HMI/controlers in terms of resources and can be set up with high security assurance protecting the industrial network, the machine, while also providing advanced functionality, timely updates and modern runtimes like [docker containers](https://industrial.softing.com/products/docker.html).

Microsoft seems to have a robust service offering in this segment, among other cloud providers.

## Manufacturers as software companies

Although *every company is a software company*, manufacturing is a rather distinct environment to work in so far. The IT and the perimeter has always been there, but the contemporary security practice is still developing. I was lucky to land in a team with already established DevOps proccess ( keep in mind that some companies only getting there) and managing security tasks in this context has been my bread and butter for a few years. Conversely, I am trying to learn more about the industry inside and outside. 

## Silver lining

Standardisation IS developing, as there are some open protocols that are heading in the right direction, like MT-connect and OPC-UA. Although MT-connect is neither authenticated, nor encrypted, it gained some ground across vendors and showed the benefits of standardised connectivity. 

The newer and better thought out OPC-UA is getting traction, as it supports strong encryption and auth, with latest versions relying on certificate services. While these services are well defined in the specification, good pluggable implementations are still around the corner. Another cool thing, was that [Pwn2Own organized an event to take a crack ](https://sector7.computest.nl/post/2022-07-opc-ua-net-standard-trusted-application-check-bypass/)at it. We can hope that this gradually improves the robustness as security holes get adressed. 

Instead of every tool manufacturing coming with their own platform, some independent entities trying to connect them all are springing up. Vendor agnostic solutions could be the way forward once industry settles on the communication standards. 

## German Machine Tool Builders association

VDW, German machine Tool Builders Association, offers some well documented guidance in their security manuals for Operators and Manufacturers. 

![Machines on the network](https://eqqn.github.io/images/VDWDiagram2.jpg)

[Security for machine tools](https://vdw.de/wp-content/uploads/2023/07/Security-for-machine-tools-EN.pdf) guide provides examples of a real-life deployment and security considerations, risk scenarios and strategy to follow in order to protect the machine or wider OT networks.
It is based around IEC 62443 Industrial security standard 

The work group consists of many industry players promoting machine tools not just in Germany, but also globally. The security workgroup further develops the guidance and shares best practices with the wider community.

## IEC 62443 Industrial security standard

While ISO 27001 covers a lot of IT security areas, it is not specific enough for industrial settings. The need for physical safety, consequences of security events and hardware in question is vastly different from standard IT environment.
IEC 62443 offers multiple chapters depending on application - whether you are developing software for your products, or expect so serve the industrial market, or operate a facility - the standard covers all areas.
This is de-facto standard to exchange security requirements, determine security levels and to ensure alignment with industrial security norms.
![IEC 62443 standards](https://eqqn.github.io/images/VDWIEC62443.jpg)

## Regulation 
Critical infrastructure and manufacturing industries are now facing regulatory pressure in various countries. Manufacturing being an important pillar of supply chain and strategic interests globally, securing it against increasing number of attacks is more important than ever. 
On the US side, government contracts have strict DoD requirements and implementing NIST SP 800-171.
On the EU side, EU Cyber Resilience Act requires baseline security to access European markets, and will come into force in late 2027.
(If you want to know what actions you need to take, I have composed a convenient checklist at [https://resilience-checklist.eu/](https://resilience-checklist.eu/) )

## Summary

Security work in manufacturing has some challenges that are common for all software companies, and some unique ones too. 

Standards help collaboration between operators, manufacturers and sets the framework to follow for industrial security.

Governments are pressuring the industry for improvement, despite the shortage of security professionals and funding in the sector.

### Links , references

[Microsoft IoT edge architecture](https://learn.microsoft.com/en-us/azure/architecture/guide/iiot-guidance/iiot-architecture)

[Cool PLC security talk](https://www.youtube.com/watch?v=JtsyyTfSP1I)

[Trend micro industry 4.0 CNC report PDF](https://documents.trendmicro.com/assets/white_papers/wp-the-security-risks-faced-by-cnc-machines-in-industry-4-0.pdf)

[Blackhat EU talk on the Trend Micro research](https://www.blackhat.com/eu-22/briefings/schedule/#abusing-cnc-technologies-28834)

[Security for Machine tool Builders - VDW Guide](https://vdw.de/wp-content/uploads/2023/07/Security-for-machine-tools-EN.pdf)

[Pwn2Own OPC-UA](https://sector7.computest.nl/post/2022-07-opc-ua-net-standard-trusted-application-check-bypass/)
