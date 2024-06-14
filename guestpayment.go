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

// GuestPaymentService contains methods and other services that help with
// interacting with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGuestPaymentService] method instead.
type GuestPaymentService struct {
	Options []option.RequestOption
}

// NewGuestPaymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGuestPaymentService(opts ...option.RequestOption) (r *GuestPaymentService) {
	r = &GuestPaymentService{}
	r.Options = opts
	return
}

// Initialize a Bolt payment token that will be used to reference this payment to
// Bolt when it is updated or finalized for guest shoppers.
func (r *GuestPaymentService) GuestPaymentsInitialize(ctx context.Context, params GuestPaymentGuestPaymentsInitializeParams, opts ...option.RequestOption) (res *GuestPaymentGuestPaymentsInitializeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "guest/payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type GuestPaymentGuestPaymentsInitializeResponse struct {
	ID     string                                            `json:"id"`
	Action GuestPaymentGuestPaymentsInitializeResponseAction `json:"action"`
	Status GuestPaymentGuestPaymentsInitializeResponseStatus `json:"status"`
	JSON   guestPaymentGuestPaymentsInitializeResponseJSON   `json:"-"`
}

// guestPaymentGuestPaymentsInitializeResponseJSON contains the JSON metadata for
// the struct [GuestPaymentGuestPaymentsInitializeResponse]
type guestPaymentGuestPaymentsInitializeResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GuestPaymentGuestPaymentsInitializeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r guestPaymentGuestPaymentsInitializeResponseJSON) RawJSON() string {
	return r.raw
}

type GuestPaymentGuestPaymentsInitializeResponseAction struct {
	Method GuestPaymentGuestPaymentsInitializeResponseActionMethod `json:"method"`
	Type   GuestPaymentGuestPaymentsInitializeResponseActionType   `json:"type"`
	URL    string                                                  `json:"url" format:"uri"`
	JSON   guestPaymentGuestPaymentsInitializeResponseActionJSON   `json:"-"`
}

// guestPaymentGuestPaymentsInitializeResponseActionJSON contains the JSON metadata
// for the struct [GuestPaymentGuestPaymentsInitializeResponseAction]
type guestPaymentGuestPaymentsInitializeResponseActionJSON struct {
	Method      apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GuestPaymentGuestPaymentsInitializeResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r guestPaymentGuestPaymentsInitializeResponseActionJSON) RawJSON() string {
	return r.raw
}

type GuestPaymentGuestPaymentsInitializeResponseActionMethod string

const (
	GuestPaymentGuestPaymentsInitializeResponseActionMethodGet  GuestPaymentGuestPaymentsInitializeResponseActionMethod = "GET"
	GuestPaymentGuestPaymentsInitializeResponseActionMethodPost GuestPaymentGuestPaymentsInitializeResponseActionMethod = "POST"
)

func (r GuestPaymentGuestPaymentsInitializeResponseActionMethod) IsKnown() bool {
	switch r {
	case GuestPaymentGuestPaymentsInitializeResponseActionMethodGet, GuestPaymentGuestPaymentsInitializeResponseActionMethodPost:
		return true
	}
	return false
}

type GuestPaymentGuestPaymentsInitializeResponseActionType string

const (
	GuestPaymentGuestPaymentsInitializeResponseActionTypeRedirect GuestPaymentGuestPaymentsInitializeResponseActionType = "redirect"
	GuestPaymentGuestPaymentsInitializeResponseActionTypeFinalize GuestPaymentGuestPaymentsInitializeResponseActionType = "finalize"
)

func (r GuestPaymentGuestPaymentsInitializeResponseActionType) IsKnown() bool {
	switch r {
	case GuestPaymentGuestPaymentsInitializeResponseActionTypeRedirect, GuestPaymentGuestPaymentsInitializeResponseActionTypeFinalize:
		return true
	}
	return false
}

type GuestPaymentGuestPaymentsInitializeResponseStatus string

