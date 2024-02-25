[![Go Reference](https://pkg.go.dev/badge/github.com/zomasec/webarchive.svg)](https://pkg.go.dev/github.com/zomasec/webarchive) [![Go Report Card](https://goreportcard.com/badge/github.com/zomasec/webarchive)](https://goreportcard.com/report/github.com/zomasec/webarchive)
# WebArchive v1.0.2
Webarchive is a Go package for pentesters and developers to interacting with the Wayback Machine's CDX API and integrate web archive utilities into your Golang projects , Developed by Hazem El-Sayed (ZomaSec).

## Installation

To use `webarchive` in your Go project, simply run:

```bash
go get -u github.com/zomasec/webarchive
```
or install it manually using :
```bash
cd /usr/lib/go-*/src/
sudo git clone https://github.com/zomasec/webarchive
```
But if you use the manual installation then you will import it like :
```go
import "webarchive"
```

## Usage 

```go
package main

import (
	"fmt"
	"log"

	"github.com/zomasec/webarchive/"
)

func main() {
	// Create a new Archive instance
	archive := webarchive.NewArchive("example.com", nil)

	// Fetch historical URLs
	result, err := archive.FetchURLs()
	if err != nil {
		log.Fatal(err)
	}

	// Print the fetched URLs
	for _, u := range result.URLs {
		fmt.Println(u)
	}
}

```
## Filter URLs by Parameters
```go 
package main

import (
	"fmt"
	"log"

	"github.com/zomasec/webarchive"
)

func main() {
	// Create a new Archive instance
	archive := webarchive.NewArchive("example.com", nil)

	// Fetch historical URLs
	result, err := archive.FetchURLs()
	if err != nil {
		log.Fatal(err)
	}

	// Filter URLs with parameters
	params, err := result.HasParams()
	if err != nil {
		log.Fatal(err)
	}

	// Print URLs with parameters
	for _, u := range params {
		fmt.Println(u)
	}
}
```
## Filter URLs by Extension
```go
package main

import (
	"fmt"
	"log"

	"github.com/zomasec/webarchive"
)

func main() {
	// Create a new Archive instance
	archive := webarchive.NewArchive("example.com", nil)

	// Fetch historical URLs
	result, err := archive.FetchURLs()
	if err != nil {
		log.Fatal(err)
	}

	// Filter URLs by extension
	ext := ".html" // specify the desired extension
	filtered, err := result.FilterByExtension(ext)
	if err != nil {
		log.Fatal(err)
	}

	// Print URLs with the specified extension
	for _, u := range filtered {
		fmt.Println(u)
	}
}
```
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

