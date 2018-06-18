package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var iStr string
	for i := range ip {
		if iStr != "" {
			iStr += "."
		}
		iStr += fmt.Sprint(ip[i])
	}
	return iStr
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
