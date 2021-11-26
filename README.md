# go-cors
A tool for scanning domains for CORS misconfigurations written in Go.  
Final project for COMP 424 Software Security  
Professor: Dr. Wonju Lee

**By:**  
Sabra Bilodeau  
Sally Chung

## Misconfigurations Tested
`go-cors` tests the follow CORS misconfigurations:  

- [Backtick Bypass](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/BACKTICK_BYPASS.md)
- [HTTP Origin](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/HTTP_ORIGIN.md)
- [Origin Reflection](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/ORIGIN_REFLECTION.md)
- [Null Origin](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/NULL_ORIGIN.md)
- [Post-Domain Bypass](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/POSTDOMAIN_BYPASS.md)
- [Pre-Domain Bypass](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/PREDOMAIN_BYPASS.md)
- [Special Characters Bypass](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/SPECIAL_CHARACTERS_BYPASS.md)
- [Third Party Origin](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/THIRD_PARTY_ORIGINS.md)
- [Underscore Bypass](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/UNDERSCORE_BYPASS.md)
- [Unescaped Dot Bypass](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/UNESCAPED_DOT_BYPASS.md)
- [Wildcard Origin](https://github.com/Cryliss/go-cors/blob/main/docs/misconfigurations/WILDCARD_ORIGIN.md)

For more information on each, including sample exploits and possible fixes for the vulnerabilities, please click the link provided.

## Installation
Clone the repository:  
`git clone https://github.com/Cryliss/go-cors.git`  

Change directories to the repository's directory:  
`cd go-cors`  

Build the application:  
`make build`  

## Usage
### Simple Scans
To run a scan on a signle URL, use `./go-cors -url https://example.com`.  

To run scans on multiple URLs, save the URLs to a `.txt` file and run the program like so:  

`./go-cors -input global_top_100_domains.txt`  

### Configurable Scans
To add additional configuration to a request, there are two options.  
1. Add any of the following command line flags to your input  
2. Update the provided `conf.json` to reflect your desired configuration.   

### CLI flags
| Flag | Description | Default |
| :--: | :---------- | :-----: |
| -url     | The URL to scan for CORS misconfiguration | "" |
| -headers | Include headers | "" |
| -method  |  Include another method other than `GET` | "GET" |
| -input   |  A text file with a list of domains or a json configuration file | "" |
| -threads |  Number of threads to use for the scan | 10 |
| -output  |  Save the results to a JSON file | true |
| -timeout |  Set requests timeout | "10s" |
| -proxy   |  Use a proxy (HTTP) | "" |
| -h       |  Show the help information & exit | N/A |
| -verbose |  Enables the UI to display realtime results | false |

## Example Usage of the CLI flags  
- **URL**:     `./go-cors -url https://example.com`
- **Headers**: `./go-cors -url https://example.com -headers "User-Agent: GoogleBot\nCookie: SESSION=Hacked"`
- **Method**:  `./go-cors -url https://example.com -method POST`
- **Input**:   `./go-cors -input global_top_100_domains.txt`
- **Threads**: `./go-cors -url https://example.com -threads 20`
- **Output**:  `./go-cors -url https://example.com -output true`
- **Timeout**: `./go-cors -url https://example.com -timeout 20s`
- **Proxy**:   `./go-cors -url https://example.com -proxy http://127.0.0.1:4545`
- **Verbose**: `./go-cors -url https://example.com -verbose true`


# Using `go-cors` in your own application

```go
package main

import (
    "fmt"
    "github.com/Cryliss/go-cors/scanner"
    "github.com/Cryliss/go-cors/log"
)

// initGoCors initializes a new go-cors scanner
func initGoCors() *scanner.Scanner {    
    log := log.New()
    log.Verbose = false

    conf := scanner.Conf{
        Output: true,
        Threads: 10,
        Timeout: "10s",
        Verbose: false,
    }
    scan := scanner.New(&conf, log)
    return scan
}

func main() {
    corsScanner := initGoCors()

    /*
    In order to start running tests with go-cors, we need to create them first.

    Creating tests requires an array of domain names, a scanner.Headers variable 
    which is a map[string]string of header name-value pairs, a request method and 
    a proxy URL.

    After creating our headers variable and domain names, then we can call the create 
    tests function, which will set scanner.Conf.Tests value at the end.
    */
    var headers scanner.Headers
    domains := []string{"https://www.instagram.com/"}
    corsScanner.CreateTests(domains, headers, "GET", "")

    // Now that we have our tests set, we can go ahead and start the scanner.
    corsScanner.Start()

    // After running the scan, you can save your results to a specific file directory like so:
    if err := corsScanner.SaveResults("/your/directories/filepath/"); err != nil {
        fmt.Printf("Error saving reults: %s\n", err.Error())
    }
}
```
