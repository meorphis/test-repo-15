// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/tidwall/gjson"
)

// AccountPaymentMethodService contains methods and other services that help with
// interacting with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPaymentMethodService] method instead.
type AccountPaymentMethodService struct {
	Options []option.RequestOption
}

// NewAccountPaymentMethodService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPaymentMethodService(opts ...option.RequestOption) (r *AccountPaymentMethodService) {
	r = &AccountPaymentMethodService{}
	r.Options = opts
	return
}

// Delete an existing payment method. Deleting a payment method does not invalidate
// transactions or orders that are associated with it.
func (r *AccountPaymentMethodService) Delete(ctx context.Context, id string, body AccountPaymentMethodDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("account/payment-methods/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Add a payment method to a shopper's Bolt account Wallet. For security purposes,
// this request must come from your backend because authentication requires the use
// of your private key.<br /> **Note**: Before using this API, the credit card
// details must be tokenized using Bolt's JavaScript library function, which is
// documented in
// [Install the Bolt Tokenizer](https://help.bolt.com/developers/references/bolt-tokenizer).
func (r *AccountPaymentMethodService) AccountAddPaymentMethod(ctx context.Context, params AccountPaymentMethodAccountAddPaymentMethodParams, opts ...option.RequestOption) (res *AccountPaymentMethodAccountAddPaymentMethodResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "account/payment-methods"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountPaymentMethodAccountAddPaymentMethodResponse struct {
	Tag AccountPaymentMethodAccountAddPaymentMethodResponseTag `json:".tag,required"`
	// The expiration date of the credit card. TODO TO MAKE EXPIRATION REUSABLE
	Expiration string `json:"expiration,required" format:"^\\d{4}-\\d{2}$"`
	// The last 4 digits of the credit card number.
	Last4 string `json:"last4,required" format:"^\\d{4}$"`
	// The credit card network.
	Network AccountPaymentMethodAccountAddPaymentMethodResponseNetwork `json:"network,required"`
	ID      string                                                     `json:"id"`
	// The ID of credit card's billing address
	BillingAddressID string                                                  `json:"billing_address_id"`
	JSON             accountPaymentMethodAccountAddPaymentMethodResponseJSON `json:"-"`
}

// accountPaymentMethodAccountAddPaymentMethodResponseJSON contains the JSON
// metadata for the struct [AccountPaymentMethodAccountAddPaymentMethodResponse]
type accountPaymentMethodAccountAddPaymentMethodResponseJSON struct {
	Tag              apijson.Field
	Expiration       apijson.Field
	Last4            apijson.Field
	Network          apijson.Field
	ID               apijson.Field
	BillingAddressID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountPaymentMethodAccountAddPaymentMethodResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPaymentMethodAccountAddPaymentMethodResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPaymentMethodAccountAddPaymentMethodResponseTag string

const (
	AccountPaymentMethodAccountAddPaymentMethodResponseTagCreditCard AccountPaymentMethodAccountAddPaymentMethodResponseTag = "credit_card"
)

func (r AccountPaymentMethodAccountAddPaymentMethodResponseTag) IsKnown() bool {
	switch r {
	case AccountPaymentMethodAccountAddPaymentMethodResponseTagCreditCard:
		return true
	}
	return false
}

// The credit card's billing address
type AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddress struct {
	// The type of address reference
	Tag AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressTag `json:".tag,required"`
	// The address's ID
	ID             string                                                                `json:"id,required"`
	FirstName      string                                                                `json:"first_name"`
	LastName       string                                                                `json:"last_name"`
	Company        string                                                                `json:"company"`
	StreetAddress1 string                                                                `json:"street_address1"`
	StreetAddress2 string                                                                `json:"street_address2"`
	Locality       string                                                                `json:"locality"`
	PostalCode     string                                                                `json:"postal_code"`
	Region         string                                                                `json:"region"`
	CountryCode    string                                                                `json:"country_code"`
	Email          string                                                                `json:"email" format:"email"`
	Phone          string                                                                `json:"phone" format:"phone"`
	JSON           accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressJSON `json:"-"`
	union          AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressUnion
}

// accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressJSON contains
// the JSON metadata for the struct
// [AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddress]
type accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressJSON struct {
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

func (r accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressJSON) RawJSON() string {
	return r.raw
}

func (r *AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddress) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddress) AsUnion() AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressUnion {
	return r.union
}

// The credit card's billing address
//
// Union satisfied by
// [AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceID]
// or
// [AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicit].
type AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressUnion interface {
	implementsAccountPaymentMethodAccountAddPaymentMethodResponseBillingAddress()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressUnion)(nil)).Elem(),
		".tag",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceID{}),
			DiscriminatorValue: "id",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicit{}),
			DiscriminatorValue: "explicit",
		},
	)
}

type AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceID struct {
	// The address's ID
	ID string `json:"id,required"`
	// The type of address reference
	Tag  AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceIDTag  `json:".tag,required"`
	JSON accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceIDJSON `json:"-"`
}

// accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceIDJSON
// contains the JSON metadata for the struct
// [AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceID]
type accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceIDJSON struct {
	ID          apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceIDJSON) RawJSON() string {
	return r.raw
}

func (r AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceID) implementsAccountPaymentMethodAccountAddPaymentMethodResponseBillingAddress() {
}

// The type of address reference
type AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceIDTag string

const (
	AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceIDTagID AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceIDTag = "id"
)

func (r AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceIDTag) IsKnown() bool {
	switch r {
	case AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceIDTagID:
		return true
	}
	return false
}

type AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicit struct {
	ID string `json:"id,required"`
	// The type of address reference
	Tag            AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicitTag  `json:".tag,required"`
	CountryCode    string                                                                                        `json:"country_code,required"`
	FirstName      string                                                                                        `json:"first_name,required"`
	LastName       string                                                                                        `json:"last_name,required"`
	Locality       string                                                                                        `json:"locality,required"`
	PostalCode     string                                                                                        `json:"postal_code,required"`
	StreetAddress1 string                                                                                        `json:"street_address1,required"`
	Company        string                                                                                        `json:"company"`
	Email          string                                                                                        `json:"email" format:"email"`
	Phone          string                                                                                        `json:"phone" format:"phone"`
	Region         string                                                                                        `json:"region"`
	StreetAddress2 string                                                                                        `json:"street_address2"`
	JSON           accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicitJSON `json:"-"`
}

// accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicitJSON
// contains the JSON metadata for the struct
// [AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicit]
type accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicitJSON struct {
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

func (r *AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicitJSON) RawJSON() string {
	return r.raw
}

func (r AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicit) implementsAccountPaymentMethodAccountAddPaymentMethodResponseBillingAddress() {
}

// The type of address reference
type AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicitTag string

const (
	AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicitTagExplicit AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicitTag = "explicit"
)

func (r AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicitTag) IsKnown() bool {
	switch r {
	case AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressAddressReferenceExplicitTagExplicit:
		return true
	}
	return false
}

// The type of address reference
type AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressTag string

const (
	AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressTagID       AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressTag = "id"
	AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressTagExplicit AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressTag = "explicit"
)

func (r AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressTag) IsKnown() bool {
	switch r {
	case AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressTagID, AccountPaymentMethodAccountAddPaymentMethodResponseBillingAddressTagExplicit:
		return true
	}
	return false
}

// The credit card network.
type AccountPaymentMethodAccountAddPaymentMethodResponseNetwork string

const (
	AccountPaymentMethodAccountAddPaymentMethodResponseNetworkVisa         AccountPaymentMethodAccountAddPaymentMethodResponseNetwork = "visa"
	AccountPaymentMethodAccountAddPaymentMethodResponseNetworkMastercard   AccountPaymentMethodAccountAddPaymentMethodResponseNetwork = "mastercard"
	AccountPaymentMethodAccountAddPaymentMethodResponseNetworkAmex         AccountPaymentMethodAccountAddPaymentMethodResponseNetwork = "amex"
	AccountPaymentMethodAccountAddPaymentMethodResponseNetworkDiscover     AccountPaymentMethodAccountAddPaymentMethodResponseNetwork = "discover"
	AccountPaymentMethodAccountAddPaymentMethodResponseNetworkJcb          AccountPaymentMethodAccountAddPaymentMethodResponseNetwork = "jcb"
	AccountPaymentMethodAccountAddPaymentMethodResponseNetworkUnionpay     AccountPaymentMethodAccountAddPaymentMethodResponseNetwork = "unionpay"
	AccountPaymentMethodAccountAddPaymentMethodResponseNetworkAlliancedata AccountPaymentMethodAccountAddPaymentMethodResponseNetwork = "alliancedata"
	AccountPaymentMethodAccountAddPaymentMethodResponseNetworkCitiplcc     AccountPaymentMethodAccountAddPaymentMethodResponseNetwork = "citiplcc"
)

func (r AccountPaymentMethodAccountAddPaymentMethodResponseNetwork) IsKnown() bool {
	switch r {
	case AccountPaymentMethodAccountAddPaymentMethodResponseNetworkVisa, AccountPaymentMethodAccountAddPaymentMethodResponseNetworkMastercard, AccountPaymentMethodAccountAddPaymentMethodResponseNetworkAmex, AccountPaymentMethodAccountAddPaymentMethodResponseNetworkDiscover, AccountPaymentMethodAccountAddPaymentMethodResponseNetworkJcb, AccountPaymentMethodAccountAddPaymentMethodResponseNetworkUnionpay, AccountPaymentMethodAccountAddPaymentMethodResponseNetworkAlliancedata, AccountPaymentMethodAccountAddPaymentMethodResponseNetworkCitiplcc:
		return true
	}
	return false
}

type AccountPaymentMethodDeleteParams struct {
	XPublishableKey param.Field[string] `header:"X-Publishable-Key,required"`
}

type AccountPaymentMethodAccountAddPaymentMethodParams struct {
	// The Bolt token associated to the credit card.
	Token param.Field[string]                                               `json:"token,required"`
	Tag   param.Field[AccountPaymentMethodAccountAddPaymentMethodParamsTag] `json:".tag,required"`
	// The credit card's billing address
	BillingAddress param.Field[AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressUnion] `json:"billing_address,required"`
	// The Bank Identification Number for the credit card. This is typically the first
	// 4-6 digits of the credit card number.
	Bin param.Field[string] `json:"bin,required" format:"^\\d+$"`
	// The expiration date of the credit card. TODO TO MAKE EXPIRATION REUSABLE
	Expiration param.Field[string] `json:"expiration,required" format:"^\\d{4}-\\d{2}$"`
	// The last 4 digits of the credit card number.
	Last4 param.Field[string] `json:"last4,required" format:"^\\d{4}$"`
	// The credit card network.
	Network         param.Field[AccountPaymentMethodAccountAddPaymentMethodParamsNetwork] `json:"network,required"`
	XPublishableKey param.Field[string]                                                   `header:"X-Publishable-Key,required"`
}

func (r AccountPaymentMethodAccountAddPaymentMethodParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountPaymentMethodAccountAddPaymentMethodParamsTag string

const (
	AccountPaymentMethodAccountAddPaymentMethodParamsTagCreditCard AccountPaymentMethodAccountAddPaymentMethodParamsTag = "credit_card"
)

func (r AccountPaymentMethodAccountAddPaymentMethodParamsTag) IsKnown() bool {
	switch r {
	case AccountPaymentMethodAccountAddPaymentMethodParamsTagCreditCard:
		return true
	}
	return false
}

// The credit card's billing address
type AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddress struct {
	// The type of address reference
	Tag param.Field[AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressTag] `json:".tag,required"`
	// The address's ID
	ID             param.Field[string] `json:"id,required"`
	FirstName      param.Field[string] `json:"first_name"`
	LastName       param.Field[string] `json:"last_name"`
	Company        param.Field[string] `json:"company"`
	StreetAddress1 param.Field[string] `json:"street_address1"`
	StreetAddress2 param.Field[string] `json:"street_address2"`
	Locality       param.Field[string] `json:"locality"`
	PostalCode     param.Field[string] `json:"postal_code"`
	Region         param.Field[string] `json:"region"`
	CountryCode    param.Field[string] `json:"country_code"`
	Email          param.Field[string] `json:"email" format:"email"`
	Phone          param.Field[string] `json:"phone" format:"phone"`
}

func (r AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddress) implementsAccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressUnion() {
}

// The credit card's billing address
//
// Satisfied by
// [AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceID],
// [AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceExplicit],
// [AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddress].
type AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressUnion interface {
	implementsAccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressUnion()
}

type AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceID struct {
	// The address's ID
	ID param.Field[string] `json:"id,required"`
	// The type of address reference
	Tag param.Field[AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceIDTag] `json:".tag,required"`
}

func (r AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceID) implementsAccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressUnion() {
}

// The type of address reference
type AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceIDTag string

const (
	AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceIDTagID AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceIDTag = "id"
)

func (r AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceIDTag) IsKnown() bool {
	switch r {
	case AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceIDTagID:
		return true
	}
	return false
}

type AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceExplicit struct {
	// The type of address reference
	Tag            param.Field[AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceExplicitTag] `json:".tag,required"`
	CountryCode    param.Field[string]                                                                                     `json:"country_code,required"`
	FirstName      param.Field[string]                                                                                     `json:"first_name,required"`
	LastName       param.Field[string]                                                                                     `json:"last_name,required"`
	Locality       param.Field[string]                                                                                     `json:"locality,required"`
	PostalCode     param.Field[string]                                                                                     `json:"postal_code,required"`
	StreetAddress1 param.Field[string]                                                                                     `json:"street_address1,required"`
	Company        param.Field[string]                                                                                     `json:"company"`
	Email          param.Field[string]                                                                                     `json:"email" format:"email"`
	Phone          param.Field[string]                                                                                     `json:"phone" format:"phone"`
	Region         param.Field[string]                                                                                     `json:"region"`
	StreetAddress2 param.Field[string]                                                                                     `json:"street_address2"`
}

func (r AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceExplicit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceExplicit) implementsAccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressUnion() {
}

// The type of address reference
type AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceExplicitTag string

const (
	AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceExplicitTagExplicit AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceExplicitTag = "explicit"
)

func (r AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceExplicitTag) IsKnown() bool {
	switch r {
	case AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceExplicitTagExplicit:
		return true
	}
	return false
}

// The type of address reference
type AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressTag string

const (
	AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressTagID       AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressTag = "id"
	AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressTagExplicit AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressTag = "explicit"
)

func (r AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressTag) IsKnown() bool {
	switch r {
	case AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressTagID, AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressTagExplicit:
		return true
	}
	return false
}

// The credit card network.
type AccountPaymentMethodAccountAddPaymentMethodParamsNetwork string

const (
	AccountPaymentMethodAccountAddPaymentMethodParamsNetworkVisa         AccountPaymentMethodAccountAddPaymentMethodParamsNetwork = "visa"
	AccountPaymentMethodAccountAddPaymentMethodParamsNetworkMastercard   AccountPaymentMethodAccountAddPaymentMethodParamsNetwork = "mastercard"
	AccountPaymentMethodAccountAddPaymentMethodParamsNetworkAmex         AccountPaymentMethodAccountAddPaymentMethodParamsNetwork = "amex"
	AccountPaymentMethodAccountAddPaymentMethodParamsNetworkDiscover     AccountPaymentMethodAccountAddPaymentMethodParamsNetwork = "discover"
	AccountPaymentMethodAccountAddPaymentMethodParamsNetworkJcb          AccountPaymentMethodAccountAddPaymentMethodParamsNetwork = "jcb"
	AccountPaymentMethodAccountAddPaymentMethodParamsNetworkUnionpay     AccountPaymentMethodAccountAddPaymentMethodParamsNetwork = "unionpay"
	AccountPaymentMethodAccountAddPaymentMethodParamsNetworkAlliancedata AccountPaymentMethodAccountAddPaymentMethodParamsNetwork = "alliancedata"
	AccountPaymentMethodAccountAddPaymentMethodParamsNetworkCitiplcc     AccountPaymentMethodAccountAddPaymentMethodParamsNetwork = "citiplcc"
)

func (r AccountPaymentMethodAccountAddPaymentMethodParamsNetwork) IsKnown() bool {
	switch r {
	case AccountPaymentMethodAccountAddPaymentMethodParamsNetworkVisa, AccountPaymentMethodAccountAddPaymentMethodParamsNetworkMastercard, AccountPaymentMethodAccountAddPaymentMethodParamsNetworkAmex, AccountPaymentMethodAccountAddPaymentMethodParamsNetworkDiscover, AccountPaymentMethodAccountAddPaymentMethodParamsNetworkJcb, AccountPaymentMethodAccountAddPaymentMethodParamsNetworkUnionpay, AccountPaymentMethodAccountAddPaymentMethodParamsNetworkAlliancedata, AccountPaymentMethodAccountAddPaymentMethodParamsNetworkCitiplcc:
		return true
	}
	return false
}
