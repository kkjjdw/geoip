# geoip - fast geoip country library

[![GoDoc](https://godoc.org/github.com/phuslu/geoip?status.svg)](http://godoc.org/github.com/phuslu/geoip)

### Getting Started

```go
package main

import (
	"net"
	"github.com/phuslu/geoip"
)

func main() {
	println(string(geoip.Country(net.ParseIP("1.1.1.1"))))
}

// Output: US
```

### Benchmarks
```
BenchmarkGeoIpCountryForIPv4-4          22.1 ns/op             0 B/op          0 allocs/op
BenchmarkGeoIpCountryForIPv6-8          70.8 ns/op             0 B/op          0 allocs/op
```

### Acknowledgment
This site or product includes IP2Location LITE data available from http://www.ip2location.com.
