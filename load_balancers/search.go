// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/apiquery"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// SearchService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r *SearchService) {
	r = &SearchService{}
	r.Options = opts
	return
}

// Search for Load Balancing resources.
func (r *SearchService) Get(ctx context.Context, accountID string, query SearchGetParams, opts ...option.RequestOption) (res *[]SearchGetResponse, err error) {
	var env SearchGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/search", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SearchGetResponse = interface{}

type SearchGetParams struct {
	Page         param.Field[interface{}]                 `query:"page"`
	PerPage      param.Field[interface{}]                 `query:"per_page"`
	SearchParams param.Field[SearchGetParamsSearchParams] `query:"search_params"`
}

// URLQuery serializes [SearchGetParams]'s query parameters as `url.Values`.
func (r SearchGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SearchGetParamsSearchParams struct {
	// Search query term.
	Query param.Field[string] `query:"query"`
	// The type of references to include ("\*" for all).
	References param.Field[SearchGetParamsSearchParamsReferences] `query:"references"`
}

// URLQuery serializes [SearchGetParamsSearchParams]'s query parameters as
// `url.Values`.
func (r SearchGetParamsSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The type of references to include ("\*" for all).
type SearchGetParamsSearchParamsReferences string

const (
	SearchGetParamsSearchParamsReferencesEmpty    SearchGetParamsSearchParamsReferences = ""
	SearchGetParamsSearchParamsReferencesStar     SearchGetParamsSearchParamsReferences = "*"
	SearchGetParamsSearchParamsReferencesReferral SearchGetParamsSearchParamsReferences = "referral"
	SearchGetParamsSearchParamsReferencesReferrer SearchGetParamsSearchParamsReferences = "referrer"
)

func (r SearchGetParamsSearchParamsReferences) IsKnown() bool {
	switch r {
	case SearchGetParamsSearchParamsReferencesEmpty, SearchGetParamsSearchParamsReferencesStar, SearchGetParamsSearchParamsReferencesReferral, SearchGetParamsSearchParamsReferencesReferrer:
		return true
	}
	return false
}

type SearchGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SearchGetResponse   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SearchGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SearchGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       searchGetResponseEnvelopeJSON       `json:"-"`
}

// searchGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SearchGetResponseEnvelope]
type searchGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SearchGetResponseEnvelopeSuccess bool

const (
	SearchGetResponseEnvelopeSuccessTrue SearchGetResponseEnvelopeSuccess = true
)

func (r SearchGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SearchGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SearchGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       searchGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// searchGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [SearchGetResponseEnvelopeResultInfo]
type searchGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
