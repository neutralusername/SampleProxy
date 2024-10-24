# sample proxy server

clients visit the proxy servers' url (default configuration: localhost:8080) in the browser and attach the desired url as query parameter such as

```
localhost:8080/?url=https://www.google.com
```

and receive the desired pages' content back as http response

<br/>
    
deploy on server with access to the desired url and contact via https such as

```
https://*your ip*:8080/?url=https://www.google.com
```

browser extensions exist that can redirect all urls to "localhost:8080/?url=...." 
