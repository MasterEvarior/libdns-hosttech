**DEVELOPER INSTRUCTIONS:**

This repo is a template for developers to use when creating new [libdns](https://github.com/libdns/libdns) provider implementations.

Be sure to update:

- ~~The package name~~
- ~~The Go module name in go.mod~~
- ~~The latest `libdns/libdns` version in go.mod~~~~
- All comments and documentation, including README below and godocs
- ~~License (must be compatible with Apache/MIT)~~
- All "TODO:"s is in the code
- All methods that currently do nothing

Remove this section from the readme before publishing.

---

hosttech for [`libdns`](https://github.com/libdns/libdns)
=======================

[![Go Reference](https://pkg.go.dev/badge/test.svg)](https://pkg.go.dev/github.com/libdns/hosttech)

This package implements the [libdns interfaces](https://github.com/libdns/libdns) for [hosttech.ch](https://hosttech.ch), allowing you to manage DNS records.

## Example Use

```go

```

## Constraints
Because the Hosttech API does not provide a way to manipulate a generic "Type,Name,Value"-Record, not every type of record can be set. Currently supported are:
- AAAA
- A
- NS
- CNAME
- MX
- TXT
- TLSA

Any unsupported record types returns an error.

## Further documentation
Any further documentation that could be helpful:
 - [Hosttech DNS API documentation](https://api.ns1.hosttech.eu/api/documentation)


