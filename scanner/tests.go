package scanner

import (
	"net/http"

	tld "github.com/jpillora/go-tld"
)

func (s *Scanner) reflectOrigin(c *http.Client, r *Request) error {
	s.l.Out("Starting reflect origins test")
	origin := "https://crylis.io/"
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return err
	}

	if acao == origin {
		s.l.OutErr("s.reflectOrigin: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
	}
	return nil
}

func (s *Scanner) httpOrigin(c *http.Client, r *Request) error {
	s.l.Out("Starting http origins test")
	origin := "http://crylis.io/"
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return err
	}

	if acao == origin {
		s.l.OutErr("s.httpOrigin: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
	}
	return nil
}

func (s *Scanner) nullOrigin(c *http.Client, r *Request) error {
	s.l.Out("Starting null origins test")
	origin := "null"
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return err
	}

	if acao == origin {
		s.l.OutErr("s.nullOrigin: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
	}
	return nil
}

func (s *Scanner) wildcardOrigin(c *http.Client, r *Request) error {
	s.l.Out("Starting wilcard origin test")
	origin := "*"
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return err
	}

	if acao == origin {
		s.l.OutErr("s.wildcardOrigin: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
	}
	return nil
}

func (s *Scanner) thirdPartyOrigin(c *http.Client, r *Request) error {
	s.l.Out("Starting third party test")
	origins := []string{
		"http://jsbin.com",
		"https://codepen.io",
		"https://jsfiddle.net",
		"https://www.webdevout.net",
		"https://repl.it",
	}
	for _, origin := range origins {
		acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
		if err != nil {
			return err
		}

		if acao == origin {
			s.l.OutErr("s.thirdpartyOrigin: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
		}
	}
	return nil
}

func (s *Scanner) backtickBypass(c *http.Client, r *Request) error {
	s.l.Out("Starting backtick bypass test")
	url, err := tld.Parse(r.URL)
	origin := "https://" + url.Subdomain + "." + url.Domain + "." + url.TLD + "`.cryls.io"
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return err
	}

	if acao == origin {
		s.l.OutErr("s.backtickBypass: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
	}
	return nil
}

func (s *Scanner) preDomainBypass(c *http.Client, r *Request) error {
	s.l.Out("Starting predomain bypass test")
	url, err := tld.Parse(r.URL)
	if err != nil {
		return err
	}
	origin := "https://" + url.Domain + ".cryls.io"
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return err
	}

	if acao == origin {
		s.l.OutErr("s.preDomainBypass: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
	}
	return nil
}

func (s *Scanner) postDomainBypass(c *http.Client, r *Request) error {
	s.l.Out("Starting postdomain bypass test")
	url, err := tld.Parse(r.URL)
	if err != nil {
		return err
	}

	origin := "https://crylis" + url.Domain + "." + url.TLD
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return err
	}

	if acao == origin {
		s.l.OutErr("s.postDomainBypass: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
		return nil
	}

	origin = "https://crylis.io" + url.Domain + "." + url.TLD
	acao, acac, err = s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return err
	}

	if acao == origin {
		s.l.OutErr("s.postDomainBypass: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
	}
	return nil
}

func (s *Scanner) underscoreBypass(c *http.Client, r *Request) error {
	s.l.Out("Starting underscore bypass test")
	url, err := tld.Parse(r.URL)
	if err != nil {
		return err
	}
	origin := "https://" + url.Subdomain + "." + url.Domain + "." + url.TLD + "_.cryls.io"
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return err
	}

	if acao == origin {
		s.l.OutErr("s.underscoreBypass: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
	}
	return nil
}

func (s *Scanner) unescapedDotBypass(c *http.Client, r *Request) error {
	s.l.Out("Starting unescaped dot bypass test")
	url, err := tld.Parse(r.URL)
	if err != nil {
		return err
	}
	origin := "https://" + url.Subdomain + "S" + url.Domain + "." + url.TLD
	acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
	if err != nil {
		return err
	}

	if acao == origin {
		s.l.OutErr("s.unescapedBypass: Misconfiguration found for %s! acao: %s & acac: %s", r.URL, acac, acao)
	}
	return nil
}

func (s *Scanner) specialCharactersBypass(c *http.Client, r *Request) error {
	s.l.Out("Starting special characters bypass test")
	url, err := tld.Parse(r.URL)
	if err != nil {
		return err
	}
	specialChars := []string{"_", "-", `"`, "{", "}", "+", "^", "%60", "!", "~", ";", "|", "&", "'", "(", ")", "*", ",", "$", "=", "+", "%0b"}
	for _, char := range specialChars {
		origin := "https://" + url.Subdomain + "." + url.Domain + "." + url.TLD + char + ".crylis.io"
		acao, acac, err := s.sendRequest(c, r.URL, origin, r.Method, r.Headers)
		if err != nil {
			return err
		}

		if acao == origin {
			s.l.OutErr("s.specialCharactersBypass: Misconfiguration found for %s using special character %s! acao: %s & acac: %s", r.URL, char, acac, acao)
		}
	}
	return nil
}
