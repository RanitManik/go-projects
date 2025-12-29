# 06 - Email Verifier Tool

This Go program is a command-line tool that verifies email domains by checking their DNS records for MX, SPF, and DMARC configurations. It reads domains from standard input and outputs results in CSV format.

### 🔧 Imports

```go
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)
```

- `bufio`: For reading input line by line.
- `fmt`, `log`: For output and logging.
- `net`: For DNS lookups.
- `os`: For standard input.
- `strings`: For string manipulation.

### 🧩 Main Functionality

#### Domain Checking Function

```go
func checkDomain(domain string) {
	// ... variable declarations ...

	// MX lookup
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("MX lookup error for %s: %v\n", domain, err)
	} else if len(mxRecords) > 0 {
		hasMX = true
	}

	// SPF lookup
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

	// DMARC lookup
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

	// Output CSV
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
```

- Performs DNS lookups for MX records (mail servers).
- Checks TXT records for SPF policy.
- Checks _dmarc subdomain for DMARC policy.
- Outputs results in CSV format.

#### Main Function

```go
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
```

- Reads domains from standard input line by line.
- Prints CSV header.
- Calls checkDomain for each input line.

### 📝 Usage

- Pipe domains to the program: `cat domains.txt | go run main.go`
- Or input manually: `echo "example.com" | go run main.go`
- Outputs CSV with verification results.