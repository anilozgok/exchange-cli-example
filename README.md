This basic CLI tool aims to bring latest exchange ratings based TRY for multiple currencies.

Supported currencies:

- USD => ``-u --usd`` for USD/TRY exchange rate
- EUR => ``-e --eur`` for EUR/TRY exchange rate
- GBP => ``-g --gbp`` for GBP/TRY exchange rate

Example Usages:

- ``go run maing.go -u``

![Screenshot_2.png](..%2F..%2FDesktop%2FScreenshot_2.png)

- ``go run main.go --eur``

![Screenshot_1.png](..%2F..%2FDesktop%2FScreenshot_1.png)

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