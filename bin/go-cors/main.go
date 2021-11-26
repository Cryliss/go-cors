package main

import (
	"flag"
	"go-cors/app"
	"os"
)

var (
	url     string
	header  string
	method  string
	file    string
	threads int
	output  bool
	timeout string
	proxy   string
	verbose bool
)

// args parses the input arguments
func args() {
	flag.StringVar(&url, "url", "", "The URL to scan for CORS misconfiguration")
	flag.StringVar(&header, "headers", "", "The headers to include in the request ")
	flag.StringVar(&method, "method", "GET", "Include another method other than `GET`")
	flag.StringVar(&file, "input", "", "A text file with a list of domains to scan for CORS misconfiguration")
	flag.IntVar(&threads, "threads", 10, "Number of threads to use for the scan. Default: 40")
	flag.BoolVar(&output, "output", true, "Save the results to a JSON file. Always saves as go-cors/results/domain_TIMESTAMP.json. Default: true")
	flag.StringVar(&timeout, "Timeout", "10s", "Set the request timeout. Default: 10s")
	flag.StringVar(&proxy, "proxy", "", "Use a HTTP address as a proxy")
	flag.BoolVar(&verbose, "verbose", false, "Enable the UI for realtime results and requests. Defualt: false")
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

	// We got valid flags / configuration files, let's start the scanning process
	a.Scan.Start()

	dir := "/Users/sabra/go/src/go-cors/results/"
	results := a.Scan.Results
	for key, r := range results {
		a.Scan.CreateOutputFile(dir, key, r)
	}
}
