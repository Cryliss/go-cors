// Package scanner ...
package scanner

import (
	"net/http"

	tld "github.com/jpillora/go-tld"
)

// Test struct holds all our data related to the test
type Test struct {
	Acao    string            `json:"access-control-allow-origins"`
	Acac    string            `json:"access-control-allow-credentials"`
	Headers map[string]string `json:"headers"`
	Method  string            `json:"method"`
	Origin  string            `json:"origin"`
	Test    string            `json:"test"`
	URL     string            `json:"url"`
}

func (s *Scanner) reflectOrigin(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting reflect origins test")

	// Set the origin value for the request we're going to make
	origin := "https://crylis.io/"

	// Perform the request
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return tests, err
	}

	// Check if the allowed origin is the origin we used
	if acao == origin {
		// Let the user know a misconfiguration was found
		s.l.OutErr("s.reflectOrigin: Misconfiguration found for '%s'! acao: %s & acac: %s", r.URL, acao, acac)
	}

	// Save the test information
	t := Test{
		Acao:    acao,
		Acac:    acac,
		Headers: r.Headers,
		Method:  r.Method,
		Origin:  origin,
		URL:     r.URL,
		Test:    "reflect origin",
	}

	// Add the test to the tests array and return
	tests = append(tests, &t)
	return tests, nil
}

func (s *Scanner) httpOrigin(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting http origins test")

	// Set the origin value for the request we're going to make
	origin := "http://crylis.io/"
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return tests, err
	}

	// Check if the allowed origin is the origin we used
	if acao == origin {
		s.l.OutErr("s.httpOrigin: Misconfiguration found for '%s'! acao: %s & acac: %s", r.URL, acao, acac)
	}

	// Save the test information
	t := Test{
		Acao:    acao,
		Acac:    acac,
		Headers: r.Headers,
		Method:  r.Method,
		Origin:  origin,
		URL:     r.URL,
		Test:    "http origin",
	}

	// Add the test to the tests array and return
	tests = append(tests, &t)
	return tests, nil
}

func (s *Scanner) nullOrigin(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting null origins test")

	// Set the origin value for the request we're going to make
	origin := "null"

	// Perform the request
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return tests, err
	}

	// Check if the allowed origin is the origin we used
	if acao == origin {
		s.l.OutErr("s.nullOrigin: Misconfiguration found for '%s'! acao: %s & acac: %s", r.URL, acao, acac)
	}

	// Save the test information
	t := Test{
		Acao:    acao,
		Acac:    acac,
		Headers: r.Headers,
		Method:  r.Method,
		Origin:  origin,
		URL:     r.URL,
		Test:    "null origin",
	}

	// Add the test to the tests array and return
	tests = append(tests, &t)
	return tests, nil
}

func (s *Scanner) wildcardOrigin(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting wilcard origin test")

	// Set the origin value for the request we're going to make
	origin := "*"

	// Perform the request
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return tests, err
	}

	// Check if the allowed origin is the origin we used
	if acao == origin {
		s.l.OutErr("s.wildcardOrigin: Misconfiguration found for '%s'! acao: %s & acac: %s", r.URL, acao, acac)
	}

	// Save the test information
	t := Test{
		Acao:    acao,
		Acac:    acac,
		Headers: r.Headers,
		Method:  r.Method,
		Origin:  origin,
		URL:     r.URL,
		Test:    "wildcard origin",
	}

	// Add the test to the tests array and return
	tests = append(tests, &t)
	return tests, nil
}

func (s *Scanner) thirdPartyOrigin(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting third party test")

	// Set the origin value for the request we're going to make
	// Since we're testing third party origins, we're going to test some
	// common third party sites
	origins := []string{
		"http://jsbin.com",
		"https://codepen.io",
		"https://jsfiddle.net",
		"https://www.webdevout.net",
		"https://repl.it",
	}

	for _, origin := range origins {
		// Perform the request
		acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
		if err != nil {
			return tests, err
		}

		// Check if the allowed origin is the origin we used
		if acao == origin {
			s.l.OutErr("s.thirdpartyOrigin: Misconfiguration found for '%s'! acao: %s & acac: %s", r.URL, acao, acac)
		}

		// Save the test information
		t := Test{
			Acao:    acao,
			Acac:    acac,
			Headers: r.Headers,
			Method:  r.Method,
			Origin:  origin,
			URL:     r.URL,
			Test:    "third party origin",
		}

		// Add the test to the tests array
		tests = append(tests, &t)
	}
	return tests, nil
}

func (s *Scanner) backtickBypass(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting backtick bypass test")

	// Parse the URL so we can use it to form our request origin
	url, err := tld.Parse(r.URL)
	if err != nil {
		return tests, err
	}

	// Set the origin value for the request we're going to make
	origin := "https://crylis.io%60" + url.Domain + "." + url.TLD

	// Perform the request
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return tests, err
	}

	// Check if the allowed origin is the origin we used
	if acao == origin {
		s.l.OutErr("s.backtickBypass: Misconfiguration found for '%s'! acao: %s & acac: %s", r.URL, acao, acac)
	}

	// Save the test information
	t := Test{
		Acao:    acao,
		Acac:    acac,
		Headers: r.Headers,
		Method:  r.Method,
		Origin:  origin,
		URL:     r.URL,
		Test:    "backtick bypass",
	}

	// Add the test to the tests array and return
	tests = append(tests, &t)
	return tests, nil
}

