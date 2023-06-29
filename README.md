This basic CLI tool aims to bring latest exchange ratings based TRY for multiple currencies.

Supported currencies:

- USD => ``-u --usd`` for USD/TRY exchange rate
- EUR => ``-e --eur`` for EUR/TRY exchange rate
- GBP => ``-g --gbp`` for GBP/TRY exchange rate

Example Usages:

- ``go run maing.go -u``

![u.png](resources%2Fu.png)

- ``go run main.go --eur``

![eur.png](resources%2Feur.png)

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