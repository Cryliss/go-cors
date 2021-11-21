# Third Party Origin

## Description 
Host has whitelisted a third party host for cross origin requests.

**Severity**: Medium

## Exploit 
If the third party host has an XSS vulnerability it can be used to exploit your CORS configuration.

### Example
`https://trusted-origin.example.com/?xss=<script>CORS-ATTACK-PAYLOAD</script>`
