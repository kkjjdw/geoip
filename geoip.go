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
	if ip4 := ip.To4(); ip4 != nil {
		country = country4(binary.BigEndian.Uint32(ip4))
	} else {
		country = country6(binary.BigEndian.Uint64(ip), binary.BigEndian.Uint64(ip[8:]))
	}
	return
}

func country4(n uint32) (country []byte) {
	i, j := 0, len(ips)
	for i < j {
		h := int(uint(i+j) >> 1)
		if ips[h] <= n {
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

func country6(high, low uint64) (country []byte) {
	i, j := 0, len(ips6)/16
	for i < j {
		h := int(uint(i+j) >> 1)
		hi := binary.BigEndian.Uint64(ips6[h*16:])
		lo := binary.BigEndian.Uint64(ips6[h*16+8:])
		if hi > high || (hi == high && lo > low) {
			j = h
		} else {
			i = h + 1
		}
	}

	if i == 0 {
		return
	}

	country = geo6[i*2-2 : i*2]

	return
}
