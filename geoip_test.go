package geoip

import (
	"encoding/binary"
	"math/rand"
	"net"
	"testing"

	"github.com/oschwald/geoip2-golang"
)

func TestGeoIpCountry(t *testing.T) {
	cases := []struct {
		IP      string
		Country string
	}{
		{"0.1.1.1", ""},
		{"1.1.1.1", "AU"},
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
	}

	for _, c := range cases {
		country := Country(net.ParseIP(c.IP))
		// t.Logf("Country(%#v) return \"%s\", expect %#v", c.IP, country, c.Country)
		if string(country) != c.Country {
			t.Errorf("Country(%#v) return \"%s\", expect %#v", c.IP, country, c.Country)
		}
	}
}

func BenchmarkGeoIpCountryByIPInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountryByIPInt(rand.Uint32())
	}
}

func BenchmarkGeoIpCountry(b *testing.B) {
	ip := make([]byte, 4)
	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint32(ip[0:], rand.Uint32())
		Country(net.IP(ip))
	}
}

func BenchmarkMaxmindGolang(b *testing.B) {
	db, err := geoip2.Open(`/usr/local/share/GeoIP/GeoIP2-Country.mmdb`)
	if err != nil {
		b.Fatalf("geoip2.Open() error: %+v", err)
	}
	defer db.Close()

	b.ResetTimer()

	ip := make([]byte, 4)
	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint32(ip[0:], rand.Uint32())
		db.Country(net.IP(ip))
	}
}
