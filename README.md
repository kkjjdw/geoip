# geoip - fast geoip country library

[![GoDoc](https://godoc.org/github.com/phuslu/geoip?status.svg)](http://godoc.org/github.com/phuslu/geoip)

## Getting Started

```go
package main

import (
	"github.com/phuslu/geoip"
)

func main() {
	println(string(geoip.CountryByIPStr("1.1.1.1")))
}

// Output: US
```

## Benchmarks
```
BenchmarkGeoIpCountryByIPInt-8   	19.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGeoIpCountry-8          	25.5 ns/op	       0 B/op	       0 allocs/op
```