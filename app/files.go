package app

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"go-cors/scanner"
	"go-cors/types"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// CreateOutputFile creates a new file and writes the test results to it
func (a *Application) CreateOutputFile(domain string, results []*types.Test) error {
	newFile := a.newFileName(domain)
	f, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		e := fmt.Sprintf("CreateOutputFile(%s): Error writing file (%s): %s\n", domain, newFile, err)
		return errors.New(e)
	}
	ioutil.WriteFile(newFile, f, 0644)
	return nil
}

// newFileName creates the name for the output file in the format of domain_TIMESTAMP.json
func (a *Application) newFileName(domain string) string {
	currTime := time.Now()
	cTimeArr := strings.Split(currTime.String(), " ")
	cDate := cTimeArr[0]
	cTime := cTimeArr[1]
	cTime = strings.Replace(cTime, ":", "-", 2)
	newFile := "/Users/sabra/go/src/go-cors/results/" + domain + "_" + cDate + "-" + cTime + ".json"
	return newFile
}

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
