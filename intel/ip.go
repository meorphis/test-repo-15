// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

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

// IPService contains methods and other services that help with interacting with
// the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPService] method instead.
type IPService struct {
	Options []option.RequestOption
}

// NewIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPService(opts ...option.RequestOption) (r *IPService) {
	r = &IPService{}
	r.Options = opts
	return
}

// Get IP Overview
func (r *IPService) Get(ctx context.Context, accountID string, query IPGetParams, opts ...option.RequestOption) (res *[]IP, err error) {
	var env IPGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/ip", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IP struct {
	// Specifies a reference to the autonomous systems (AS) that the IP address belongs
	// to.
	BelongsToRef IPBelongsToRef `json:"belongs_to_ref"`
	IP           string         `json:"ip" format:"ipv4"`
	RiskTypes    []interface{}  `json:"risk_types"`
	JSON         ipJSON         `json:"-"`
}

// ipJSON contains the JSON metadata for the struct [IP]
type ipJSON struct {
	BelongsToRef apijson.Field
	IP           apijson.Field
	RiskTypes    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipJSON) RawJSON() string {
	return r.raw
}

// Specifies a reference to the autonomous systems (AS) that the IP address belongs
// to.
type IPBelongsToRef struct {
	ID          string `json:"id"`
	Country     string `json:"country"`
	Description string `json:"description"`
	// Infrastructure type of this ASN.
	Type  IPBelongsToRefType `json:"type"`
	Value string             `json:"value"`
	JSON  ipBelongsToRefJSON `json:"-"`
}

// ipBelongsToRefJSON contains the JSON metadata for the struct [IPBelongsToRef]
type ipBelongsToRefJSON struct {
	ID          apijson.Field
	Country     apijson.Field
	Description apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPBelongsToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipBelongsToRefJSON) RawJSON() string {
	return r.raw
}

// Infrastructure type of this ASN.
type IPBelongsToRefType string

const (
	IPBelongsToRefTypeHostingProvider IPBelongsToRefType = "hosting_provider"
	IPBelongsToRefTypeISP             IPBelongsToRefType = "isp"
	IPBelongsToRefTypeOrganization    IPBelongsToRefType = "organization"
)

func (r IPBelongsToRefType) IsKnown() bool {
	switch r {
	case IPBelongsToRefTypeHostingProvider, IPBelongsToRefTypeISP, IPBelongsToRefTypeOrganization:
		return true
	}
	return false
}

type IPGetParams struct {
	IPV4 param.Field[string] `query:"ipv4"`
	IPV6 param.Field[string] `query:"ipv6"`
}

// URLQuery serializes [IPGetParams]'s query parameters as `url.Values`.
func (r IPGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []IP                  `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IPGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IPGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ipGetResponseEnvelopeJSON       `json:"-"`
}

// ipGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPGetResponseEnvelope]
type ipGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPGetResponseEnvelopeSuccess bool

const (
	IPGetResponseEnvelopeSuccessTrue IPGetResponseEnvelopeSuccess = true
)

func (r IPGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IPGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                             `json:"total_count"`
	JSON       ipGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// ipGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [IPGetResponseEnvelopeResultInfo]
type ipGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
