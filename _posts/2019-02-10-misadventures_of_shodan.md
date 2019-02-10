# Adventure with HIBP and Shodan

I am a fan of Troy Hunt. I find "HaveIBeenPwned"( https://haveibeenpwned.com/ ) quite genius. 
It provides value to wide population, via simple website and improves the security of peoples data by letting them know which of their passwords or emails could have been leaked in the past. 
Working in a corporate environment, it is sometimes hard to make an impact outside of your bubble during your day-to-day work.
But sometimes you will stumble upon servers or websites that could use a little nudge to prevent damage. This is a story about how I tried to locate owners of a random AWS server in a bid for them to fix the issue. With a light dose of OSINT and forensics.

I've been looking at the Domain Search functionality of the portal, which allows you to check 
if any emais under your email server were involved in a breach. For that, you need to verify the ownership of the domain,
by responding from a certain email, i.e. Security@eqqn.com, or embedding TXT DNS record with a certain string.
Another easy method is uploading a file called "have-i-been-pwned-verification.txt" to your root directory of the web.
![Image](https://eqqn.github.io/images/HIBP_verification.jpg)

So, let's look it up on our favorite search engine, to see who is concerned about their breaches! 
![Image](https://eqqn.github.io/images/shodan_hibp.jpg)

Only two entries. Both of them are the same AWS server, on different IPs.
Following it leads to...
![Image](https://eqqn.github.io/images/web-root-redacted.jpg)

Oh. That's not good. Someone misconfigured their server, or inadvertedly exposed it  
by trying to provide the verification file and left the website root open for a few years.
Another error is putting SQL dumps on the webserver. ( thx friend for pointing it out)
So now it's up to me to find out who owns this server and ~~crack the passwords from SQL dump and take over the CEO's account~~ 
**notify them**,because thats what the good guys do.
Analyzing the files served, you can see that it hosts some wordpress files from 2013 to 2016. 
Although I am not very familiar with Wordpress, I found some configuration files in them that are quite humorous (and dangerous).
Worpress uses MySQL database to store the blogs and information, so each blog has their own database. It turns out the configuration files
contain information about said database, including credentials.
![Image](https://eqqn.github.io/images/wordpress_backup_code.JPG)

The SQL dumps are from 2 websites, dated 2013, containing some (12) user account password hashes, news entries, and some test data.
There are also redirects to the same website versions in Polish, Finnish and Swedish.
One of the links lead to a small payment service company based in Finland. Looking deeper, 
it seems three different fintech(loantech?) companies are managed by the same business group. Looking up the small Dev company name on 
LinkedIn results in one of the board members of the group.
![Image](https://eqqn.github.io/images/fintech_board_member.jpg)

And the same consulting group serviced both the Small Dev company and Main Loan company.
![Image](https://eqqn.github.io/images/consulting.jpg)

Also seen in the publicly accessible annual review for NASDAQ Finland
![Image](https://eqqn.github.io/images/annual_review.jpg)

By this point we know with a decent amount of certainty who to notify.
I contacted one of the product managers of Small dev company and CC'd the info@thatcompany just in case.
The PM email bounced, info made it through.

Some couple weeks later I revisited the website and it gave me proper 403 response. I did not receive a response or acknowledgement,
but I'd like to think I helped. If my email was read and forwarded to right people, perhaps it is in their best interest to sweep it under a rug and pretend it never happened after the fix. Noone wants to deal with GDPR notices.


