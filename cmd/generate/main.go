package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"text/template"
)

// srcDir returns the directory containing this source file.
func srcDir() string {
	_, f, _, _ := runtime.Caller(0)
	return filepath.Dir(f)
}

func main() {
	var pkg, output, typ string
	flag.StringVar(&pkg, "package", os.Getenv("GOPACKAGE"), "target package")
	flag.StringVar(&output, "output", "", "output file (relative)")
	flag.StringVar(&typ, "type", "", "type to generate (e.g. alpha2, alpha3)")
	flag.Parse()

	if output == "" {
		output = filepath.Join(".", fmt.Sprintf("code_%s.go", pkg))
	}

	switch pkg {
	case "iso3166":
		if typ == "" {
			panic("iso3166: -type flag is required (alpha2 or alpha3)")
		}

		countries, err := fetchCountries("https://raw.githubusercontent.com/lukes/ISO-3166-Countries-with-Regional-Codes/refs/heads/master/all/all.csv")
		if err != nil {
			panic(err)
		}

		tmplFile := fmt.Sprintf("iso3166_%s.gotmpl", typ)
		data := map[string]any{
			"generator": "from CSV",
			"countries": countries,
		}

		if err := generate(tmplFile, output, data); err != nil {
			panic(err)
		}

		log.Printf("iso3166/%s: generated %d countries", typ, len(countries))

	case "iso4217":
		currencies, err := fetchCurrencies("https://raw.githubusercontent.com/datasets/currency-codes/refs/heads/main/data/codes-all.csv")
		if err != nil {
			panic(err)
		}

		data := map[string]any{
			"generator":  "from CSV",
			"currencies": currencies,
		}

		if err := generate("iso4217.gotmpl", output, data); err != nil {
			panic(err)
		}

		log.Printf("iso4217: generated %d currencies", len(currencies))

	default:
		panic(fmt.Sprintf("unsupported package: %s", pkg))
	}

	log.Printf("Generated %s for package %s", output, pkg)
}

// generate parses a template from srcDir and executes it into outfile.
func generate(tmplFile, outfile string, data any) error {
	tmpl, err := template.ParseFiles(filepath.Join(srcDir(), tmplFile))
	if err != nil {
		return err
	}

	w, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer w.Close()

	return tmpl.Execute(w, data)
}

// fetchCSV fetches a CSV from url and returns rows as []map[string]string.
func fetchCSV(url string) ([]map[string]string, error) {
	resp, err := http.Get(url) //nolint:gosec,noctx
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := csv.NewReader(resp.Body)

	headers, err := decoder.Read()
	if err != nil {
		return nil, err
	}

	var records []map[string]string

	for {
		row, err := decoder.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		record := make(map[string]string, len(headers))
		for i, value := range row {
			record[headers[i]] = value
		}

		records = append(records, record)
	}

	return records, nil
}

type tplCountry struct {
	Alpha2  string
	Alpha3  string
	Numeric string
	Name    string
}

func fetchCountries(url string) ([]tplCountry, error) {
	records, err := fetchCSV(url)
	if err != nil {
		return nil, err
	}

	var countries []tplCountry

	for _, r := range records {
		if r["alpha-3"] == "" {
			continue
		}

		countries = append(countries, tplCountry{
			Name:    r["name"],
			Alpha2:  r["alpha-2"],
			Alpha3:  r["alpha-3"],
			Numeric: r["country-code"],
		})
	}

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Alpha2 < countries[j].Alpha2
	})

	return countries, nil
}

type tplCurrency struct {
	Code string
	Name string
}

func fetchCurrencies(url string) ([]tplCurrency, error) {
	records, err := fetchCSV(url)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]bool)

	var currencies []tplCurrency

	for _, r := range records {
		code := r["AlphabeticCode"]
		if code == "" {
			continue
		}

		// skip withdrawn currencies
		if r["WithdrawalDate"] != "" {
			continue
		}

		if seen[code] {
			continue
		}

		seen[code] = true

		currencies = append(currencies, tplCurrency{
			Code: code,
			Name: r["Currency"],
		})
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Code < currencies[j].Code
	})

	return currencies, nil
}
