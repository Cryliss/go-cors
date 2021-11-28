// Package app ...
package app

import (
	"bufio"
	"encoding/json"
	"github.com/Cryliss/gocors/scanner"
	"io/ioutil"
	"os"
)

// parseDomainFile parses a .txt file that contains a list of domain names
// to run scans on
func (a *Application) parseDomainFile(f *os.File) error {
	var urls []string

	// Create a new bufio scanner so we can read line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Add each line to the URLs array
		urls = append(urls, scanner.Text())
	}
	// Set our app domains & return
	a.domains = urls
	return nil
}

// parseConfigFile parses a .json file that contains the configuration settings
// for the scans as well as a list of domains to perform the scans on.
func (a *Application) parseConfigFile(f *os.File) error {
	var conf scanner.Conf

	// Read the file as an of array of bytes
	byteData, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	// Unmarshal the byte data into our Conf struct
	json.Unmarshal(byteData, &conf)

	// Set the logger's verbose settings based on the config file
	a.log.Verbose = conf.Verbose

	// Create a new scanner, passing it the parsed configuration and our logger
	a.Scan = scanner.New(&conf, a.log)
	return nil
}
