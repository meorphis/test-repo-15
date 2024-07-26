# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountConfiguration">AccountConfiguration</a>

Methods:

- <code title="get /accounts/{account_token}">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountToken <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountConfiguration">AccountConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_token}">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountToken <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountUpdateParams">AccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountConfiguration">AccountConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CreditConfiguration

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#BusinessAccount">BusinessAccount</a>

Methods:

- <code title="get /accounts/{account_token}/credit_configuration">client.Accounts.CreditConfiguration.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountCreditConfigurationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountToken <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#BusinessAccount">BusinessAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_token}/credit_configuration">client.Accounts.CreditConfiguration.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountCreditConfigurationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountToken <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountCreditConfigurationUpdateParams">AccountCreditConfigurationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#BusinessAccount">BusinessAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cards

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#CardProvisionResponse">CardProvisionResponse</a>

Methods:

- <code title="post /cards">client.Cards.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#CardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#CardNewParams">CardNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cards/{card_token}">client.Cards.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#CardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardToken <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /cards/{card_token}">client.Cards.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#CardService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardToken <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#CardUpdateParams">CardUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cards/{card_token}/provision">client.Cards.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#CardService.Provision">Provision</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardToken <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#CardProvisionParams">CardProvisionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#CardProvisionResponse">CardProvisionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FinancialTransactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#FinancialTransaction">FinancialTransaction</a>

Methods:

- <code title="get /cards/{card_token}/financial_transactions/{financial_transaction_token}">client.Cards.FinancialTransactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#CardFinancialTransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardToken <a href="https://pkg.go.dev/builtin#string">string</a>, financialTransactionToken <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#FinancialTransaction">FinancialTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Status

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#StatusListResponse">StatusListResponse</a>

Methods:

- <code title="get /status">client.Status.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#StatusService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#StatusListResponse">StatusListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
