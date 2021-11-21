# Unescaped Dot Bypass

## Description 
The regex used for origin verification contains an unescaped dot (`.`) character.

`wwww.example.com` trusts `wwwaexample.com`, which could be an attacker's domain.

**Severity**: High

## Exploit 
If the target is `sub.example.com`, make requests from `subxexample.com`.
