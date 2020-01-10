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

type uint128 struct {
	H, L uint64
}

// Find country of IP.
func Country(ip net.IP) (country []byte) {
	if ip == nil {
		return
	}
	if ip4 := ip.To4(); ip4 != nil {
		country = country4(binary.BigEndian.Uint32(ip4))
	} else {
		ip6 := uint128{binary.BigEndian.Uint64(ip), binary.BigEndian.Uint64(ip[8:])}
		country = country6(ip6)
	}
	return
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

func country6(ipInt uint128) (country []byte) {
	country = []byte{'Z', 'Z'}
	return
}
