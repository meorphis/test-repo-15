// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// MerchantIdentifierService contains methods and other services that help with
// interacting with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMerchantIdentifierService] method instead.
type MerchantIdentifierService struct {
	Options []option.RequestOption
}

// NewMerchantIdentifierService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMerchantIdentifierService(opts ...option.RequestOption) (r *MerchantIdentifierService) {
	r = &MerchantIdentifierService{}
	r.Options = opts
	return
}

// Return several identifiers for the merchant, such as public IDs, publishable
// keys, signing secrets, etc...
func (r *MerchantIdentifierService) List(ctx context.Context, opts ...option.RequestOption) (res *MerchantIdentifierListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "merchant/identifiers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MerchantIdentifierListResponse struct {
	MerchantDivisions []MerchantIdentifierListResponseMerchantDivision `json:"merchant_divisions,required"`
	MerchantID        string                                           `json:"merchant_id,required"`
	SigningSecret     string                                           `json:"signing_secret,required"`
	JSON              merchantIdentifierListResponseJSON               `json:"-"`
}

// merchantIdentifierListResponseJSON contains the JSON metadata for the struct
// [MerchantIdentifierListResponse]
type merchantIdentifierListResponseJSON struct {
	MerchantDivisions apijson.Field
	MerchantID        apijson.Field
	SigningSecret     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MerchantIdentifierListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r merchantIdentifierListResponseJSON) RawJSON() string {
	return r.raw
}

type MerchantIdentifierListResponseMerchantDivision struct {
	PublishableKey string                                             `json:"publishable_key,required"`
	JSON           merchantIdentifierListResponseMerchantDivisionJSON `json:"-"`
}

// merchantIdentifierListResponseMerchantDivisionJSON contains the JSON metadata
// for the struct [MerchantIdentifierListResponseMerchantDivision]
type merchantIdentifierListResponseMerchantDivisionJSON struct {
	PublishableKey apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MerchantIdentifierListResponseMerchantDivision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r merchantIdentifierListResponseMerchantDivisionJSON) RawJSON() string {
	return r.raw
}
