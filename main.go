package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

/*
readSites reads a file containing a list of sites
and returns the list as a slice of strings.

It takes a file name as a string parameter
*/
func readSites(fileName string) ([]string, error) {
	var sites []string
	file, err := os.Open(fileName)
	if err != nil {
		return sites, err
	}
	defer file.Close()

	fScanner := bufio.NewScanner(file)
	for fScanner.Scan() {
		sites = append(sites, fScanner.Text())
	}

	return sites, nil
}

func main() {
	var (
		filePath string
		errMsg   = color.New(color.Bold, color.FgRed)
		success  = color.New(color.Bold, color.FgGreen)
	)

	flag.StringVar(&filePath, "f", "./sites.txt", "The file location (full or relative path)\nExample: ~/Desktop/some_file.txt")
	flag.Parse()

	sitesToCheck, err := readSites(filePath)
	if err != nil {
		errMsg.Println("Could not open file", err)
		os.Exit(1)
	}

	if len(sitesToCheck) < 1 {
		errMsg.Println("No sites found, try adding some\nUse --help to see how to do this")
	}

	for _, site := range sitesToCheck {
		resp, err := http.Get(site)
		if err != nil {
			log.Fatal(errMsg.Println("x may be down", err))
		}
		if resp.StatusCode == 200 {
			success.Printf("✓ (%d) %s\n", resp.StatusCode, site)
		}
		if resp.StatusCode != 200 {
			errMsg.Printf("x (%d) %s\n", resp.StatusCode, site)
		}
	}
}
