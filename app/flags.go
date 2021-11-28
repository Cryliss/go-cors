package app

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Cryliss/gocors/scanner"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Flags holds our command line argument flag values for later parsing
type Flags struct {
	// The URL to scan for CORS misconfiguration
	URL string
	// The headers to include in the request
	Header string
	// Additional methods
	Method string
	// A text file with a list of domains to scan for CORS misconfiguration
	File string
	// Number of threads to use for the scan
	Threads int
	// Save the results to a JSON file. Always saves as gocors/results/domain_TIMESTAMP.json
	Output string
	// Set requests timeout
	Timeout string
	// Use a HTTP address as a proxy
	Proxy string
	// Enable the UI for realtime results and requests
	Verbose bool
}

// usage prints out how to use the program
func usage() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(-1)
}

// checkInputFile checks the given input file, determines if it's a .json or .txt file
// and passes it off the proper parsing function.
func (flags *Flags) checkInputFile(a *Application) (bool, error) {
	isConfig := false

	if flags.File != "" {
		// Open the file
		f, err := os.OpenFile(flags.File, os.O_RDWR, 0644)
		if err != nil {
			e := fmt.Sprintf("a.flags.CheckInputFile: failed to open input file - %s", err.Error())
			return isConfig, errors.New(e)
		}
		defer f.Close()

		// Get the files extension type
		ext := filepath.Ext(flags.File)
		switch ext {
		case ".txt":
			if err := a.parseDomainFile(f); err != nil {
				e := fmt.Sprintf("a.flags.CheckInputFile: failed to parse input file - %s", err.Error())
				return isConfig, errors.New(e)
			}
			return isConfig, nil
		case ".json":
			isConfig = true
			if err := a.parseConfigFile(f); err != nil {
				e := fmt.Sprintf("a.flags.CheckInputFile: failed to parse config file - %s", err.Error())
				return isConfig, errors.New(e)
			}
			return isConfig, nil
		default:
			e := fmt.Sprintf("CheckFlags(%+v): Cannot parse input file - invalid file type", flags)
			return isConfig, errors.New(e)
		}
	}
	return isConfig, nil
}

// checkURL checks the URL flag
func (flags *Flags) checkURL(a *Application) {
	// We only get here if we didn't have any type of input file, so we must get a URL flag
	// so if we don't, we'll display the usage and exit
	if flags.URL == "" {
		usage()
	}
	a.domains = append(a.domains, flags.URL)
}

var allowedHeaders = scanner.Headers{
	"access-control-allow-origin":      "",
	"access-control-allow-credentials": "",
	"access-control-allow-headers":     "",
	"access-control-allow-methods":     "",
	"access-control-expose-headers":    "",
	"access-control-max-Age":           "",
	"access-control-request-headers":   "",
	"access-control-request-method":    "",
	"origin":                           "",
	"from":                             "",
	"host":                             "",
	"referer":                          "",
	"referer-policy":                   "",
	"user-agent":                       "",
	"cookie":                           "",
}

// checkHeader checks the header flag
func (flags *Flags) checkHeader(a *Application) scanner.Headers {
	h := make(scanner.Headers)

	if flags.Header != "" {
		// Our headers are in the format "header-name: header-value", so we're gonna
		// use regular expressions to help us parse the headers
		r := regexp.MustCompile(`(.*):\s(.*)`)

		// Since a user can add multiple header values, in the format
		// "header-name: header-value\nheader-name: header-value",
		// we're going to split the input string at the carriage return character
		headers := strings.Split(flags.Header, "\n")
		for _, header := range headers {
			/*  Example usage of FindStringSubmatch

			header := "User-Agent: GoogleBot"
			pairs := r.FindStringSubmatch(header)
			pairs would be: [User-Agent: GoogleBot User-Agent GoogleBot]

			If no pairs are found, matches would nil
			*/
			matches := r.FindStringSubmatch(header)
			if matches == nil {
				a.log.OutErr("a.flags.CheckHeader: failed to find a correctly formatted header from the string: %s", header)
				a.log.Out("a.flags.CheckHeader: continuing to check for additional headers")
				continue
			}

			// If the provided header value is one of the allowed headers
			if _, ok := allowedHeaders[strings.ToLower(matches[1])]; ok {
				// Add it to our headers map
				h[matches[1]] = matches[2]
			}
		}
	}
	return h
}

var allowedMethods = map[string]string{
	"DELETE": "",
	"GET":    "",
	"PUT":    "",
	"POST":   "",
	"HEAD":   "",
}

// validateMethod ensure we were given a valid request method from the flags
func (flags *Flags) validateMethod(a *Application) bool {
	if _, ok := allowedMethods[flags.Method]; ok {
		return true
	}
	a.log.OutErr("a.flags.validateMethod: ignoring given method %s, only DELETE, GET, PUT, POST, and HEAD methods are allowed", flags.Method)
	return false
}
