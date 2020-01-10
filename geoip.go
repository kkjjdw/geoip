// A pure Go interface to the free IP2Location GeoIP database.
//
// eg.
//
//      country := geoip.Country(net.ParseIP("1.1.1.1"))
//      fmt.Printf("%s\n", country)
package geoip

import (
	"encoding/binary"
	"net"
)

// Find country of IP.
func Country(ip net.IP) (country []byte) {
	if ip == nil {
		return
	}
	return country4(binary.BigEndian.Uint32(ip.To4()))
}

func country4(ipInt uint32) (country []byte) {
	i, j := 0, len(ips)
	for i < j {
		h := int(uint(i+j) >> 1)
		if ips[h] <= ipInt {
			i = h + 1
		} else {
			j = h
		}
	}

	if i == 0 {
		return
	}

	country = geo[i*2-2 : i*2]

	return
}
