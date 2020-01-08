// A pure Go interface to the free MaxMind GeoIP database.
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
	return CountryByIPInt(binary.BigEndian.Uint32(ip.To4()))
}

// Find country of IP string
func CountryByIPStr(ipStr string) (country []byte) {
	return CountryByIPInt(newIPInt(ipStr))
}

// Find country of IP uint32 value
func CountryByIPInt(ipInt uint32) (country []byte) {
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

func newIPInt(ipStr string) uint32 {
	var dots int
	var i, j uint32
	for _, c := range ipStr {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			j = j*10 + uint32(c-'0')
		case '.':
			if j >= 256 {
				return 0
			}
			i = i*256 + j
			j = 0
			dots++
		default:
			return 0
		}
	}
	if dots != 3 || j >= 256 {
		return 0
	}
	return i*256 + j
}

