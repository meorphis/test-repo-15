# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAccountGetResponse">AccountAccountGetResponse</a>

Methods:

- <code title="get /account">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountService.AccountGet">AccountGet</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAccountGetParams">AccountAccountGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAccountGetResponse">AccountAccountGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Addresses

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAddressUpdateResponse">AccountAddressUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAddressAccountAddressNewResponse">AccountAddressAccountAddressNewResponse</a>

Methods:

- <code title="put /account/addresses/{id}">client.Accounts.Addresses.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAddressService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAddressUpdateParams">AccountAddressUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAddressUpdateResponse">AccountAddressUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /account/addresses/{id}">client.Accounts.Addresses.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAddressService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAddressDeleteParams">AccountAddressDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /account/addresses">client.Accounts.Addresses.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAddressService.AccountAddressNew">AccountAddressNew</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAddressAccountAddressNewParams">AccountAddressAccountAddressNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountAddressAccountAddressNewResponse">AccountAddressAccountAddressNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Exists

Methods:

- <code title="get /account/exists">client.Accounts.Exists.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountExistService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountExistListParams">AccountExistListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## PaymentMethods

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountPaymentMethodAccountAddPaymentMethodResponse">AccountPaymentMethodAccountAddPaymentMethodResponse</a>

Methods:

- <code title="delete /account/payment-methods/{id}">client.Accounts.PaymentMethods.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountPaymentMethodService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountPaymentMethodDeleteParams">AccountPaymentMethodDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /account/payment-methods">client.Accounts.PaymentMethods.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountPaymentMethodService.AccountAddPaymentMethod">AccountAddPaymentMethod</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountPaymentMethodAccountAddPaymentMethodParams">AccountPaymentMethodAccountAddPaymentMethodParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#AccountPaymentMethodAccountAddPaymentMethodResponse">AccountPaymentMethodAccountAddPaymentMethodResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#PaymentNewResponse">PaymentNewResponse</a>

Methods:

- <code title="post /payments">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#PaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#PaymentNewParams">PaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#PaymentNewResponse">PaymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Guests

## Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#GuestPaymentGuestPaymentsInitializeResponse">GuestPaymentGuestPaymentsInitializeResponse</a>

Methods:

- <code title="post /guest/payments">client.Guests.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#GuestPaymentService.GuestPaymentsInitialize">GuestPaymentsInitialize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#GuestPaymentGuestPaymentsInitializeParams">GuestPaymentGuestPaymentsInitializeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#GuestPaymentGuestPaymentsInitializeResponse">GuestPaymentGuestPaymentsInitializeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Merchants

## Callbacks

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantCallbackListResponse">MerchantCallbackListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantCallbackMerchantCallbacksUpdateResponse">MerchantCallbackMerchantCallbacksUpdateResponse</a>

Methods:

- <code title="get /merchant/callbacks">client.Merchants.Callbacks.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantCallbackService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantCallbackListParams">MerchantCallbackListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantCallbackListResponse">MerchantCallbackListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /merchant/callbacks">client.Merchants.Callbacks.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantCallbackService.MerchantCallbacksUpdate">MerchantCallbacksUpdate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantCallbackMerchantCallbacksUpdateParams">MerchantCallbackMerchantCallbacksUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantCallbackMerchantCallbacksUpdateResponse">MerchantCallbackMerchantCallbacksUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Identifiers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantIdentifierListResponse">MerchantIdentifierListResponse</a>

Methods:

- <code title="get /merchant/identifiers">client.Merchants.Identifiers.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantIdentifierService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#MerchantIdentifierListResponse">MerchantIdentifierListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookNewResponse">WebhookNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookGetResponse">WebhookGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookListResponse">WebhookListResponse</a>

Methods:

- <code title="put /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookNewParams">WebhookNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookNewResponse">WebhookNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks/{id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookGetResponse">WebhookGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookListParams">WebhookListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/{id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Testings

## Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#TestingAccountTestingAccountNewResponse">TestingAccountTestingAccountNewResponse</a>

Methods:

- <code title="post /testing/accounts">client.Testings.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#TestingAccountService.TestingAccountNew">TestingAccountNew</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#TestingAccountTestingAccountNewParams">TestingAccountTestingAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#TestingAccountTestingAccountNewResponse">TestingAccountTestingAccountNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CreditCards

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#TestingCreditCardListResponse">TestingCreditCardListResponse</a>

Methods:

- <code title="get /testing/credit-cards">client.Testings.CreditCards.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#TestingCreditCardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#TestingCreditCardListResponse">TestingCreditCardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Shipments

Methods:

- <code title="post /testing/shipments">client.Testings.Shipments.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#TestingShipmentService.TestingShipmentTrackingNew">TestingShipmentTrackingNew</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go">meorphistest40</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/meorphis-test-40-go#TestingShipmentTestingShipmentTrackingNewParams">TestingShipmentTestingShipmentTrackingNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
