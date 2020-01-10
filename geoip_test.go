package geoip

import (
	"encoding/binary"
	"math/rand"
	"net"
	"testing"
	"time"
)

func TestGeoIpCountry(t *testing.T) {
	cases := []struct {
		IP      string
		Country string
	}{
		{"0.1.1.1", "ZZ"},
		{"1.1.1.1", "US"},
		{"121.229.143.64", "CN"},
		{"122.96.43.186", "CN"},
		{"123.249.20.198", "CN"},
		{"153.3.123.160", "CN"},
		{"153.3.131.201", "CN"},
		{"180.109.81.198", "CN"},
		{"180.111.103.88", "CN"},
		{"183.206.11.225", "CN"},
		{"192.210.171.249", "US"},
		{"223.112.9.2", "CN"},
		{"23.16.28.232", "CA"},
		{"58.240.115.210", "CN"},
		{"61.155.4.66", "CN"},
		{"255.255.255.255", "ZZ"},
		{"::", "ZZ"},
		{"::1", "ZZ"},
	}

	for _, c := range cases {
		country := Country(net.ParseIP(c.IP))
		if string(country) != c.Country {
			t.Errorf("Country(%#v) return \"%s\", expect %#v", c.IP, country, c.Country)
		}
	}
}

func BenchmarkGeoIpCountryForIPv4(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	ip := make([]byte, 4)
	binary.LittleEndian.PutUint32(ip[0:], rand.Uint32())

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Country(net.IP(ip))
	}
}
