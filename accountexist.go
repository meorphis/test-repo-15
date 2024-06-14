// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apiquery"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// AccountExistService contains methods and other services that help with
// interacting with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountExistService] method instead.
type AccountExistService struct {
	Options []option.RequestOption
}

// NewAccountExistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountExistService(opts ...option.RequestOption) (r *AccountExistService) {
	r = &AccountExistService{}
	r.Options = opts
	return
}

// Determine whether or not an identifier is associated with an existing Bolt
// account.
func (r *AccountExistService) List(ctx context.Context, params AccountExistListParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "account/exists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type AccountExistListParams struct {
	// A type and value combination that defines the identifier used to detect the
	// existence of an account.
	Identifier      param.Field[AccountExistListParamsIdentifier] `query:"identifier,required"`
	XPublishableKey param.Field[string]                           `header:"X-Publishable-Key,required"`
}

// URLQuery serializes [AccountExistListParams]'s query parameters as `url.Values`.
func (r AccountExistListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A type and value combination that defines the identifier used to detect the
// existence of an account.
type AccountExistListParamsIdentifier struct {
	// The type of identifier
	IdentifierType param.Field[AccountExistListParamsIdentifierIdentifierType] `query:"identifier_type,required"`
	// The value of the identifier. The value must be valid for the specified
	// `identifier_type`
	IdentifierValue param.Field[string] `query:"identifier_value,required"`
}

// URLQuery serializes [AccountExistListParamsIdentifier]'s query parameters as
// `url.Values`.
func (r AccountExistListParamsIdentifier) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of identifier
type AccountExistListParamsIdentifierIdentifierType string

const (
	AccountExistListParamsIdentifierIdentifierTypeEmail       AccountExistListParamsIdentifierIdentifierType = "email"
	AccountExistListParamsIdentifierIdentifierTypeEmailSha256 AccountExistListParamsIdentifierIdentifierType = "email_sha256"
)

func (r AccountExistListParamsIdentifierIdentifierType) IsKnown() bool {
	switch r {
	case AccountExistListParamsIdentifierIdentifierTypeEmail, AccountExistListParamsIdentifierIdentifierTypeEmailSha256:
		return true
	}
	return false
}
