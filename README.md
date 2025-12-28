[English version](README.md) | [中文版](README_CN.md)

[![Release](https://github.com/bynow2code/urlcheck/actions/workflows/release.yml/badge.svg)](https://github.com/bynow2code/rotail/actions/workflows/release.yml)

# urlcheck: Concurrent URL Availability Detection Tool

Quickly check the availability of batch URLs, supporting concurrency control, timeout settings, and CSV result export.

## Features
- Read URL list from command line or file
- Customizable concurrency count and timeout duration
- Output status code, response time, and failure reasons
- Support for exporting results to CSV files

## Usage Examples
# Check a single URL
```
./urlcheck https://google.cn
```

# Read from file, with 20 concurrent requests, 3 second timeout, export results
```
./urlcheck -f urls.txt -c 20 -t 3 -o result.csv
```

## Support

If you find this tool helpful, please give the project a star! ✨
