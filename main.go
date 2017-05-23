package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello this is the SSL checker app!")

	certPEM, err := ioutil.ReadFile("cert.pem")
	check(err)

	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		panic("Failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("Failed to parse certificate: " + err.Error())
	}
	fmt.Println(cert.PermittedDNSDomains)
}
