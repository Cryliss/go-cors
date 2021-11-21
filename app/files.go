package app

import (
	"bufio"
	"encoding/json"
	"go-cors/scanner"
	"io/ioutil"
	"os"
)

func (a *Application) parseDomainFile(f *os.File) error {
	var urls []string
	// Create a new bufio scanner so we can read line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	a.domains = urls
	return nil
}

func (a *Application) parseConfigFile(f *os.File) error {
	var conf scanner.Conf
	byteData, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	json.Unmarshal(byteData, &conf)
	a.log.Verbose = conf.Verbose
	a.Scan = scanner.New(&conf, a.log)
	return nil
}