func (s *Scanner) postDomainBypass(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting postdomain bypass test")

	// Parse the URL so we can use it to form our request origin
	url, err := tld.Parse(r.URL)
	if err != nil {
		return tests, err
	}

	// Set the origin value for the request we're going to make
	origin := "https://" + url.Domain + "." + url.TLD + ".cryls.io"

	// Perform the request
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return tests, err
	}

	// Check if the allowed origin is the origin we used
	if acao == origin {
		s.l.OutErr("s.postDomainBypass: Misconfiguration found for '%s'! acao: %s & acac: %s", r.URL, acao, acac)
	}

	// Save the test information
	t := Test{
		Acao:    acao,
		Acac:    acac,
		Headers: r.Headers,
		Method:  r.Method,
		Origin:  origin,
		URL:     r.URL,
		Test:    "post domain bypass",
	}

	// Add the test to the tests array and return
	tests = append(tests, &t)
	return tests, nil
}

func (s *Scanner) preDomainBypass(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting predomain bypass test")

	// Parse the URL so we can use it to form our request origin
	url, err := tld.Parse(r.URL)
	if err != nil {
		return tests, err
	}

	// Set the origin value for the request we're going to make
	origin := url.Scheme + "://crylis.io." + url.Domain + "." + url.TLD

	// Perform the request
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return tests, err
	}

	// Check if the allowed origin is the origin we used
	if acao == origin {
		s.l.OutErr("s.preDomainBypass: Misconfiguration found for '%s'! acao: %s & acac: %s", r.URL, acao, acac)
		return tests, nil
	}

	// Save the test information
	t := Test{
		Acao:    acao,
		Acac:    acac,
		Headers: r.Headers,
		Method:  r.Method,
		Origin:  origin,
		URL:     r.URL,
		Test:    "pre domain bypass",
	}

	// Add the test to the tests array and return
	tests = append(tests, &t)
	return tests, nil
}

func (s *Scanner) underscoreBypass(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting underscore bypass test")

	// Parse the URL so we can use it to form our request origin
	url, err := tld.Parse(r.URL)
	if err != nil {
		return tests, err
	}

	// Set the origin value for the request we're going to make
	origin := "https://crylis.io_"+ url.Domain + "." + url.TLD

	// Perform the request
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return tests, err
	}

	// Check if the allowed origin is the origin we used
	if acao == origin {
		s.l.OutErr("s.underscoreBypass: Misconfiguration found for '%s'! acao: %s & acac: %s", r.URL, acao, acac)
	}

	// Save the test information
	t := Test{
		Acao:    acao,
		Acac:    acac,
		Headers: r.Headers,
		Method:  r.Method,
		Origin:  origin,
		URL:     r.URL,
		Test:    "underscore bypass",
	}

	// Add the test to the tests array and return
	tests = append(tests, &t)
	return tests, nil
}

func (s *Scanner) unescapedDotBypass(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting unescaped dot bypass test")

	// Parse the URL so we can use it to form our request origin
	url, err := tld.Parse(r.URL)
	if err != nil {
		return tests, err
	}

	// Set the origin value for the request we're going to make
	origin := "https://" + url.Subdomain + "S" + url.Domain + "." + url.TLD

	// Perform the request
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return tests, err
	}

	// Check if the allowed origin is the origin we used
	if acao == origin {
		s.l.OutErr("s.unescapedBypass: Misconfiguration found for '%s'! acao: %s & acac: %s", r.URL, acao, acac)
	}

	// Save the test information
	t := Test{
		Acao:    acao,
		Acac:    acac,
		Headers: r.Headers,
		Method:  r.Method,
		Origin:  origin,
		URL:     r.URL,
		Test:    "unescaped dot bypass",
	}

	// Add the test to the tests array and return
	tests = append(tests, &t)
	return tests, nil
}

func (s *Scanner) specialCharactersBypass(c *http.Client, r *Request, tests []*Test) ([]*Test, error) {
	s.l.Out("Starting special characters bypass test")

	// Parse the URL so we can use it to form our request origin
	url, err := tld.Parse(r.URL)
	if err != nil {
		return tests, err
	}

	// An array of special characters we're going to use in our tests
	specialChars := []string{"-", `"`, "{", "}", "+", "^", "%60", "!", "~", ";", "|", "&", "'", "(", ")", "*", ",", "$", "=", "+", "%0b"}
	for _, char := range specialChars {
		// Set the origin value for the request we're going to make
		origin := "https://crylis.io" + char + url.Domain + "." + url.TLD

		// Perform the request
		acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
		if err != nil {
			return tests, err
		}

		// Check if the allowed origin is the origin we used
		if acao == origin {
			s.l.OutErr("s.specialCharactersBypass: Misconfiguration found for %s using special character '%s'! acao: %s & acac: %s", r.URL, char, acao, acac)
		}

		// Save the test information
		t := Test{
			Acao:    acao,
			Acac:    acac,
			Headers: r.Headers,
			Method:  r.Method,
			Origin:  origin,
			URL:     r.URL,
			Test:    "special character '" + char + "' bypass",
		}

		// Add the test to the tests array
		tests = append(tests, &t)
	}
	return tests, nil
}
