package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

func init() {
	LoadENV()
}

var currencySymbol = map[string]string{
	"USD": "$",
	"INR": "₹",
	"EUR": "€",
	"GBP": "£",
	"JPY": "¥",
}

var currencyOptions = []huh.Option[string]{
	huh.NewOption("EUR", "EUR"),
	huh.NewOption("INR", "INR"),
	huh.NewOption("USD", "USD"),
	huh.NewOption("GBP", "GBP"),
	huh.NewOption("JPY", "JPY"),
}

func main() {
	var currencyFrom string
	var currencyTo string
	var amount string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a base currency").
				Options(currencyOptions...).
				Value(&currencyFrom),
			huh.NewSelect[string]().
				Title("Select a currency you want to convert into").
				Options(currencyOptions...).
				Value(&currencyTo),
			huh.NewInput().
				Title("How much to convert?").
				Value(&amount),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	if currencyFrom == "" {
		log.Fatalln("form: you didn't input from currency!")
	}

	if currencyTo == "" {
		log.Fatalln("form: you didn't input to currency!")
	}

	if amount == "" {
		log.Fatalln("form: You didn't input amount!")
	}

	data := FetchConversion()

	rateFrom := data.Rates[currencyFrom]
	rateTo := data.Rates[currencyTo]

	amt, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		log.Fatalln("form: error while parsing amount of form")
	}

	converted := amt * (rateTo / rateFrom)

	fmt.Printf("%s%.2f converts to %s%.2f", currencySymbol[currencyFrom], amt, currencySymbol[currencyTo], converted)
}
