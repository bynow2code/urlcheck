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

### Check a single URL

```
./urlcheck https://google.cn
```

### Read from file, with 20 concurrent requests, 3 second timeout, export results

```
./urlcheck -f urls.txt -c 20 -t 3 -o result.csv
```

## Quick Installation

### Method 1: One-click Installation Script (Recommended)

For Linux and macOS systems:

```
curl -sfL https://raw.githubusercontent.com/bynow2code/urlcheck/main/install.sh | bash
```

### Method 2: Manual Download

Visit the [GitHub Releases](https://github.com/bynow2code/urlcheck/releases/latest) page to download the pre-compiled
version suitable for your system:

| System  | Architecture | Filename                     |
|---------|--------------|------------------------------|
| Windows | AMD64        | `urlcheck-windows-amd64.exe` |
| macOS   | AMD64        | `urlcheck-darwin-amd64`      |
| macOS   | ARM64        | `urlcheck-darwin-arm64`      |
| Linux   | AMD64        | `urlcheck-linux-amd64`       |
| Linux   | ARM64        | `urlcheck-linux-arm64`       |

After downloading, Windows users need to rename the file to `urlcheck.exe`, while other systems should add execute
permissions as needed.

### Verify Installation

After installation, run the following command to verify that the installation was successful:

```
urlcheck -v
```

## Support

If you find this tool helpful, please give the project a star! ✨
