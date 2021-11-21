# Underscore Bypass

## Description 
The regex used for origin verification contains an underscore (`_`) character.

`wwww.example.com` trusts `www.sub_example.com`, which could be an attacker's domain.

**Severity**: High

## Exploit 
If the target is `sub.example.com`, make requests from `sub_example.com`.
