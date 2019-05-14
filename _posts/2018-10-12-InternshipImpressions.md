# Your containers are (NOT) Secure
## or every container hysteria reviewed
###### Disclaimer: I work in a company with many Kubernetes deployments in Google Cloud environments, so this post will mostly cover my perspective from inside the filter bubble.

So you are developing your services/applications on the cloud.
Naturally you, or the management, or the customer asks the question:__"How secure is this?"__
If you listen to the buzz, it's a scary world out there. 
## What if you pull a compromised container?! [2]
Most articles fail to mention that most of the downloads are done on compromised orchestration/deployment infrastructure, rather than some clueless developer starting their Dockerfile with
```
FROM ubuntu:15.04_deprecated_4_year_old_version_tutorial_example
```
or [3]
```
FROM docker123332/mysql
```
Assumptions are not a good basis for your security policy. But it is easy to verify if the DockerHub repository you are pulling from can be trusted. Most relevant OS/DB/language container images are published by the companies themselves. For others, Docker also maintains image packages. 

> Docker, Inc. sponsors a dedicated team that is responsible for reviewing and publishing all content in the Official Images. This team works in collaboration with upstream software maintainers, security experts, and the broader Docker community.[4]

You can at the very least assume that Docker is acting in it's self interest and would avoid doing things that *make them look bad* while they secure another round of funding. [5] Docker also performs scans on the images going to official repositories, but thats for another post.

### Prevention:
Verify the image publisher, avoid old versions, use common sense. 

## What if someone hacks my cluster?!
Is your kubernetes API publicly exposed? 
If you answered, "Yes, I did go out of my way to expose my most critical orchestration component to the public internet without any authentication" , the miners will thank you.

# But Alpine Linux had default Root password as NULL for years![6]
Did you know that you are running as root inside of the container all along? The container doesn't even come with "suid" binaries (there are no users, no logins). Unless you actively make effort to install the packages, expose a port, install a ssh server on a container, enable root logins with empty passwords, it does not concern you. That is not a sane use case for a container. 

# The vendor told me my container has XX vulnerabilities! 
Most of the containers you see are Linux systems, that have the same vulnerabilies as you'd have on your own system. None of them make your machine a dumpster fire. For containers it is even more true, since you only use limited functionalities of the linux system. If your server technology is not outdated, and web application framework is sound, there is nothing to worry about your machine being taken over by hackers.

# Takeaway
You can't always expect your developers, Docker, open source community ( and anywhere else in supply chain) to do the right thing in terms of security. Mistakes happen. Reliable and secure systems must be fault tolerant and a lapse somewhere must not compromise the whole system. 3rd party tooling and solutions can help you find these lapses or allow visibility to prove that things are working correctly.

Understand the system. Containerization should not be seen as a black-box. Look into security controls each layer provides.
Be informed before jumping to a vendor solution with another blackbox to stack on top of your cluster and developers mindspace.


[1] https://www.zdnet.com/article/kubernetes-first-major-security-hole-discovered/

[2] https://threatpost.com/malicious-docker-containers-earn-crypto-miners-90000/132816/

[3] https://www.bleepingcomputer.com/news/security/17-backdoored-docker-images-removed-from-docker-hub/

[4] https://docs.docker.com/docker-hub/official_images/ , https://docs.docker.com/docker-hub/publish/trustchain/

[5] https://www.crunchbase.com/organization/docker#section-funding-rounds

[6] https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-5021

29 Docker Security tools compared : https://sysdig.com/blog/20-docker-security-tools/


container security vendors
container security vulns
debian stretch Oracle BerkeleyDB critical
