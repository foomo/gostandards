package iso3166

// Alpha3 is an ISO 3166-1 three-letter alphabetic country code.
//
//go:generate go run ../cmd/generate/main.go -package=iso3166 -type=alpha3 -output=alpha3_gen.go
type Alpha3 string
