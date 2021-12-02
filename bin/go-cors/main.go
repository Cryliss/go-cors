package main

import (
	"flag"
	"fmt"
	"github.com/Cryliss/gocors/app"
	"os"
	"time"
)

var (
	url     string
	header  string
	method  string
	file    string
	threads int
	output  string
	timeout string
	proxy   string
	verbose bool
)

// args parses the input arguments
func args() {
	flag.StringVar(&url, "url", "", "The URL to scan for CORS misconfiguration")
	flag.StringVar(&header, "headers", "", "The headers to include in the request ")
	flag.StringVar(&method, "method", "GET", "Include another method other than GET")
	flag.StringVar(&file, "input", "", "A text file with a list of domains to scan for CORS misconfiguration")
	flag.IntVar(&threads, "threads", 10, "Number of threads to use for the scan.")
	flag.StringVar(&output, "output", "", "Directory to save the results to a JSON file. Always saves the file as domain_TIMESTAMP.json")
	flag.StringVar(&timeout, "timeout", "10s", "Set the request timeout")
	flag.StringVar(&proxy, "proxy", "", "Use a HTTP address as a proxy")
	flag.BoolVar(&verbose, "verbose", false, "Enable the UI for realtime results and requests")
	flag.Parse()
}

func main() {
	// Parse input flags (arguments)
	args()

	f := app.Flags{
		URL:     url,
		Header:  header,
		Method:  method,
		File:    file,
		Threads: threads,
		Output:  output,
		Timeout: timeout,
		Proxy:   proxy,
		Verbose: verbose,
	}

	// Create a new app
	a := app.New()

	// Validate the given input arguments
	if err := a.CheckFlags(&f); err != nil {
		os.Exit(-1)
	}

	start := time.Now()
	// We got valid flags / configuration files, let's start the scanning process
	a.Scan.Start()
	elapsed := time.Since(start)

	var tests int
	for _, r := range a.Scan.Results {
		for _, t := range r {
			tests = tests+len(t)
		}
	}
	fmt.Printf("\nRan %d tests in %v\n", tests,elapsed)
}
