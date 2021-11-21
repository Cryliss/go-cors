# Special Character Bypass

## Description
The origin verification is flawed and can be bypassed using a special character, that is not an underscore or back tick, such as: `- " { } + ^ %60 ! ~ ; | & ' ( ) * , $ = + %0b`.

**Severity**: High

## Exploit
If the target is `sub.example.com`, make requests from `sub<special_char>example.com`.
