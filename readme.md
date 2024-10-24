sample proxy server

clients visit the proxy servers' url (default configuration: localhost:8080) in the browser and attach the desired url as query parameter such as

```
localhost:8080/?url=https://www.google.com
```

and receive the desired pages' content back as http response

browser extensions exist that can redirect all urls to "localhost:8080/?url=...." 

configure https if usage goes beyond testing