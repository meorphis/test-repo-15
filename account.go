// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"context"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/tidwall/gjson"
)

// AccountService contains methods and other services that help with interacting
// with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options        []option.RequestOption
	Addresses      *AccountAddressService
	Exists         *AccountExistService
	PaymentMethods *AccountPaymentMethodService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Addresses = NewAccountAddressService(opts...)
	r.Exists = NewAccountExistService(opts...)
	r.PaymentMethods = NewAccountPaymentMethodService(opts...)
	return
}

// Retrieve a shopper's account details, such as addresses and payment information
func (r *AccountService) AccountGet(ctx context.Context, query AccountAccountGetParams, opts ...option.RequestOption) (res *AccountAccountGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccountGetResponse struct {
	Addresses      []AccountAccountGetResponseAddress       `json:"addresses,required"`
	PaymentMethods []AccountAccountGetResponsePaymentMethod `json:"payment_methods,required"`
	Profile        AccountAccountGetResponseProfile         `json:"profile"`
	JSON           accountAccountGetResponseJSON            `json:"-"`
}

// accountAccountGetResponseJSON contains the JSON metadata for the struct
// [AccountAccountGetResponse]
type accountAccountGetResponseJSON struct {
	Addresses      apijson.Field
	PaymentMethods apijson.Field
	Profile        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccountGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccountGetResponseAddress struct {
	ID             string                               `json:"id,required"`
	CountryCode    string                               `json:"country_code,required"`
	FirstName      string                               `json:"first_name,required"`
	LastName       string                               `json:"last_name,required"`
	Locality       string                               `json:"locality,required"`
	PostalCode     string                               `json:"postal_code,required"`
	StreetAddress1 string                               `json:"street_address1,required"`
	Company        string                               `json:"company"`
	Email          string                               `json:"email" format:"email"`
	IsDefault      bool                                 `json:"is_default"`
	Phone          string                               `json:"phone" format:"phone"`
	Region         string                               `json:"region"`
	StreetAddress2 string                               `json:"street_address2"`
	JSON           accountAccountGetResponseAddressJSON `json:"-"`
}

// accountAccountGetResponseAddressJSON contains the JSON metadata for the struct
// [AccountAccountGetResponseAddress]
type accountAccountGetResponseAddressJSON struct {
	ID             apijson.Field
	CountryCode    apijson.Field
	FirstName      apijson.Field
	LastName       apijson.Field
	Locality       apijson.Field
	PostalCode     apijson.Field
	StreetAddress1 apijson.Field
	Company        apijson.Field
	Email          apijson.Field
	IsDefault      apijson.Field
	Phone          apijson.Field
	Region         apijson.Field
	StreetAddress2 apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccountGetResponseAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccountGetResponseAddressJSON) RawJSON() string {
	return r.raw
}

type AccountAccountGetResponsePaymentMethod struct {
	Tag AccountAccountGetResponsePaymentMethodsTag `json:".tag,required"`
	// The expiration date of the credit card. TODO TO MAKE EXPIRATION REUSABLE
	Expiration string `json:"expiration,required" format:"^\\d{4}-\\d{2}$"`
	// The last 4 digits of the credit card number.
	Last4 string `json:"last4,required" format:"^\\d{4}$"`
	// The credit card network.
	Network AccountAccountGetResponsePaymentMethodsNetwork `json:"network,required"`
	ID      string                                         `json:"id"`
	// The ID of credit card's billing address
	BillingAddressID string                                     `json:"billing_address_id"`
	JSON             accountAccountGetResponsePaymentMethodJSON `json:"-"`
}

// accountAccountGetResponsePaymentMethodJSON contains the JSON metadata for the
// struct [AccountAccountGetResponsePaymentMethod]
type accountAccountGetResponsePaymentMethodJSON struct {
	Tag              apijson.Field
	Expiration       apijson.Field
	Last4            apijson.Field
	Network          apijson.Field
	ID               apijson.Field
	BillingAddressID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAccountGetResponsePaymentMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccountGetResponsePaymentMethodJSON) RawJSON() string {
	return r.raw
}

type AccountAccountGetResponsePaymentMethodsTag string

const (
	AccountAccountGetResponsePaymentMethodsTagCreditCard AccountAccountGetResponsePaymentMethodsTag = "credit_card"
)

func (r AccountAccountGetResponsePaymentMethodsTag) IsKnown() bool {
	switch r {
	case AccountAccountGetResponsePaymentMethodsTagCreditCard:
		return true
	}
	return false
}

// The credit card's billing address
type AccountAccountGetResponsePaymentMethodsBillingAddress struct {
	// The type of address reference
	Tag AccountAccountGetResponsePaymentMethodsBillingAddressTag `json:".tag,required"`
	// The address's ID
	ID             string                                                    `json:"id,required"`
	FirstName      string                                                    `json:"first_name"`
	LastName       string                                                    `json:"last_name"`
	Company        string                                                    `json:"company"`
	StreetAddress1 string                                                    `json:"street_address1"`
	StreetAddress2 string                                                    `json:"street_address2"`
	Locality       string                                                    `json:"locality"`
	PostalCode     string                                                    `json:"postal_code"`
	Region         string                                                    `json:"region"`
	CountryCode    string                                                    `json:"country_code"`
	Email          string                                                    `json:"email" format:"email"`
	Phone          string                                                    `json:"phone" format:"phone"`
	JSON           accountAccountGetResponsePaymentMethodsBillingAddressJSON `json:"-"`
	union          AccountAccountGetResponsePaymentMethodsBillingAddressUnion
}

// accountAccountGetResponsePaymentMethodsBillingAddressJSON contains the JSON
// metadata for the struct [AccountAccountGetResponsePaymentMethodsBillingAddress]
type accountAccountGetResponsePaymentMethodsBillingAddressJSON struct {
	Tag            apijson.Field
	ID             apijson.Field
	FirstName      apijson.Field
	LastName       apijson.Field
	Company        apijson.Field
	StreetAddress1 apijson.Field
	StreetAddress2 apijson.Field
	Locality       apijson.Field
	PostalCode     apijson.Field
	Region         apijson.Field
	CountryCode    apijson.Field
	Email          apijson.Field
	Phone          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r accountAccountGetResponsePaymentMethodsBillingAddressJSON) RawJSON() string {
	return r.raw
}

func (r *AccountAccountGetResponsePaymentMethodsBillingAddress) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r AccountAccountGetResponsePaymentMethodsBillingAddress) AsUnion() AccountAccountGetResponsePaymentMethodsBillingAddressUnion {
	return r.union
}

// The credit card's billing address
//
// Union satisfied by
// [AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceID] or
// [AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicit].
type AccountAccountGetResponsePaymentMethodsBillingAddressUnion interface {
	implementsAccountAccountGetResponsePaymentMethodsBillingAddress()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAccountGetResponsePaymentMethodsBillingAddressUnion)(nil)).Elem(),
		".tag",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceID{}),
			DiscriminatorValue: "id",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicit{}),
			DiscriminatorValue: "explicit",
		},
	)
}

type AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceID struct {
	// The address's ID
	ID string `json:"id,required"`
	// The type of address reference
	Tag  AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceIDTag  `json:".tag,required"`
	JSON accountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceIDJSON `json:"-"`
}

// accountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceIDJSON
// contains the JSON metadata for the struct
// [AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceID]
type accountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceIDJSON struct {
	ID          apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceIDJSON) RawJSON() string {
	return r.raw
}

func (r AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceID) implementsAccountAccountGetResponsePaymentMethodsBillingAddress() {
}

// The type of address reference
type AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceIDTag string

const (
	AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceIDTagID AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceIDTag = "id"
)

func (r AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceIDTag) IsKnown() bool {
	switch r {
	case AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceIDTagID:
		return true
	}
	return false
}

type AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicit struct {
	ID string `json:"id,required"`
	// The type of address reference
	Tag            AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicitTag  `json:".tag,required"`
	CountryCode    string                                                                            `json:"country_code,required"`
	FirstName      string                                                                            `json:"first_name,required"`
	LastName       string                                                                            `json:"last_name,required"`
	Locality       string                                                                            `json:"locality,required"`
	PostalCode     string                                                                            `json:"postal_code,required"`
	StreetAddress1 string                                                                            `json:"street_address1,required"`
	Company        string                                                                            `json:"company"`
	Email          string                                                                            `json:"email" format:"email"`
	Phone          string                                                                            `json:"phone" format:"phone"`
	Region         string                                                                            `json:"region"`
	StreetAddress2 string                                                                            `json:"street_address2"`
	JSON           accountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicitJSON `json:"-"`
}

// accountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicitJSON
// contains the JSON metadata for the struct
// [AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicit]
type accountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicitJSON struct {
	ID             apijson.Field
	Tag            apijson.Field
	CountryCode    apijson.Field
	FirstName      apijson.Field
	LastName       apijson.Field
	Locality       apijson.Field
	PostalCode     apijson.Field
	StreetAddress1 apijson.Field
	Company        apijson.Field
	Email          apijson.Field
	Phone          apijson.Field
	Region         apijson.Field
	StreetAddress2 apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicitJSON) RawJSON() string {
	return r.raw
}

func (r AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicit) implementsAccountAccountGetResponsePaymentMethodsBillingAddress() {
}

// The type of address reference
type AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicitTag string

const (
	AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicitTagExplicit AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicitTag = "explicit"
)

func (r AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicitTag) IsKnown() bool {
	switch r {
	case AccountAccountGetResponsePaymentMethodsBillingAddressAddressReferenceExplicitTagExplicit:
		return true
	}
	return false
}

// The type of address reference
type AccountAccountGetResponsePaymentMethodsBillingAddressTag string

const (
	AccountAccountGetResponsePaymentMethodsBillingAddressTagID       AccountAccountGetResponsePaymentMethodsBillingAddressTag = "id"
	AccountAccountGetResponsePaymentMethodsBillingAddressTagExplicit AccountAccountGetResponsePaymentMethodsBillingAddressTag = "explicit"
)

func (r AccountAccountGetResponsePaymentMethodsBillingAddressTag) IsKnown() bool {
	switch r {
	case AccountAccountGetResponsePaymentMethodsBillingAddressTagID, AccountAccountGetResponsePaymentMethodsBillingAddressTagExplicit:
		return true
	}
	return false
}

// The credit card network.
type AccountAccountGetResponsePaymentMethodsNetwork string

const (
	AccountAccountGetResponsePaymentMethodsNetworkVisa         AccountAccountGetResponsePaymentMethodsNetwork = "visa"
	AccountAccountGetResponsePaymentMethodsNetworkMastercard   AccountAccountGetResponsePaymentMethodsNetwork = "mastercard"
	AccountAccountGetResponsePaymentMethodsNetworkAmex         AccountAccountGetResponsePaymentMethodsNetwork = "amex"
	AccountAccountGetResponsePaymentMethodsNetworkDiscover     AccountAccountGetResponsePaymentMethodsNetwork = "discover"
	AccountAccountGetResponsePaymentMethodsNetworkJcb          AccountAccountGetResponsePaymentMethodsNetwork = "jcb"
	AccountAccountGetResponsePaymentMethodsNetworkUnionpay     AccountAccountGetResponsePaymentMethodsNetwork = "unionpay"
	AccountAccountGetResponsePaymentMethodsNetworkAlliancedata AccountAccountGetResponsePaymentMethodsNetwork = "alliancedata"
	AccountAccountGetResponsePaymentMethodsNetworkCitiplcc     AccountAccountGetResponsePaymentMethodsNetwork = "citiplcc"
)

func (r AccountAccountGetResponsePaymentMethodsNetwork) IsKnown() bool {
	switch r {
	case AccountAccountGetResponsePaymentMethodsNetworkVisa, AccountAccountGetResponsePaymentMethodsNetworkMastercard, AccountAccountGetResponsePaymentMethodsNetworkAmex, AccountAccountGetResponsePaymentMethodsNetworkDiscover, AccountAccountGetResponsePaymentMethodsNetworkJcb, AccountAccountGetResponsePaymentMethodsNetworkUnionpay, AccountAccountGetResponsePaymentMethodsNetworkAlliancedata, AccountAccountGetResponsePaymentMethodsNetworkCitiplcc:
		return true
	}
	return false
}

type AccountAccountGetResponseProfile struct {
	// An email address.
	Email string `json:"email"`
	// The given name of the person associated with this record.
	FirstName string `json:"first_name"`
	// The last name of the person associated with this record.
	LastName string `json:"last_name"`
	// A phone number following E164 standards, in its globalized format, i.e.
	// prepended with a plus sign.
	Phone string                               `json:"phone"`
	JSON  accountAccountGetResponseProfileJSON `json:"-"`
}

// accountAccountGetResponseProfileJSON contains the JSON metadata for the struct
// [AccountAccountGetResponseProfile]
type accountAccountGetResponseProfileJSON struct {
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	Phone       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccountGetResponseProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccountGetResponseProfileJSON) RawJSON() string {
	return r.raw
}

type AccountAccountGetParams struct {
	XPublishableKey param.Field[string] `header:"X-Publishable-Key,required"`
}
