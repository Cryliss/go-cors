// Package app ...
package app

import (
	"github.com/Cryliss/gocors/log"
	"github.com/Cryliss/gocors/scanner"
)

// New creates and returns a new Application with a default scanner initialized
func New() *Application {
	var domains []string
	a := Application{
		domains: domains,
		log:     log.New(),
	}
	return &a
}

// CheckFlags checks the input flags that were given and sets up the application accordingly
func (a *Application) CheckFlags(flags *Flags) error {
	isConfig, err := flags.checkInputFile(a)
	if err != nil {
		a.log.OutErr("a.Start: failed to parse input flags")
		a.log.OutErr("%s", err.Error())
		return err
	}

	// If config is true, then that means the input file was the configuration file,
	// so we don't need to continue checking the other input arguments.
	if isConfig {
		return nil
	}

	// Set the logs verbose value
	a.log.Verbose = flags.Verbose

	// Define the scanner configuration based on the given flags & create a new scanner
	conf := scanner.Conf{
		Output:  flags.Output,
		Threads: flags.Threads,
		Timeout: flags.Timeout,
		Verbose: flags.Verbose,
	}
	a.Scan = scanner.New(&conf, a.log)

	// Okay, so let's check if we added any domains while checking the input file
	// if we didn't, then the input flag wasn't set so let's check the URL given
	if len(a.domains) == 0 {
		flags.checkURL(a)
	}

	// Get the header flags
	headers := flags.checkHeader(a)

	// The default method for each request is 'GET'
	method := "GET"
	if flags.validateMethod(a) {
		// If we were given a valid request method in our flags, let's override the default
		method = flags.Method
	}

	// Create the test structs for scanner to use during testing
	a.Scan.CreateTests(a.domains, headers, method, flags.Proxy)
	return nil
}
