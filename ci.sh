#!/bin/bash -xe

if [ ${#GITHUB_TOKEN} -eq 0 ]; then
	echo "WARNING: \$GITHUB_TOKEN is not set!"
fi

function build() {
	cd /tmp
	rm -rf GeoIPCountryCSV.zip
	curl -O http://www.maxmind.com/download/geoip/database/GeoIPCountryCSV.zip
	yes | unzip GeoIPCountryCSV.zip
	git clone https://${GITHUB_TOKEN}@github.com/phuslu/geoip

	cd geoip

	python << END
import csv
import sys

fp = open('geoip_ranges.go', 'wb')
fp.write(b'package geoip\n\nvar ranges = []byte("')
for row in csv.reader(open('../GeoIPCountryWhois.csv')):
    fp.write((''.join('\\\\x%02x' % int(x) for x in row[0].split('.')) + row[-2]).encode())
fp.write(b'")\n')

END

	go test -v .
}

function release() {
	git diff --quiet && exit 0
	git add * && \
	git commit -m "[skip ci] update geoip ranges" && \
	git push origin master -f
}

build
release
