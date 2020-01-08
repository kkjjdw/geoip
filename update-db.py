#!/usr/bin/env python3
# pylint: disable=too-many-statements, line-too-long, W0703

import io
import sys
import urllib.request
import zipfile


def main():
    """convert ip2location country csv to geoip_db.go"""
    filename = 'http://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.CSV.ZIP'
    if len(sys.argv) > 1:
        filename = sys.argv[1]
    if filename.startswith(('http://', 'https://')):
        file = io.BytesIO(urllib.request.urlopen(filename).read())
    else:
        file = open(filename, 'rb')
    if filename.lower().endswith('.zip'):
        text = zipfile.ZipFile(file).read('IP2LOCATION-LITE-DB1.CSV')
    else:
        text = file.read()
    # generate geoip_db.go
    ips, geo = [], []
    for line in io.BytesIO(text):
        parts = line.strip().decode().split(',')
        ip = parts[0].strip('"')
        country = parts[2].strip('"')
        if country == '-':
            country = 'ZZ'
        ips.append(ip)
        geo.append(country)
    with open('geoip_db.go', 'wb') as file:
        file.write(('''package geoip

var ips = []uint32{%s}
var geo = []byte("%s")
''' % (','.join(ips), ''.join(geo))).encode())


if __name__ == '__main__':
    main()
