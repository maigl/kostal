package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	resp, err := http.Get("https://coolors.co/palettes/trending")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read: %v\n", err)
		os.Exit(1)
	}

	re := regexp.MustCompile(`page_data_encoded\s*=\s*"([^"]+)"`)
	matches := re.FindSubmatch(body)
	if len(matches) < 2 {
		fmt.Fprintf(os.Stderr, "Could not find page_data_encoded\n")
		os.Exit(1)
	}

	encoded := matches[1]
	fmt.Fprintf(os.Stderr, "Encoded length: %d\n", len(encoded))
	fmt.Fprintf(os.Stderr, "First 100 chars: %s\n", string(encoded[:min(100, len(encoded))]))
	
	// Try URL-safe base64
	decoded, err := base64.URLEncoding.DecodeString(string(encoded))
	if err != nil {
		fmt.Fprintf(os.Stderr, "URL base64 failed: %v\n", err)
		// Try raw
		decoded, err = base64.RawURLEncoding.DecodeString(string(encoded))
		if err != nil {
			fmt.Fprintf(os.Stderr, "RawURL base64 failed: %v\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("%s\n", decoded)
}
