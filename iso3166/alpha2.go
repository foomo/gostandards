package iso3166

// Alpha2 is an ISO 3166-1 two-letter alphabetic country code.
//
//go:generate go run ../cmd/generate/main.go -package=iso3166 -type=alpha2 -output=alpha2_gen.go
type Alpha2 string