const (
	GuestPaymentGuestPaymentsInitializeResponseStatusAwaitingUserConfirmation GuestPaymentGuestPaymentsInitializeResponseStatus = "awaiting_user_confirmation"
	GuestPaymentGuestPaymentsInitializeResponseStatusPaymentReady             GuestPaymentGuestPaymentsInitializeResponseStatus = "payment_ready"
	GuestPaymentGuestPaymentsInitializeResponseStatusUpdatePaymentMethod      GuestPaymentGuestPaymentsInitializeResponseStatus = "update_payment_method"
	GuestPaymentGuestPaymentsInitializeResponseStatusSuccess                  GuestPaymentGuestPaymentsInitializeResponseStatus = "success"
)

func (r GuestPaymentGuestPaymentsInitializeResponseStatus) IsKnown() bool {
	switch r {
	case GuestPaymentGuestPaymentsInitializeResponseStatusAwaitingUserConfirmation, GuestPaymentGuestPaymentsInitializeResponseStatusPaymentReady, GuestPaymentGuestPaymentsInitializeResponseStatusUpdatePaymentMethod, GuestPaymentGuestPaymentsInitializeResponseStatusSuccess:
		return true
	}
	return false
}

type GuestPaymentGuestPaymentsInitializeParams struct {
	Cart            param.Field[GuestPaymentGuestPaymentsInitializeParamsCart]          `json:"cart,required"`
	PaymentMethod   param.Field[GuestPaymentGuestPaymentsInitializeParamsPaymentMethod] `json:"payment_method,required"`
	XPublishableKey param.Field[string]                                                 `header:"X-Publishable-Key,required"`
}

func (r GuestPaymentGuestPaymentsInitializeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentGuestPaymentsInitializeParamsCart struct {
	Amounts param.Field[GuestPaymentGuestPaymentsInitializeParamsCartAmounts] `json:"amounts,required"`
	// This value is used by Bolt as an external reference to a given order. This
	// reference must be unique per successful transaction.
	OrderReference param.Field[string]                                                  `json:"order_reference,required"`
	Discounts      param.Field[[]GuestPaymentGuestPaymentsInitializeParamsCartDiscount] `json:"discounts"`
	// This field corresponds to the merchant's order reference associated with this
	// Bolt transaction.
	DisplayID param.Field[string]                                              `json:"display_id"`
	Items     param.Field[[]GuestPaymentGuestPaymentsInitializeParamsCartItem] `json:"items"`
	// Used optionally to pass additional information like order numbers or other IDs
	// as needed.
	OrderDescription param.Field[string]                                                  `json:"order_description"`
	Shipments        param.Field[[]GuestPaymentGuestPaymentsInitializeParamsCartShipment] `json:"shipments"`
}

func (r GuestPaymentGuestPaymentsInitializeParamsCart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentGuestPaymentsInitializeParamsCartAmounts struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r GuestPaymentGuestPaymentsInitializeParamsCartAmounts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentGuestPaymentsInitializeParamsCartDiscount struct {
	Amounts param.Field[GuestPaymentGuestPaymentsInitializeParamsCartDiscountsAmounts] `json:"amounts,required"`
	// Discount code.
	Code param.Field[string] `json:"code"`
	// Used to provide a link to additional details, such as a landing page, associated
	// with the discount offering.
	DetailsURL param.Field[string] `json:"details_url" format:"uri"`
}

func (r GuestPaymentGuestPaymentsInitializeParamsCartDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentGuestPaymentsInitializeParamsCartDiscountsAmounts struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r GuestPaymentGuestPaymentsInitializeParamsCartDiscountsAmounts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentGuestPaymentsInitializeParamsCartItem struct {
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

func (r GuestPaymentGuestPaymentsInitializeParamsCartItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentGuestPaymentsInitializeParamsCartShipment struct {
	// The Address object is used for shipping, and physical store address use cases.
	Address param.Field[GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressUnion] `json:"address"`
	// The name of the carrier selected.
	Carrier param.Field[string]                                                     `json:"carrier"`
	Cost    param.Field[GuestPaymentGuestPaymentsInitializeParamsCartShipmentsCost] `json:"cost"`
}

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Address object is used for shipping, and physical store address use cases.
type GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddress struct {
	// The type of address reference
	Tag param.Field[GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressTag] `json:".tag,required"`
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

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddress) implementsGuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressUnion() {
}

// The Address object is used for shipping, and physical store address use cases.
//
// Satisfied by
// [GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceID],
// [GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceExplicit],
// [GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddress].
type GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressUnion interface {
	implementsGuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressUnion()
}

type GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceID struct {
	// The address's ID
	ID param.Field[string] `json:"id,required"`
	// The type of address reference
	Tag param.Field[GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceIDTag] `json:".tag,required"`
}

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceID) implementsGuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressUnion() {
}

// The type of address reference
type GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceIDTag string

const (
	GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceIDTagID GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceIDTag = "id"
)

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceIDTag) IsKnown() bool {
	switch r {
	case GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceIDTagID:
		return true
	}
	return false
}

type GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceExplicit struct {
	// The type of address reference
	Tag            param.Field[GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceExplicitTag] `json:".tag,required"`
	CountryCode    param.Field[string]                                                                                   `json:"country_code,required"`
	FirstName      param.Field[string]                                                                                   `json:"first_name,required"`
	LastName       param.Field[string]                                                                                   `json:"last_name,required"`
	Locality       param.Field[string]                                                                                   `json:"locality,required"`
	PostalCode     param.Field[string]                                                                                   `json:"postal_code,required"`
	StreetAddress1 param.Field[string]                                                                                   `json:"street_address1,required"`
	Company        param.Field[string]                                                                                   `json:"company"`
	Email          param.Field[string]                                                                                   `json:"email" format:"email"`
	Phone          param.Field[string]                                                                                   `json:"phone" format:"phone"`
	Region         param.Field[string]                                                                                   `json:"region"`
	StreetAddress2 param.Field[string]                                                                                   `json:"street_address2"`
}

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceExplicit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceExplicit) implementsGuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressUnion() {
}

// The type of address reference
type GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceExplicitTag string

const (
	GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceExplicitTagExplicit GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceExplicitTag = "explicit"
)

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceExplicitTag) IsKnown() bool {
	switch r {
	case GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceExplicitTagExplicit:
		return true
	}
	return false
}

// The type of address reference
type GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressTag string

const (
	GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressTagID       GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressTag = "id"
	GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressTagExplicit GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressTag = "explicit"
)

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressTag) IsKnown() bool {
	switch r {
	case GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressTagID, GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressTagExplicit:
		return true
	}
	return false
}

type GuestPaymentGuestPaymentsInitializeParamsCartShipmentsCost struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r GuestPaymentGuestPaymentsInitializeParamsCartShipmentsCost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentGuestPaymentsInitializeParamsPaymentMethod struct {
	Tag param.Field[GuestPaymentGuestPaymentsInitializeParamsPaymentMethodTag] `json:".tag,required"`
	// Redirect URL for canceled PayPal transaction.
	Cancel param.Field[string] `json:"cancel,required" format:"uri"`
	// Redirect URL for successful PayPal transaction.
	Success param.Field[string] `json:"success,required" format:"uri"`
}

func (r GuestPaymentGuestPaymentsInitializeParamsPaymentMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentGuestPaymentsInitializeParamsPaymentMethodTag string

const (
	GuestPaymentGuestPaymentsInitializeParamsPaymentMethodTagPaypal GuestPaymentGuestPaymentsInitializeParamsPaymentMethodTag = "paypal"
)

func (r GuestPaymentGuestPaymentsInitializeParamsPaymentMethodTag) IsKnown() bool {
	switch r {
	case GuestPaymentGuestPaymentsInitializeParamsPaymentMethodTagPaypal:
		return true
	}
	return false
}
