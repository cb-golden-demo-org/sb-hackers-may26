package main

import (
	"crypto/tls"
	"net/http"
)

// peUAT7aInsecureClient is a deliberately insecure fixture for Policy Engine UAT test 7a
// (Gosec rule G402, TLS certificate verification disabled). It is not called anywhere.
// It exists only on branch pe-uat-7a-gosec. Safe to delete.
func peUAT7aInsecureClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}
