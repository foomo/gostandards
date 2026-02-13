//go:generate go run ../cmd/generate/main.go -package=iso4217 -output=code_gen.go
package iso4217

// Code is an ISO 4217 https://en.wikipedia.org/wiki/ISO_4217 currency code.
type Code string
