package app

import (
	"go-cors/log"
	"go-cors/scanner"
)

// Application is our structure to hold data related to the program and its configuration.
type Application struct {
	domains []string
	log     *log.Logger
	Scan    *scanner.Scanner
	flags   *Flags
}
