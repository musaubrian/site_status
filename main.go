package main

import (
	"github.com/fatih/color"
	"net/http"
)

func main() {
	sites_to_check := []string{"https://musaubrian.vercel.app", "https://openjournal.vercel.app", "https://github.com/musuabrian"}
	err_msg := color.New(color.Bold, color.FgRed)
	success := color.New(color.Bold, color.FgGreen)

	for _, site := range sites_to_check {
		resp, err := http.Get(site)
		if err != nil {
			err_msg.Printf("x %s may be down\n", err)
		}
		if resp.StatusCode != 200 {
			err_msg.Printf("x (%d) %s not okay\n", resp.StatusCode, site)
		}
		success.Printf("✓ (%d) %s\n", resp.StatusCode, site)
	}
}
