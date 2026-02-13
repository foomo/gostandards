# Getting Started

## Installation

```bash
go get github.com/foomo/gostandards
```

## Available Packages

| Package | Import | Description |
| --- | --- | --- |
| [ISO 3166](/standards/iso3166) | `github.com/foomo/gostandards/iso3166` | Country codes (Alpha-2 and Alpha-3) |
| [ISO 4217](/standards/iso4217) | `github.com/foomo/gostandards/iso4217` | Currency codes |
| [HTTP](/standards/http) | `github.com/foomo/gostandards/http` | Headers, encodings, and auth prefixes |

## Quick Example

```go
package main

import (
	"fmt"

	"github.com/foomo/gostandards/iso3166"
	"github.com/foomo/gostandards/iso4217"
)

func main() {
	country := iso3166.Alpha2US
	fmt.Println(country)          // "US"
	fmt.Println(country.Valid())  // true

	currency := iso4217.CodeEUR
	fmt.Println(currency)         // "EUR"
	fmt.Println(currency.Valid()) // true
}
```
