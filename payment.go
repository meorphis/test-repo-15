// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// PaymentService contains methods and other services that help with interacting
// with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	Options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r *PaymentService) {
	r = &PaymentService{}
	r.Options = opts
	return
}

// Initialize a Bolt payment token that will be used to reference this payment to
// Bolt when it is updated or finalized for logged in shoppers.
func (r *PaymentService) New(ctx context.Context, params PaymentNewParams, opts ...option.RequestOption) (res *PaymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PaymentNewResponse struct {
	ID     string                   `json:"id"`
	Action PaymentNewResponseAction `json:"action"`
	Status PaymentNewResponseStatus `json:"status"`
	JSON   paymentNewResponseJSON   `json:"-"`
}

// paymentNewResponseJSON contains the JSON metadata for the struct
// [PaymentNewResponse]
type paymentNewResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentNewResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentNewResponseAction struct {
	Method PaymentNewResponseActionMethod `json:"method"`
	Type   PaymentNewResponseActionType   `json:"type"`
	URL    string                         `json:"url" format:"uri"`
	JSON   paymentNewResponseActionJSON   `json:"-"`
}

// paymentNewResponseActionJSON contains the JSON metadata for the struct
// [PaymentNewResponseAction]
type paymentNewResponseActionJSON struct {
	Method      apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentNewResponseActionJSON) RawJSON() string {
	return r.raw
}

type PaymentNewResponseActionMethod string

const (
	PaymentNewResponseActionMethodGet  PaymentNewResponseActionMethod = "GET"
	PaymentNewResponseActionMethodPost PaymentNewResponseActionMethod = "POST"
)

func (r PaymentNewResponseActionMethod) IsKnown() bool {
	switch r {
	case PaymentNewResponseActionMethodGet, PaymentNewResponseActionMethodPost:
		return true
	}
	return false
}

type PaymentNewResponseActionType string

const (
	PaymentNewResponseActionTypeRedirect PaymentNewResponseActionType = "redirect"
	PaymentNewResponseActionTypeFinalize PaymentNewResponseActionType = "finalize"
)

func (r PaymentNewResponseActionType) IsKnown() bool {
	switch r {
	case PaymentNewResponseActionTypeRedirect, PaymentNewResponseActionTypeFinalize:
		return true
	}
	return false
}

type PaymentNewResponseStatus string

const (
	PaymentNewResponseStatusAwaitingUserConfirmation PaymentNewResponseStatus = "awaiting_user_confirmation"
	PaymentNewResponseStatusPaymentReady             PaymentNewResponseStatus = "payment_ready"
	PaymentNewResponseStatusUpdatePaymentMethod      PaymentNewResponseStatus = "update_payment_method"
	PaymentNewResponseStatusSuccess                  PaymentNewResponseStatus = "success"
)

func (r PaymentNewResponseStatus) IsKnown() bool {
	switch r {
	case PaymentNewResponseStatusAwaitingUserConfirmation, PaymentNewResponseStatusPaymentReady, PaymentNewResponseStatusUpdatePaymentMethod, PaymentNewResponseStatusSuccess:
		return true
	}
	return false
}

type PaymentNewParams struct {
	Cart            param.Field[PaymentNewParamsCart]          `json:"cart,required"`
	PaymentMethod   param.Field[PaymentNewParamsPaymentMethod] `json:"payment_method,required"`
	XPublishableKey param.Field[string]                        `header:"X-Publishable-Key,required"`
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCart struct {
	Amounts param.Field[PaymentNewParamsCartAmounts] `json:"amounts,required"`
	// This value is used by Bolt as an external reference to a given order. This
	// reference must be unique per successful transaction.
	OrderReference param.Field[string]                         `json:"order_reference,required"`
	Discounts      param.Field[[]PaymentNewParamsCartDiscount] `json:"discounts"`
	// This field corresponds to the merchant's order reference associated with this
	// Bolt transaction.
	DisplayID param.Field[string]                     `json:"display_id"`
	Items     param.Field[[]PaymentNewParamsCartItem] `json:"items"`
	// Used optionally to pass additional information like order numbers or other IDs
	// as needed.
	OrderDescription param.Field[string]                         `json:"order_description"`
	Shipments        param.Field[[]PaymentNewParamsCartShipment] `json:"shipments"`
}

func (r PaymentNewParamsCart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCartAmounts struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r PaymentNewParamsCartAmounts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCartDiscount struct {
	Amounts param.Field[PaymentNewParamsCartDiscountsAmounts] `json:"amounts,required"`
	// Discount code.
	Code param.Field[string] `json:"code"`
	// Used to provide a link to additional details, such as a landing page, associated
	// with the discount offering.
	DetailsURL param.Field[string] `json:"details_url" format:"uri"`
}

func (r PaymentNewParamsCartDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCartDiscountsAmounts struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r PaymentNewParamsCartDiscountsAmounts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCartItem struct {
	// The name of a given item.
	Name param.Field[string] `json:"name,required"`
	// The number of units that comprise this cart item.
	Quantity param.Field[int64] `json:"quantity,required"`
	// This value is used by Bolt as an external reference to a given item.
	Reference param.Field[string] `json:"reference,required"`
	// The total amount, in cents, of the item including its taxes if applicable.
	TotalAmount param.Field[int64] `json:"total_amount,required"`
	// The price of one unit of the item; for example, the price of one pack of socks.
	UnitPrice param.Field[int64] `json:"unit_price,required"`
	// A human-readable description of this cart item.
	Description param.Field[string] `json:"description"`
	// Used to provide a link to the image associated with the item.
	ImageURL param.Field[string] `json:"image_url" format:"uri"`
}

func (r PaymentNewParamsCartItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCartShipment struct {
	// The Address object is used for shipping, and physical store address use cases.
	Address param.Field[PaymentNewParamsCartShipmentsAddressUnion] `json:"address"`
	// The name of the carrier selected.
	Carrier param.Field[string]                            `json:"carrier"`
	Cost    param.Field[PaymentNewParamsCartShipmentsCost] `json:"cost"`
}

func (r PaymentNewParamsCartShipment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Address object is used for shipping, and physical store address use cases.
type PaymentNewParamsCartShipmentsAddress struct {
	// The type of address reference
	Tag param.Field[PaymentNewParamsCartShipmentsAddressTag] `json:".tag,required"`
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

func (r PaymentNewParamsCartShipmentsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PaymentNewParamsCartShipmentsAddress) implementsPaymentNewParamsCartShipmentsAddressUnion() {}

// The Address object is used for shipping, and physical store address use cases.
//
// Satisfied by [PaymentNewParamsCartShipmentsAddressAddressReferenceID],
// [PaymentNewParamsCartShipmentsAddressAddressReferenceExplicit],
// [PaymentNewParamsCartShipmentsAddress].
type PaymentNewParamsCartShipmentsAddressUnion interface {
	implementsPaymentNewParamsCartShipmentsAddressUnion()
}

type PaymentNewParamsCartShipmentsAddressAddressReferenceID struct {
	// The address's ID
	ID param.Field[string] `json:"id,required"`
	// The type of address reference
	Tag param.Field[PaymentNewParamsCartShipmentsAddressAddressReferenceIDTag] `json:".tag,required"`
}

func (r PaymentNewParamsCartShipmentsAddressAddressReferenceID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PaymentNewParamsCartShipmentsAddressAddressReferenceID) implementsPaymentNewParamsCartShipmentsAddressUnion() {
}

// The type of address reference
type PaymentNewParamsCartShipmentsAddressAddressReferenceIDTag string

const (
	PaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID PaymentNewParamsCartShipmentsAddressAddressReferenceIDTag = "id"
)

func (r PaymentNewParamsCartShipmentsAddressAddressReferenceIDTag) IsKnown() bool {
	switch r {
	case PaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID:
		return true
	}
	return false
}

type PaymentNewParamsCartShipmentsAddressAddressReferenceExplicit struct {
	// The type of address reference
	Tag            param.Field[PaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTag] `json:".tag,required"`
	CountryCode    param.Field[string]                                                          `json:"country_code,required"`
	FirstName      param.Field[string]                                                          `json:"first_name,required"`
	LastName       param.Field[string]                                                          `json:"last_name,required"`
	Locality       param.Field[string]                                                          `json:"locality,required"`
	PostalCode     param.Field[string]                                                          `json:"postal_code,required"`
	StreetAddress1 param.Field[string]                                                          `json:"street_address1,required"`
	Company        param.Field[string]                                                          `json:"company"`
	Email          param.Field[string]                                                          `json:"email" format:"email"`
	Phone          param.Field[string]                                                          `json:"phone" format:"phone"`
	Region         param.Field[string]                                                          `json:"region"`
	StreetAddress2 param.Field[string]                                                          `json:"street_address2"`
}

func (r PaymentNewParamsCartShipmentsAddressAddressReferenceExplicit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PaymentNewParamsCartShipmentsAddressAddressReferenceExplicit) implementsPaymentNewParamsCartShipmentsAddressUnion() {
}

// The type of address reference
type PaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTag string

const (
	PaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTagExplicit PaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTag = "explicit"
)

func (r PaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTag) IsKnown() bool {
	switch r {
	case PaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTagExplicit:
		return true
	}
	return false
}

// The type of address reference
type PaymentNewParamsCartShipmentsAddressTag string

const (
	PaymentNewParamsCartShipmentsAddressTagID       PaymentNewParamsCartShipmentsAddressTag = "id"
	PaymentNewParamsCartShipmentsAddressTagExplicit PaymentNewParamsCartShipmentsAddressTag = "explicit"
)

func (r PaymentNewParamsCartShipmentsAddressTag) IsKnown() bool {
	switch r {
	case PaymentNewParamsCartShipmentsAddressTagID, PaymentNewParamsCartShipmentsAddressTagExplicit:
		return true
	}
	return false
}

type PaymentNewParamsCartShipmentsCost struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r PaymentNewParamsCartShipmentsCost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsPaymentMethod struct {
	// Payment ID of the saved Bolt Payment method.
	ID  param.Field[string]                           `json:"id,required"`
	Tag param.Field[PaymentNewParamsPaymentMethodTag] `json:".tag,required"`
}

func (r PaymentNewParamsPaymentMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsPaymentMethodTag string

const (
	PaymentNewParamsPaymentMethodTagSavedPaymentMethod PaymentNewParamsPaymentMethodTag = "saved_payment_method"
)

func (r PaymentNewParamsPaymentMethodTag) IsKnown() bool {
	switch r {
	case PaymentNewParamsPaymentMethodTagSavedPaymentMethod:
		return true
	}
	return false
}
