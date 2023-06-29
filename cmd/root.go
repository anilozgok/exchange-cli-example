/*
Copyright © 2023 Anil Can Ozgok
*/
package cmd

import (
	"encoding/xml"
	"exchange-cli-example/model"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exchange-cli-example",
	Short: "CLI tool for exchange ratings",
	Long: `
This basic CLI tool aims to bring latest exchange ratings based TRY for multiple currencies.

Supported currencies:
- USD => -u --usd for USD/TRY exchange rate
- EUR => -e --eur for EUR/TRY exchange rate
- GBP => -g --gbp for GBP/TRY exchange rate
Following currencies will be added soon.
- JPY
- DKK
- CHF
- SEK
- KWD
- NOK
- BGN
- RUB
- PKR
- QAR
- AED

NOTE: The exchange rates retrieved from https://www.tcmb.gov.tr/kurlar/today.xml
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		isUSD, _ := cmd.Flags().GetBool("usd")
		isEUR, _ := cmd.Flags().GetBool("eur")
		isGBP, _ := cmd.Flags().GetBool("gbp")

		if isUSD {
			getExchangeRate(model.USD.String())
		} else if isEUR {
			getExchangeRate(model.EUR.String())
		} else if isGBP {
			getExchangeRate(model.GBP.String())
		} else {
			panic("Unknown Currency Type")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.exchange-cli-example.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("usd", "u", false, "USD/TRY exchange rate")
	rootCmd.Flags().BoolP("eur", "e", false, "EUR/TRY exchange rate")
	rootCmd.Flags().BoolP("gbp", "g", false, "GBP/TRY exchange rate")
	rootCmd.Flags().BoolP("help", "h", false, "Help message")
}

func getExchangeRate(exchangeType string) {

	res, err := http.Get("https://www.tcmb.gov.tr/kurlar/today.xml")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var exchangeModelXML model.ExchangeModelXML
	err = xml.Unmarshal(body, &exchangeModelXML)
	if err != nil {
		panic(err)
	}

	writer(exchangeType, exchangeModelXML)
}

func writer(exchangeType string, exchangeModelXML model.ExchangeModelXML) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Currency Code", "Unit", "Forex Buying", "Forex Selling", "Banknote Buying", "Banknote Selling"})

	for item := range exchangeModelXML.Currency {
		if exchangeModelXML.Currency[item].CurrencyCode == strings.ToUpper(exchangeType) {
			currency := exchangeModelXML.Currency[item]
			currencyCode := currency.CurrencyCode + "/TRY"
			table.Append([]string{currencyCode, currency.Unit, currency.ForexBuying, currency.ForexSelling, currency.BanknoteBuying, currency.BanknoteSelling})
			table.SetFooter([]string{"", "", "", "", "Date", exchangeModelXML.Date})
			table.SetAutoMergeCells(true)
			table.SetRowLine(true)
			break
		}
	}

	table.Render()
}
