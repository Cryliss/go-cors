# What is CORS?

As defined by the [MDN Web Docs](https://developer.mozilla.org/en-US/docs/Glossary/CORS)  
> Cross-Origin Resource Sharing (CORS) is a system, consisting of transmitting HTTP headers, that determines whether browsers block frontend JavaScript code from accessing responses for cross-origin requests.  
The (same-origin security policy)[https://developer.mozilla.org/en-US/docs/Web/Security/Same-origin_policy] forbids cross-origin access to resources. But CORS gives web servers the ability to say they want to opt into allowing cross-origin access to their resources.

Essentially, CORS adds new [HTTP headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers) that let servers describe which origins (domain, scheme, or port) other than its own are permitted read and load resources.

CORS failures result in errors that are *not available to JavaScript* for security reasons. These errors may only be viewed by looking at the browser's console for details, by right-clicking on the page and selecting `inspect`.


# What requests use CORS?
- [Fetch APIs](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API)
- Web Fonts
- [WebGL textures](https://developer.mozilla.org/en-US/docs/Web/API/WebGL_API/Tutorial/Using_textures_in_WebGL)
- Images/video frames drawn to a canvas using [drawImage()](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/drawImage)
- [CSS Shapes from images](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Shapes/Shapes_From_Images)

CORS relies on a mechanism by which browsers make a "preflight" request to the server hosting the cross-origin resource, in order to check that the server will permit the actual request with the HTTP `OPTIONS` request method. In that preflight, the browser sends headers that indicate the HTTP method and headers that will be used in the actual request. Upon approval, the client will then send the actual request to the server.

Such requests are usually done using [XMLHttpRequest](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest), which can make cross-site requests in any supporting browser. An example preflighted request:

```js
const xhr = new XMLHttpRequest();
xhr.open('POST', 'https://bar.other/resources/post-here/');
xhr.setRequestHeader('X-PINGOTHER', 'pingpong');
xhr.setRequestHeader('Content-Type', 'application/xml');
xhr.onreadystatechange = handler;
xhr.send('<person><name>Arun</name></person>');
```

Some requests do not triggers a CORS preflight request, such as `GET`, `HEAD`, and `POST`, and are usually performed using [Fetch](https://fetch.spec.whatwg.org/). Such requests are called *simple requests*, which must meet the following conditions:

-  The only headers which are allowed to be manually set are those which the Fetch spec defines as a CORS-safelisted request-header, which are:
    - Accept
    - Accept-Language
    - Content-Language
    - Content-Type

- The only allowed values for the Content-Type header are:
    - application/x-www-form-urlencoded
    - multipart/form-data
    - text/plain

- If the request is made using an XMLHttpRequest object, no event listeners are registered on the object returned by the XMLHttpRequest.upload property used in the request; that is, given an XMLHttpRequest instance xhr, no code has called xhr.upload.addEventListener() to add an event listener to monitor the upload.

- No ReadableStream object is used in the request. 
