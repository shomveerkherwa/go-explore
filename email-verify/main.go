package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the email to check domain, hasMX, hasSPF ,sprRecord, hasDMARC, dmarcRecord \n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v \n", err)
	}
}

func checkDomain(domain string) {
	// 1. Define all the variable needed
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// 1.1 try to lookuo mx records
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	// 2. set the hasMX variable
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// 3. look for txt records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	// 3.1 try to set the spfRecord and hasSPF variable
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v \n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	fmt.Printf("domain %v\n hasMX %v\n hasSPF %v\n spfRecord %v\n hasDMARC %v\n dmarcRecord %v\n",
		domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
