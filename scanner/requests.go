package scanner

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// newRequest creates a new HTTP request with the provided method and URL
func (s *Scanner) newRequest(method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		s.l.Log.Err(err).Msg("s.addRequestMethod")
	}
	return req
}

// createClient creates a new HTTP client
func (s *Scanner) createClient() *http.Client {
	// Set the request timeout
	timeout, err := time.ParseDuration(s.conf.Timeout)
	if err != nil {
		// If we weren't able to parse the configuration's timeout value, let's
		// make it equal to our default value of 10 seconds.
		timeout, _ = time.ParseDuration("10s")
	}

	transport := &http.Transport{
		MaxIdleConns:    30,
		IdleConnTimeout: time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: time.Second,
		}).DialContext,
	}

	redirect := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	c := http.Client{
		Transport:     transport,
		CheckRedirect: redirect,
		Timeout:       timeout,
	}
	return &c
}

// sendRequest creates a new requests and performs it
func (s *Scanner) sendRequest(c *http.Client, url, origin, method string, headers Headers) (string, string, error) {
	// Make sure method has a value
	if method == "" {
		method = "GET"
	}

	// Create a new request
	req := s.newRequest(method, url)

	// Set the request origin header
	req.Header.Set("Origin", origin)

	// Now add the other request headers
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	// Perform the HTTP requests
	resp, err := c.Do(req)
	if err != nil {
		s.l.OutErr("s.sendRequest: failed to send request to %s - %+v", url, err)
		return "", "", err
	}

	// Get the response headers and return them 
	acao := resp.Header.Get("Access-Control-Allow-Origin")
	acac := resp.Header.Get("Access-Control-Allow-Credentials")

	return acao, acac, nil
}
