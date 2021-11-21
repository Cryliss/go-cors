package scanner

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

func (s *Scanner) newRequest(method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		s.l.Log.Err(err).Msg("s.addRequestMethod")
	}
	return req
}

func (s *Scanner) createClient() *http.Client {
	timeout, err := time.ParseDuration(s.conf.Timeout)
	if err != nil {
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

func (s *Scanner) sendRequest(c *http.Client, url, origin, method string, headers Headers) (string, string, error) {
	req := s.newRequest(method, url)
	req.Header.Set("Origin", origin)
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	resp, err := c.Do(req)
	if err != nil {
		s.l.OutErr("s.sendRequest: failed to send request to %s - %+v", url, err)
		return "", "", err
	}

	acao := resp.Header.Get("Access-Control-Allow-Origin")
	acac := resp.Header.Get("Access-Control-Allow-Credentials")

	return acao, acac, nil
}
