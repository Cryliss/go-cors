package main

import (
	"github.com/Cryliss/gocors/log"
	"github.com/Cryliss/gocors/scanner"
)

// InitGoCors can be called from other applications to initialize a new scanner
func InitGoCors(output, timeout string, threads int) *scanner.Scanner {
	log := log.New()
	log.Verbose = false

	conf := scanner.Conf{
		Output:  output,
		Threads: threads,
		Timeout: timeout,
		Verbose: false,
	}
	scan := scanner.New(&conf, log)
	return scan
}
