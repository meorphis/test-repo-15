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

// MerchantCallbackService contains methods and other services that help with
// interacting with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMerchantCallbackService] method instead.
type MerchantCallbackService struct {
	Options []option.RequestOption
}

// NewMerchantCallbackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMerchantCallbackService(opts ...option.RequestOption) (r *MerchantCallbackService) {
	r = &MerchantCallbackService{}
	r.Options = opts
	return
}

// Return callback URLs configured on the merchant such as OAuth URLs.
func (r *MerchantCallbackService) List(ctx context.Context, query MerchantCallbackListParams, opts ...option.RequestOption) (res *MerchantCallbackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "merchant/callbacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update and configure callback URLs on the merchant such as OAuth URLs.
func (r *MerchantCallbackService) MerchantCallbacksUpdate(ctx context.Context, params MerchantCallbackMerchantCallbacksUpdateParams, opts ...option.RequestOption) (res *MerchantCallbackMerchantCallbacksUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "merchant/callbacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type MerchantCallbackListResponse struct {
	AccountPage                   string                           `json:"account_page" format:"uri"`
	BaseDomain                    string                           `json:"base_domain" format:"uri"`
	ConfirmationRedirect          string                           `json:"confirmation_redirect" format:"uri"`
	CreateOrder                   string                           `json:"create_order" format:"uri"`
	Debug                         string                           `json:"debug" format:"uri"`
	GetAccount                    string                           `json:"get_account" format:"uri"`
	MobileAppDomain               string                           `json:"mobile_app_domain" format:"uri"`
	OAuthLogout                   string                           `json:"oauth_logout" format:"uri"`
	OAuthRedirect                 string                           `json:"oauth_redirect" format:"uri"`
	PrivacyPolicy                 string                           `json:"privacy_policy" format:"uri"`
	ProductInfo                   string                           `json:"product_info" format:"uri"`
	RemoteAPI                     string                           `json:"remote_api" format:"uri"`
	Shipping                      string                           `json:"shipping" format:"uri"`
	SupportPage                   string                           `json:"support_page" format:"uri"`
	Tax                           string                           `json:"tax" format:"uri"`
	TermsOfService                string                           `json:"terms_of_service" format:"uri"`
	UniversalMerchantAPI          string                           `json:"universal_merchant_api" format:"uri"`
	UpdateCart                    string                           `json:"update_cart" format:"uri"`
	ValidateAdditionalAccountData string                           `json:"validate_additional_account_data" format:"uri"`
	JSON                          merchantCallbackListResponseJSON `json:"-"`
}

// merchantCallbackListResponseJSON contains the JSON metadata for the struct
// [MerchantCallbackListResponse]
type merchantCallbackListResponseJSON struct {
	AccountPage                   apijson.Field
	BaseDomain                    apijson.Field
	ConfirmationRedirect          apijson.Field
	CreateOrder                   apijson.Field
	Debug                         apijson.Field
	GetAccount                    apijson.Field
	MobileAppDomain               apijson.Field
	OAuthLogout                   apijson.Field
	OAuthRedirect                 apijson.Field
	PrivacyPolicy                 apijson.Field
	ProductInfo                   apijson.Field
	RemoteAPI                     apijson.Field
	Shipping                      apijson.Field
	SupportPage                   apijson.Field
	Tax                           apijson.Field
	TermsOfService                apijson.Field
	UniversalMerchantAPI          apijson.Field
	UpdateCart                    apijson.Field
	ValidateAdditionalAccountData apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *MerchantCallbackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r merchantCallbackListResponseJSON) RawJSON() string {
	return r.raw
}

type MerchantCallbackMerchantCallbacksUpdateResponse struct {
	AccountPage                   string                                              `json:"account_page" format:"uri"`
	BaseDomain                    string                                              `json:"base_domain" format:"uri"`
	ConfirmationRedirect          string                                              `json:"confirmation_redirect" format:"uri"`
	CreateOrder                   string                                              `json:"create_order" format:"uri"`
	Debug                         string                                              `json:"debug" format:"uri"`
	GetAccount                    string                                              `json:"get_account" format:"uri"`
	MobileAppDomain               string                                              `json:"mobile_app_domain" format:"uri"`
	OAuthLogout                   string                                              `json:"oauth_logout" format:"uri"`
	OAuthRedirect                 string                                              `json:"oauth_redirect" format:"uri"`
	PrivacyPolicy                 string                                              `json:"privacy_policy" format:"uri"`
	ProductInfo                   string                                              `json:"product_info" format:"uri"`
	RemoteAPI                     string                                              `json:"remote_api" format:"uri"`
	Shipping                      string                                              `json:"shipping" format:"uri"`
	SupportPage                   string                                              `json:"support_page" format:"uri"`
	Tax                           string                                              `json:"tax" format:"uri"`
	TermsOfService                string                                              `json:"terms_of_service" format:"uri"`
	UniversalMerchantAPI          string                                              `json:"universal_merchant_api" format:"uri"`
	UpdateCart                    string                                              `json:"update_cart" format:"uri"`
	ValidateAdditionalAccountData string                                              `json:"validate_additional_account_data" format:"uri"`
	JSON                          merchantCallbackMerchantCallbacksUpdateResponseJSON `json:"-"`
}

// merchantCallbackMerchantCallbacksUpdateResponseJSON contains the JSON metadata
// for the struct [MerchantCallbackMerchantCallbacksUpdateResponse]
type merchantCallbackMerchantCallbacksUpdateResponseJSON struct {
	AccountPage                   apijson.Field
	BaseDomain                    apijson.Field
	ConfirmationRedirect          apijson.Field
	CreateOrder                   apijson.Field
	Debug                         apijson.Field
	GetAccount                    apijson.Field
	MobileAppDomain               apijson.Field
	OAuthLogout                   apijson.Field
	OAuthRedirect                 apijson.Field
	PrivacyPolicy                 apijson.Field
	ProductInfo                   apijson.Field
	RemoteAPI                     apijson.Field
	Shipping                      apijson.Field
	SupportPage                   apijson.Field
	Tax                           apijson.Field
	TermsOfService                apijson.Field
	UniversalMerchantAPI          apijson.Field
	UpdateCart                    apijson.Field
	ValidateAdditionalAccountData apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *MerchantCallbackMerchantCallbacksUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r merchantCallbackMerchantCallbacksUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type MerchantCallbackListParams struct {
	XPublishableKey param.Field[string] `header:"X-Publishable-Key,required"`
}

type MerchantCallbackMerchantCallbacksUpdateParams struct {
	XPublishableKey               param.Field[string] `header:"X-Publishable-Key,required"`
	AccountPage                   param.Field[string] `json:"account_page" format:"uri"`
	BaseDomain                    param.Field[string] `json:"base_domain" format:"uri"`
	ConfirmationRedirect          param.Field[string] `json:"confirmation_redirect" format:"uri"`
	CreateOrder                   param.Field[string] `json:"create_order" format:"uri"`
	Debug                         param.Field[string] `json:"debug" format:"uri"`
	GetAccount                    param.Field[string] `json:"get_account" format:"uri"`
	MobileAppDomain               param.Field[string] `json:"mobile_app_domain" format:"uri"`
	OAuthLogout                   param.Field[string] `json:"oauth_logout" format:"uri"`
	OAuthRedirect                 param.Field[string] `json:"oauth_redirect" format:"uri"`
	PrivacyPolicy                 param.Field[string] `json:"privacy_policy" format:"uri"`
	ProductInfo                   param.Field[string] `json:"product_info" format:"uri"`
	RemoteAPI                     param.Field[string] `json:"remote_api" format:"uri"`
	Shipping                      param.Field[string] `json:"shipping" format:"uri"`
	SupportPage                   param.Field[string] `json:"support_page" format:"uri"`
	Tax                           param.Field[string] `json:"tax" format:"uri"`
	TermsOfService                param.Field[string] `json:"terms_of_service" format:"uri"`
	UniversalMerchantAPI          param.Field[string] `json:"universal_merchant_api" format:"uri"`
	UpdateCart                    param.Field[string] `json:"update_cart" format:"uri"`
	ValidateAdditionalAccountData param.Field[string] `json:"validate_additional_account_data" format:"uri"`
}

func (r MerchantCallbackMerchantCallbacksUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
