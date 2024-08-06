# Currency Converter with TUI

## Overview

This project is a Currency Converter application written in Go. It features a Text User Interface (TUI) to provide a user-friendly way to convert between different currencies. The application fetches real-time currency exchange rates from a public API.

![App's Demo](/demo.png "App's Demo")

## Packages Used

- **net/http**: Used for making HTTP requests to the currency exchange API.
- **github.com/charmbracelet/huh**: Utilized for building the TUI interface.
- **encoding/json**: Used to marshal and unmarshal JSON data from the API.

## Technical Considerations

### Currency API

To obtain exchange data, this project uses the API provided by [Open Exchange Rates](https://openexchangerates.org). The API offers a free tier, but it requires signing up for an API key.

### Conversion Logic

The free tier of the Open Exchange Rates API provides conversion rates with USD as the base currency. However, you can use these rates to convert between any two currencies by following these steps:

1. Fetch the exchange rate for the source currency (currencyFrom) and the target currency (currencyTo).
2. Use the following formula to convert the amount:

```go
rateFrom := data.Rates[currencyFrom]
rateTo := data.Rates[currencyTo]

converted := amount * (rateTo / rateFrom)
```

### Supported Currencies

For the initial implementation, the application supports the following currencies:

- USD (United States Dollar)
- EUR (Euro)
- GBP (British Pound Sterling)
- JPY (Japanese Yen)
- INR (Indian Rupee)

## Setup and Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/currency-converter.git
   cd currency-converter
   ```

2. **Install dependencies:**
   Make sure you have Go installed on your system. Then run:
   ```sh
   go get ./...
   ```

3. **Set up your API key:**
   - Rename the `.env.example` file to `.env`:
     ```sh
     mv .env.example .env
     ```
   - Open the `.env` file and add your API key

4. **Run the application:**
   ```sh
   go run .
   ```

## Usage

The application will launch with a TUI, allowing you to select the source and target currencies, and input the amount to be converted. The converted amount will be displayed on the interface.

## Contribution

Feel free to fork this repository and contribute by submitting a pull request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License.


## Acknowledgements

This project idea was inspired by the [Currency Converter project](https://github.com/dreamsofcode-io/goprojects/tree/main/05-currency-converter) by [dreamsofcode-io](https://github.com/dreamsofcode-io).
