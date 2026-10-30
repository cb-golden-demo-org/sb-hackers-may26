package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

// Deliberately insecure fixture for Policy Engine UAT test 7a-2 (Gosec rule G402, TLS
// certificate verification disabled). Standard library only, so it needs no module downloads.
// Exists only on branch pe-uat-7a2-gosec-standalone-module. Safe to delete.
func main() {
	fmt.Println(insecureClient() != nil)
}

func insecureClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}
