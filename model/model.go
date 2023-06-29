package model

import "encoding/xml"

type ExchangeModelJson struct {
	BaseCurrency string `json:"base"`
	Date         string `json:"date"`
	Rates        struct {
		USD float64 `json:"USD"`
		EUR float64 `json:"EUR"`
		TRY float64 `json:"TRY"`
		GBP float64 `json:"GBP"`
	} `json:"rates"`
}

type ExchangeModelXML struct {
	XMLName  xml.Name `xml:"Tarih_Date"`
	Date     string   `xml:"Date,attr"`
	Currency []struct {
		CurrencyCode    string `xml:"CurrencyCode,attr"`
		Unit            string `xml:"Unit"`
		ForexBuying     string `xml:"ForexBuying"`
		ForexSelling    string `xml:"ForexSelling"`
		BanknoteBuying  string `xml:"BanknoteBuying"`
		BanknoteSelling string `xml:"BanknoteSelling"`
	} `xml:"Currency"`
}
type CurrencyType int

const (
	USD CurrencyType = iota
	EUR
	GBP

	// other currency types will be added
)

func (c CurrencyType) String() string {
	switch c {
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	case GBP:
		return "GBP"
	}
	return "Unknown Currency Type"
}
