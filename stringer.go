package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	address := ""
	for i,oct := range ip {
		if i < len(ip)-1 {
		    address += fmt.Sprintf("%d", oct) + "."
		} else {
		    address += fmt.Sprintf("%d", oct)
		}
		
	}
	return address
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
