#!/usr/bin/env python3
# pylint: disable=too-many-statements, line-too-long, W0703

__version__ = '1.0'

import os
import struct


def download():
    """download ip2location country csv"""
    assert os.system('rm -rf IP2LOCATION-LITE-DB1.CSV') == 0
    assert os.system('curl http://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.CSV.ZIP >IP2LOCATION-LITE-DB1.CSV.ZIP') == 0, 'download ip2location database error'
    assert os.system('unzip IP2LOCATION-LITE-DB1.CSV.ZIP') == 0, 'unzip ip2location database error'
    assert os.system('rm -rf IP2LOCATION-LITE-DB1.CSV.ZIP LICENSE-CC-BY-SA-4.0.TXT README_LITE.TXT') == 0
    return 'IP2LOCATION-LITE-DB1.CSV'


def generate(filename):
    """generate geoip_db.go"""
    with open('geoip_db.go', 'wb') as file:
        file.write(b'package geoip\n\nvar db = []byte("')
        for line in open(filename):
            parts = line.strip().split(',')
            ip_str = ''.join('\\x%02x' % x for x in struct.pack('>I', int(parts[0].strip('"'))))
            country = parts[2].strip('"')
            if country == '-':
                country = 'ZZ'
            file.write((ip_str + country).encode())
        file.write(b'")\n')


if __name__ == '__main__':
    generate(download())
