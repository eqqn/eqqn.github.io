# Your containers are (NOT) Secure

###### Disclaimer: I work in a company with many Kubernetes deployments in Google Cloud environments, so this post will mostly cover my perspective from inside the bubble.

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



[1] https://www.zdnet.com/article/kubernetes-first-major-security-hole-discovered/

[2] https://threatpost.com/malicious-docker-containers-earn-crypto-miners-90000/132816/

[3] https://www.bleepingcomputer.com/news/security/17-backdoored-docker-images-removed-from-docker-hub/

[4] https://docs.docker.com/docker-hub/official_images/ , https://docs.docker.com/docker-hub/publish/trustchain/

[5] https://www.crunchbase.com/organization/docker#section-funding-rounds


29 Docker Security tools compared : https://sysdig.com/blog/20-docker-security-tools/





hello
container security vendors
container security vulns
debian stretch Oracle BerkeleyDB critical
