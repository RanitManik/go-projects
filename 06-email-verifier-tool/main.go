package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string) {
	if strings.TrimSpace(domain) == "" {
		return
	}

	var (
		hasMX       bool
		hasSPF      bool
		hasDMARC    bool
		spfRecord   string
		dmarcRecord string
	)

	// MX
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("MX lookup error for %s: %v\n", domain, err)
	} else if len(mxRecords) > 0 {
		hasMX = true
	}

	// SPF
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("TXT lookup error for %s: %v\n", domain, err)
	} else {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	}

	// DMARC
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("DMARC lookup error for %s: %v\n", domain, err)
	} else {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	}

	// CSV output
	fmt.Printf(
		"%s,%t,%t,%q,%t,%q\n",
		domain,
		hasMX,
		hasSPF,
		spfRecord,
		hasDMARC,
		dmarcRecord,
	)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("could not read from the input: %v\n", err)
	}
}
