package scanner

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-cors/log"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/jpillora/go-tld"
)

// Scanner structure to hold the client connection we'll be making requests
// from as well as the scanner configuration and results from the scan
type Scanner struct {
	conf    *Conf
	l       *log.Logger
	mu      sync.Mutex
	Results map[string][][]*Test
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
	var r map[string][][]*Test
	r = make(map[string][][]*Test)
	s := Scanner{
		conf:    conf,
		l:       log,
		Results: r,
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
func (s *Scanner) Start() {
	// Log the current scanner configuration, now that it's been setup
	s.l.Log.Debug().Interface("conf", s.conf).Send()

	// Create a new sync.WaitGroup and the # of threads we need to it
	processGroup := new(sync.WaitGroup)
	processGroup.Add(s.conf.Threads)

	// For each thread, create a new client and begin performing the tests.
	for i := 0; i < s.conf.Threads; i++ {
		c := s.createClient()
		go func() {
			defer processGroup.Done()
			for _, t := range s.conf.Tests {
				s.runTests(c, t)
			}
		}()
	}
	processGroup.Wait()
}

// runTests runs all the tests and saves their results in array.
func (s *Scanner) runTests(c *http.Client, r *Request) {
	var t []*Test
	var err error

	t, err = s.reflectOrigin(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: reflect origins test failed - %s", err.Error())
	}

	t, err = s.httpOrigin(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: http origins test failed - %s", err.Error())
	}

	t, err = s.nullOrigin(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: null origins test failed - %s", err.Error())
	}

	t, err = s.wildcardOrigin(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: wildcard origins test failed - %s", err.Error())
	}

	t, err = s.thirdPartyOrigin(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: third party origins test failed - %s", err.Error())
	}

	t, err = s.backtickBypass(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: backtick bypass test failed - %s", err.Error())
	}

	t, err = s.preDomainBypass(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: prefix domain bypass test failed - %s", err.Error())
	}

	t, err = s.postDomainBypass(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: suffix domain bypass test failed - %s", err.Error())
	}

	t, err = s.underscoreBypass(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: underscore bypass test failed - %s", err.Error())
	}

	t, err = s.unescapedDotBypass(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: unescaped dot bypass test failed - %s", err.Error())
	}

	t, err = s.specialCharactersBypass(c, r, t)
	if err != nil {
		s.l.OutErr("s.runTests: special characters bypass test failed - %s", err.Error())
	}

	// Now that we've fnished running all the tests, let's all the results to our
	// Scanner's main results variable. First, let's parse the reqests URL
	url, _ := tld.Parse(r.URL)

	// Now, use the domain to index into our scanner Results map
	tests, ok := s.Results[url.Domain]
	if !ok {
		// If we were unable to retrieve anything from the map, then the entry doesn't
		// exist and we need to add it.
		var urlTests [][]*Test
		urlTests = append(urlTests, t)
		s.Results[url.Domain] = urlTests
		return
	}

	// Add this round of test results to the tests array, and update our map
	tests = append(tests, t)
	s.Results[url.Domain] = tests
}

// CreateOutputFile creates a new file and writes the test results to it
func (s *Scanner) CreateOutputFile(directory, domain string, results [][]*Test) error {
	newFile := s.newFileName(directory, domain)
	f, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		e := fmt.Sprintf("CreateOutputFile(%s): Error writing file (%s): %s\n", domain, newFile, err.Error())
		return errors.New(e)
	}
	err = ioutil.WriteFile(newFile, f, 0644)
	if err != nil {
		s.l.OutErr("s.CreateOutputFile: failed to save output file - %s", err.Error())
	}
	return nil
}

// newFileName creates the name for the output file in the format of domain_TIMESTAMP.json
func (s *Scanner) newFileName(directory, domain string) string {
	// Get the current time and parse it into a date and time string
	currTime := time.Now()
	cTimeArr := strings.Split(currTime.String(), " ")
	cDate := cTimeArr[0]
	cTime := cTimeArr[1]
	cTime = strings.Replace(cTime, ":", "-", 2)

	// Let's check the directory that we were given and ensure it has a "/" at the end
	if directory[len(directory)-1] != '/' {
		directory = directory + "/"
	}

	// Create the new file name and return it.
	newFile := directory + domain + "_" + cDate + "-" + cTime + ".json"
	return newFile
}
