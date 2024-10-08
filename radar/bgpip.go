// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/apiquery"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// BGPIPService contains methods and other services that help with interacting with
// the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPIPService] method instead.
type BGPIPService struct {
	Options []option.RequestOption
}

// NewBGPIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBGPIPService(opts ...option.RequestOption) (r *BGPIPService) {
	r = &BGPIPService{}
	r.Options = opts
	return
}

// Gets time-series data for the announced IP space count, represented as the
// number of IPv4 /24s and IPv6 /48s, for a given ASN.
func (r *BGPIPService) Timeseries(ctx context.Context, query BGPIPTimeseriesParams, opts ...option.RequestOption) (res *BgpipTimeseriesResponse, err error) {
	var env BgpipTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/ips/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BgpipTimeseriesResponse struct {
	Meta     BgpipTimeseriesResponseMeta     `json:"meta,required"`
	Serie174 BgpipTimeseriesResponseSerie174 `json:"serie_174,required"`
	SerieCn  BgpipTimeseriesResponseSerieCn  `json:"serie_cn,required"`
	JSON     bgpipTimeseriesResponseJSON     `json:"-"`
}

// bgpipTimeseriesResponseJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponse]
type bgpipTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie174    apijson.Field
	SerieCn     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMeta struct {
	DateRange []BgpipTimeseriesResponseMetaDateRange `json:"dateRange,required"`
	JSON      bgpipTimeseriesResponseMetaJSON        `json:"-"`
}

// bgpipTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseMeta]
type bgpipTimeseriesResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      bgpipTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// bgpipTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [BgpipTimeseriesResponseMetaDateRange]
type bgpipTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseSerie174 struct {
	IPV4       []string                            `json:"ipv4,required"`
	IPV6       []string                            `json:"ipv6,required"`
	Timestamps []time.Time                         `json:"timestamps,required" format:"date-time"`
	JSON       bgpipTimeseriesResponseSerie174JSON `json:"-"`
}

// bgpipTimeseriesResponseSerie174JSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseSerie174]
type bgpipTimeseriesResponseSerie174JSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseSerie174) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseSerie174JSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseSerieCn struct {
	IPV4       []string                           `json:"ipv4,required"`
	IPV6       []string                           `json:"ipv6,required"`
	Timestamps []time.Time                        `json:"timestamps,required" format:"date-time"`
	JSON       bgpipTimeseriesResponseSerieCnJSON `json:"-"`
}

// bgpipTimeseriesResponseSerieCnJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseSerieCn]
type bgpipTimeseriesResponseSerieCnJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseSerieCn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseSerieCnJSON) RawJSON() string {
	return r.raw
}

type BGPIPTimeseriesParams struct {
	// Comma separated list of ASNs.
	ASN param.Field[string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[BgpipTimeseriesParamsFormat] `query:"format"`
	// Include data delay meta information
	IncludeDelay param.Field[bool] `query:"includeDelay"`
	// Comma separated list of locations.
	Location param.Field[string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [BGPIPTimeseriesParams]'s query parameters as `url.Values`.
func (r BGPIPTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type BgpipTimeseriesParamsFormat string

const (
	BgpipTimeseriesParamsFormatJson BgpipTimeseriesParamsFormat = "JSON"
	BgpipTimeseriesParamsFormatCsv  BgpipTimeseriesParamsFormat = "CSV"
)

func (r BgpipTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case BgpipTimeseriesParamsFormatJson, BgpipTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type BgpipTimeseriesResponseEnvelope struct {
	Result  BgpipTimeseriesResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    bgpipTimeseriesResponseEnvelopeJSON `json:"-"`
}

// bgpipTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseEnvelope]
type bgpipTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
