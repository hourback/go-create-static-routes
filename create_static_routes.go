package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Great!  Looks like we're running on Windows.  Commencing to improve routing table. . . .")
	} else {
		fmt.Println("This program only runs on Windows for now.  Exiting. . . .")
		os.Exit(1)
	}

	var domains []string

	if len(os.Args) == 1 {
		domains = []string{"pastebin.com", "archive.ubuntu.com"}
	} else {
		domains = os.Args[1:]
	}

	for i := range domains {
		domain := domains[i]
		fmt.Println("\n\nWorking on domain ", domain)
		// Do a single lookup
		/*		ip, err := net.ResolveIPAddr("ip4", domain)
				if err == nil {
					fmt.Printf("%s ok %s\n", domain, ip)
				} else {
					fmt.Printf("%s error: %s\n", domain, err)
				}*/

		// Use LookupIP instead to get an array of IPs
		ips, err := net.LookupIP(domain)
		if err == nil {
			for i := range ips {
				ip := ips[i]
				cmd := exec.Command("cmd", "/C", "route", "add", ip.String(), "mask", "255.255.255.255", "172.16.1.1")
				var out bytes.Buffer
				cmd.Stdout = &out
				err := cmd.Run()
				if err != nil {
					log.Println("\n\nCommand failed:", out.String())
					log.Fatal(err)
				}
				fmt.Println("Added route for ", ip.String())
			}
		} else {
			fmt.Printf("\n\n%s error: %s\n", domain, err)
		}
	}
}
