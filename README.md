[![Build Status](https://github.com/foomo/gostandards/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/foomo/gostandards/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/gostandards)](https://goreportcard.com/report/github.com/foomo/gostandards)
[![GoDoc](https://godoc.org/github.com/foomo/gostandards?status.svg)](https://godoc.org/github.com/foomo/gostandards)

<p align="center">
  <img alt="gostandards" src="docs/public/logo.png" height="400" width="400"/>
</p>

# Go Standards

> Collection of standards as types

## Standards

- [ISO 3166-1](https://en.wikipedia.org/wiki/ISO_3166-1) - Country codes (Alpha-2 & Alpha-3)
- [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) - Currency codes
- HTTP headers, content encodings & authorization prefixes

## Usage

```go
import (
	"github.com/foomo/gostandards/iso3166"
	"github.com/foomo/gostandards/iso4217"
	gohttp "github.com/foomo/gostandards/http"
)

// ISO 3166-1 country codes
country := iso3166.Alpha2US   // "US"
alpha3 := iso3166.Alpha3USA   // "USA"

// ISO 4217 currency codes
currency := iso4217.CodeUSD   // "USD"

// HTTP constants
header := gohttp.HeaderContentType  // "Content-Type"
```

## How to Contribute

Contributions are welcome! Please read the [contributing guide](docs/CONTRIBUTING.md).

![Contributors](https://contributors-table.vercel.app/image?repo=foomo/gostandards&width=50&columns=15)

## License

Distributed under MIT License, please read the [license file](LICENSE) for more details.

_Made with ♥ [foomo](https://www.foomo.org) by [bestbytes](https://www.bestbytes.com)_
