# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#AccountConfiguration">AccountConfiguration</a>

Methods:

- <code title="get /accounts/{account_token}">client.Accounts.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountToken <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#AccountConfiguration">AccountConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_token}">client.Accounts.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#AccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountToken <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#AccountUpdateParams">AccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#AccountConfiguration">AccountConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CreditConfiguration

Response Types:

- <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#BusinessAccount">BusinessAccount</a>

Methods:

- <code title="get /accounts/{account_token}/credit_configuration">client.Accounts.CreditConfiguration.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#AccountCreditConfigurationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountToken <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#BusinessAccount">BusinessAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_token}/credit_configuration">client.Accounts.CreditConfiguration.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#AccountCreditConfigurationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountToken <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#AccountCreditConfigurationUpdateParams">AccountCreditConfigurationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#BusinessAccount">BusinessAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cards

Response Types:

- <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#CardProvisionResponse">CardProvisionResponse</a>

Methods:

- <code title="post /cards">client.Cards.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#CardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#CardNewParams">CardNewParams</a>) (<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cards/{card_token}">client.Cards.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#CardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardToken <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /cards/{card_token}">client.Cards.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#CardService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardToken <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#CardUpdateParams">CardUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cards/{card_token}/provision">client.Cards.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#CardService.Provision">Provision</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardToken <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#CardProvisionParams">CardProvisionParams</a>) (<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#CardProvisionResponse">CardProvisionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FinancialTransactions

Response Types:

- <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#FinancialTransaction">FinancialTransaction</a>

Methods:

- <code title="get /cards/{card_token}/financial_transactions/{financial_transaction_token}">client.Cards.FinancialTransactions.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#CardFinancialTransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardToken <a href="https://pkg.go.dev/builtin#string">string</a>, financialTransactionToken <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#FinancialTransaction">FinancialTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Status

Response Types:

- <a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#StatusListResponse">StatusListResponse</a>

Methods:

- <code title="get /status">client.Status.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#StatusService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/meorphis/test-repo-15/v2#StatusListResponse">StatusListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
