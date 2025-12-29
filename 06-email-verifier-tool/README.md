# 06 - Email Verifier Tool

A command-line tool built with Go that verifies email domains by checking DNS records for MX, SPF, and DMARC configurations. It reads domains from standard input and outputs verification results in CSV format.

## Features

- Checks MX records (mail servers)
- Validates SPF policies
- Verifies DMARC configurations
- CSV output format
- Reads from stdin for batch processing

## Project Structure

- `main.go`: Main application logic for domain verification
- `domains.txt`: Sample list of domains to verify
- `go.mod`, `go.sum`: Go module files for dependency management

## How to Run

1. Navigate to the project directory:
   ```sh
   cd "06-email-verifier-tool"
   ```
2. Download dependencies:
   ```sh
   go mod tidy
   ```
3. Run the tool with domains:
   ```sh
   cat domains.txt | go run main.go
   ```
   Or manually:
   ```sh
   echo "example.com" | go run main.go
   ```

## Output Format

The tool outputs CSV with the following columns:
- `domain`: The domain name
- `hasMX`: Boolean indicating if MX records exist
- `hasSPF`: Boolean indicating if SPF record exists
- `spfRecord`: The SPF record content
- `hasDMARC`: Boolean indicating if DMARC record exists
- `dmarcRecord`: The DMARC record content

## Example Output

```
domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord
example.com,true,true,"v=spf1 -all",true,"v=DMARC1; p=reject"
```

---

This project is part of my Go learning journey!