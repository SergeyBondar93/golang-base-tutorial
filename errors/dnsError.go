package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("https://google.com")

	if err != nil {
		var dnsErr *net.DNSError

		if errors.As(err, &dnsErr) {

			if dnsErr.Timeout() {
				fmt.Println("Timeout error", dnsErr)
			}
			if dnsErr.Temporary() {
				fmt.Println("Temporary error", dnsErr)
			}

			fmt.Println("DNS GENERIC error", err)
			return
		}
		fmt.Println("GENERIC error", err)
		return
	}

	fmt.Println(addr)

}
