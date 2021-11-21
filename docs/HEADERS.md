# CORS Headers

## [Access-Control-Allow-Origin](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin)  
Indicates whether the response can be shared.

### Syntax
`Access-Control-Allow-Origin: *
Access-Control-Allow-Origin: <origin>
Access-Control-Allow-Origin: null`

### Directives
`*`  
For requests without credentials, the literal value "*" can be specified as a wildcard; the value tells browsers to allow requesting code from any origin to access the resource. Attempting to use the wildcard with credentials results in an error.

`<origin>`  
Specifies an origin. Only a single origin can be specified. If the server supports clients from multiple origins, it must return the origin for the specific client making the request.

`null`  
Specifies the origin "null".

## [Access-Control-Allow-Credentials](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials)  
Indicates whether or not the response to the request can be exposed when the credentials flag is true.

### Syntax
Access-Control-Allow-Credentials: true

### Directives
`true`  
The only valid value for this header is true (case-sensitive). If you don't need credentials, omit this header entirely (rather than setting its value to false).

## [Access-Control-Allow-Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers)  
Used in response to a preflight request to indicate which HTTP headers can be used when making the actual request.

### Syntax
`Access-Control-Allow-Headers: [<header-name>[, <header-name>]*]
Access-Control-Allow-Headers: *`

### Directives
`<header-name>`  
The name of a supported request header. The header may list any number of headers, separated by commas.

`*` (wildcard)  
The value `*` only counts as a special wildcard value for requests without credentials (requests without HTTP cookies or HTTP authentication information). In requests with credentials, it is treated as the literal header name `*` without special semantics. Note that the Authorization header can't be wildcarded and always needs to be listed explicitly.

## [Access-Control-Allow-Methods](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods)  
Specifies the method or methods allowed when accessing the resource in response to a preflight request.

### Syntax
`Access-Control-Allow-Methods: <method>, <method>, ...
Access-Control-Allow-Methods: *`

### Directives
<method>
A comma-delimited list of the allowed HTTP request methods.

`*` (wildcard)  
The value `*` only counts as a special wildcard value for requests without credentials (requests without HTTP cookies or HTTP authentication information). In requests with credentials, it is treated as the literal method name `*` without special semantics.

##[Access-Control-Expose-Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers)  
Indicates which headers can be exposed as part of the response by listing their names.

### Syntax
`Access-Control-Expose-Headers: [<header-name>[, <header-name>]*]
Access-Control-Expose-Headers: *`

### Directives
`<header-name>`  
A list of zero or more comma-separated header names that clients are allowed to access from a response. These are in addition to the CORS-safelisted response headers.

`*` (wildcard)  
The value `*` only counts as a special wildcard value for requests without credentials (requests without HTTP cookies or HTTP authentication information). In requests with credentials, it is treated as the literal header name `*` without special semantics. Note that the Authorization header can't be wildcarded and always needs to be listed explicitly.

## [Access-Control-Max-Age](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age)  
Indicates how long the results of a preflight request can be cached.

### Syntax
Access-Control-Max-Age: <delta-seconds>

### Directives
`<delta-seconds>`  
Maximum number of seconds the results can be cached, as an unsigned non-negative integer. Firefox caps this at 24 hours (86400 seconds). Chromium (prior to v76) caps at 10 minutes (600 seconds). Chromium (starting in v76) caps at 2 hours (7200 seconds). The default value is 5 seconds.

## [Access-Control-Request-Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Request-Headers)  
Used when issuing a preflight request to let the server know which HTTP headers will be used when the actual request is made.

### Syntax
`Access-Control-Request-Headers: <header-name>, <header-name>, ...`

### Directives
`<header-name>`  
A comma-delimited list of HTTP headers that are included in the request.

## [Access-Control-Request-Method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Request-Method)  
Used when issuing a preflight request to let the server know which HTTP method will be used when the actual request is made.

### Syntax
`Access-Control-Request-Method: <method>`

### Directives
`<method>`  
One of the HTTP request methods, for example GET, POST, or DELETE.

## [Origin](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin)  
Indicates where a fetch originates from.

### Syntax
`Origin: null
Origin: <scheme> "://" <hostname> [ ":" <port> ]`

### Directives
`<scheme>`  
The protocol that is used. Usually, it is the HTTP protocol or its secured version, HTTPS.

`<hostname>`  
The domain name of the server (for virtual hosting) or the IP.

`<port>` Optional  
TCP port number on which the server is listening. If no port is given, the default port for the service requested (e.g., "80" for an HTTP URL) is implied.
