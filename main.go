package main

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func readSites() ([]string, error) {
	var sites []string
	file, err := os.Open("./sites.txt")
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
	err_msg := color.New(color.Bold, color.FgRed)
	success := color.New(color.Bold, color.FgGreen)

	sites_to_check, err := readSites()
	if err != nil {
		err_msg.Println("Could not open file", err)
		os.Exit(1)
	}

	if len(sites_to_check) < 1 {
		err_msg.Println("No sites found, try adding some")
	}

	for _, site := range sites_to_check {
		resp, err := http.Get(site)
		if err != nil {
			log.Fatal(err_msg.Println("x may be down", err))
		}
		if resp.StatusCode == 200 {
			success.Printf("✓ (%d) %s\n", resp.StatusCode, site)
		}
		if resp.StatusCode != 200 {
			err_msg.Printf("x (%d) %s\n", resp.StatusCode, site)
		}
	}
}
