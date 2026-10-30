package main

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"net/http"
)

// Deliberately insecure fixture for Policy Engine UAT tests 7a-2 and 7a-3. Standard library only.
// Contains: G402 (TLS verification disabled), G501/G401 (md5 import and use), G101 (hard-coded
// credential; the value below is a made-up fake, not a real secret).
// Exists only on branches pe-uat-7a2-... and pe-uat-7a3-... Safe to delete.

const fixturePassword = "Zq8!rT4#vN7@pL2xW9" // fake value for the Gosec G101 test

func main() {
	fmt.Println(insecureClient() != nil)
	fmt.Printf("%x\n", md5.Sum([]byte(fixturePassword)))
}

func insecureClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}
