package scanner

import (
	"go-cors/log"
	"go-cors/types"
	"net/http"
	"sync"
)

// Make a new request (uses a client, method, url, origins and header )
// URL parser
// Function to create an array of origins

// Scanner structure to hold the client connection we'll be making requests
// from as well as the scanner configuration and results from the scan
type Scanner struct {
	conf    *Conf
	l       *log.Logger
	Results []*types.Test
}

// Conf structure to hold configuration settings for the scna
type Conf struct {
	Output  bool       `json:"output"`
	Tests   []*Request `json:"tests"`
	Threads int        `json:"threads"`
	Timeout string     `json:"timeout"`
	Verbose bool       `json:"verbose"`
}

// Request structure to hold configuration settings for each request
type Request struct {
	URL     string  `json:"url"`
	Headers Headers `json:"headers"`
	Method  string  `json:"method"`
	Proxy   string  `json:"proxy"`
}

// Headers is a map of our header key,value pairs
type Headers map[string]string

// New creates and returns a new scanner
func New(conf *Conf, log *log.Logger) *Scanner {
	s := Scanner{
		conf: conf,
		l:    log,
	}
	return &s
}

// CreateTests initializes our array of tests
func (s *Scanner) CreateTests(domains []string, headers Headers, method, proxy string) {
	var tests []*Request
	for _, domain := range domains {
		r := Request{
			URL:     domain,
			Headers: headers,
			Method:  method,
			Proxy:   proxy,
		}
		tests = append(tests, &r)
	}
	s.conf.Tests = tests
}

// Start begins the scanning procedure
func (s *Scanner) Start(a types.Application) {
	// Log the current scanner configuration, now that it's been setup
	s.l.Log.Debug().Interface("conf", s.conf).Send()

	processGroup := new(sync.WaitGroup)
	processGroup.Add(s.conf.Threads)

	for i := 0; i < s.conf.Threads; i++ {
		c := s.createClient()
		go func() {
			defer processGroup.Done()
			for _, t := range s.conf.Tests {
				s.runTests(a, c, t)
			}
		}()
	}
	processGroup.Wait()
}

func (s *Scanner) runTests(a types.Application, c *http.Client, r *Request) {
	var t []*types.Test
	if err := s.reflectOrigin(c, r, t); err != nil {
		s.l.OutErr("s.runTests: reflect origins test failed - %s", err.Error())
	}
	if err := s.httpOrigin(c, r, t); err != nil {
		s.l.OutErr("s.runTests: http origins test failed - %s", err.Error())
	}
	if err := s.nullOrigin(c, r, t); err != nil {
		s.l.OutErr("s.runTests: null origins test failed - %s", err.Error())
	}
	if err := s.wildcardOrigin(c, r, t); err != nil {
		s.l.OutErr("s.runTests: wildcard origins test failed - %s", err.Error())
	}
	if err := s.thirdPartyOrigin(c, r, t); err != nil {
		s.l.OutErr("s.runTests: third party origins test failed - %s", err.Error())
	}
	if err := s.backtickBypass(c, r, t); err != nil {
		s.l.OutErr("s.runTests: backtick bypass test failed - %s", err.Error())
	}
	if err := s.preDomainBypass(c, r, t); err != nil {
		s.l.OutErr("s.runTests: prefix domain bypass test failed - %s", err.Error())
	}
	if err := s.postDomainBypass(c, r, t); err != nil {
		s.l.OutErr("s.runTests: suffix domain bypass test failed - %s", err.Error())
	}
	if err := s.underscoreBypass(c, r, t); err != nil {
		s.l.OutErr("s.runTests: underscore bypass test failed - %s", err.Error())
	}
	if err := s.unescapedDotBypass(c, r, t); err != nil {
		s.l.OutErr("s.runTests: unescaped dot bypass test failed - %s", err.Error())
	}
	if err := s.specialCharactersBypass(c, r, t); err != nil {
		s.l.OutErr("s.runTests: special characters bypass test failed - %s", err.Error())
	}
	// Saving to output does not work yet :(
	//a.CreateOutputFile(r.URL, t)
}
