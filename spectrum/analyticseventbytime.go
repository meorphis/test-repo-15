// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/apiquery"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// AnalyticsEventBytimeService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsEventBytimeService] method instead.
type AnalyticsEventBytimeService struct {
	Options []option.RequestOption
}

// NewAnalyticsEventBytimeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsEventBytimeService(opts ...option.RequestOption) (r *AnalyticsEventBytimeService) {
	r = &AnalyticsEventBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
func (r *AnalyticsEventBytimeService) Get(ctx context.Context, zone string, query AnalyticsEventBytimeGetParams, opts ...option.RequestOption) (res *AnalyticsEventBytimeGetResponseUnion, err error) {
	var env AnalyticsEventBytimeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zone == "" {
		err = errors.New("missing required zone parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/analytics/events/bytime", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [spectrum.AnalyticsEventBytimeGetResponseUnknown] or
// [shared.UnionString].
type AnalyticsEventBytimeGetResponseUnion interface {
	ImplementsSpectrumAnalyticsEventBytimeGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsEventBytimeGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AnalyticsEventBytimeGetParams struct {
	// Can be used to break down the data by given attributes. Options are:
	//
	// | Dimension | Name                          | Example                                                    |
	// | --------- | ----------------------------- | ---------------------------------------------------------- |
	// | event     | Connection Event              | connect, progress, disconnect, originError, clientFiltered |
	// | appID     | Application ID                | 40d67c87c6cd4b889a4fd57805225e85                           |
	// | coloName  | Colo Name                     | SFO                                                        |
	// | ipVersion | IP version used by the client | 4, 6.                                                      |
	Dimensions param.Field[[]Dimension] `query:"dimensions"`
	// Used to filter rows by one or more dimensions. Filters can be combined using OR
	// and AND boolean logic. AND takes precedence over OR in all the expressions. The
	// OR operator is defined using a comma (,) or OR keyword surrounded by whitespace.
	// The AND operator is defined using a semicolon (;) or AND keyword surrounded by
	// whitespace. Note that the semicolon is a reserved character in URLs (rfc1738)
	// and needs to be percent-encoded as %3B. Comparison options are:
	//
	// | Operator | Name                     | URL Encoded |
	// | -------- | ------------------------ | ----------- |
	// | ==       | Equals                   | %3D%3D      |
	// | !=       | Does not equals          | !%3D        |
	// | >        | Greater Than             | %3E         |
	// | <        | Less Than                | %3C         |
	// | >=       | Greater than or equal to | %3E%3D      |
	// | <=       | Less than or equal to    | %3C%3D .    |
	Filters param.Field[string] `query:"filters"`
	// One or more metrics to compute. Options are:
	//
	// | Metric         | Name                                | Example | Unit                  |
	// | -------------- | ----------------------------------- | ------- | --------------------- |
	// | count          | Count of total events               | 1000    | Count                 |
	// | bytesIngress   | Sum of ingress bytes                | 1000    | Sum                   |
	// | bytesEgress    | Sum of egress bytes                 | 1000    | Sum                   |
	// | durationAvg    | Average connection duration         | 1.0     | Time in milliseconds  |
	// | durationMedian | Median connection duration          | 1.0     | Time in milliseconds  |
	// | duration90th   | 90th percentile connection duration | 1.0     | Time in milliseconds  |
	// | duration99th   | 99th percentile connection duration | 1.0     | Time in milliseconds. |
	Metrics param.Field[[]AnalyticsEventBytimeGetParamsMetric] `query:"metrics"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The sort order for the result set; sort fields must be included in `metrics` or
	// `dimensions`.
	Sort param.Field[[]interface{}] `query:"sort"`
	// Used to select time series resolution.
	TimeDelta param.Field[AnalyticsEventBytimeGetParamsTimeDelta] `query:"time_delta"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AnalyticsEventBytimeGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticsEventBytimeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AnalyticsEventBytimeGetParamsMetric string

const (
	AnalyticsEventBytimeGetParamsMetricCount          AnalyticsEventBytimeGetParamsMetric = "count"
	AnalyticsEventBytimeGetParamsMetricBytesIngress   AnalyticsEventBytimeGetParamsMetric = "bytesIngress"
	AnalyticsEventBytimeGetParamsMetricBytesEgress    AnalyticsEventBytimeGetParamsMetric = "bytesEgress"
	AnalyticsEventBytimeGetParamsMetricDurationAvg    AnalyticsEventBytimeGetParamsMetric = "durationAvg"
	AnalyticsEventBytimeGetParamsMetricDurationMedian AnalyticsEventBytimeGetParamsMetric = "durationMedian"
	AnalyticsEventBytimeGetParamsMetricDuration90th   AnalyticsEventBytimeGetParamsMetric = "duration90th"
	AnalyticsEventBytimeGetParamsMetricDuration99th   AnalyticsEventBytimeGetParamsMetric = "duration99th"
)

func (r AnalyticsEventBytimeGetParamsMetric) IsKnown() bool {
	switch r {
	case AnalyticsEventBytimeGetParamsMetricCount, AnalyticsEventBytimeGetParamsMetricBytesIngress, AnalyticsEventBytimeGetParamsMetricBytesEgress, AnalyticsEventBytimeGetParamsMetricDurationAvg, AnalyticsEventBytimeGetParamsMetricDurationMedian, AnalyticsEventBytimeGetParamsMetricDuration90th, AnalyticsEventBytimeGetParamsMetricDuration99th:
		return true
	}
	return false
}

// Used to select time series resolution.
type AnalyticsEventBytimeGetParamsTimeDelta string

const (
	AnalyticsEventBytimeGetParamsTimeDeltaYear       AnalyticsEventBytimeGetParamsTimeDelta = "year"
	AnalyticsEventBytimeGetParamsTimeDeltaQuarter    AnalyticsEventBytimeGetParamsTimeDelta = "quarter"
	AnalyticsEventBytimeGetParamsTimeDeltaMonth      AnalyticsEventBytimeGetParamsTimeDelta = "month"
	AnalyticsEventBytimeGetParamsTimeDeltaWeek       AnalyticsEventBytimeGetParamsTimeDelta = "week"
	AnalyticsEventBytimeGetParamsTimeDeltaDay        AnalyticsEventBytimeGetParamsTimeDelta = "day"
	AnalyticsEventBytimeGetParamsTimeDeltaHour       AnalyticsEventBytimeGetParamsTimeDelta = "hour"
	AnalyticsEventBytimeGetParamsTimeDeltaDekaminute AnalyticsEventBytimeGetParamsTimeDelta = "dekaminute"
	AnalyticsEventBytimeGetParamsTimeDeltaMinute     AnalyticsEventBytimeGetParamsTimeDelta = "minute"
)

func (r AnalyticsEventBytimeGetParamsTimeDelta) IsKnown() bool {
	switch r {
	case AnalyticsEventBytimeGetParamsTimeDeltaYear, AnalyticsEventBytimeGetParamsTimeDeltaQuarter, AnalyticsEventBytimeGetParamsTimeDeltaMonth, AnalyticsEventBytimeGetParamsTimeDeltaWeek, AnalyticsEventBytimeGetParamsTimeDeltaDay, AnalyticsEventBytimeGetParamsTimeDeltaHour, AnalyticsEventBytimeGetParamsTimeDeltaDekaminute, AnalyticsEventBytimeGetParamsTimeDeltaMinute:
		return true
	}
	return false
}

type AnalyticsEventBytimeGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                `json:"errors,required"`
	Messages []shared.ResponseInfo                `json:"messages,required"`
	Result   AnalyticsEventBytimeGetResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AnalyticsEventBytimeGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    analyticsEventBytimeGetResponseEnvelopeJSON    `json:"-"`
}

// analyticsEventBytimeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnalyticsEventBytimeGetResponseEnvelope]
type analyticsEventBytimeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventBytimeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventBytimeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AnalyticsEventBytimeGetResponseEnvelopeSuccess bool

const (
	AnalyticsEventBytimeGetResponseEnvelopeSuccessTrue AnalyticsEventBytimeGetResponseEnvelopeSuccess = true
)

func (r AnalyticsEventBytimeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyticsEventBytimeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
