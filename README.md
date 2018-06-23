# geoip - fast geoip country library

[![Build Status](https://travis-ci.org/phuslu/geoip.svg)](https://travis-ci.org/phuslu/geoip)
[![GoDoc](https://godoc.org/github.com/phuslu/geoip?status.svg)](http://godoc.org/github.com/phuslu/geoip)

# Benchmarks
```
BenchmarkGeoIpCountryByIPInt-2   	10000000	       190 ns/op
BenchmarkGeoIpCountry-2          	10000000	       200 ns/op
BenchmarkMaxmindGolang-2         	  200000	     12289 ns/op
```