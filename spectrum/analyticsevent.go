// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

import (
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// AnalyticsEventService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsEventService] method instead.
type AnalyticsEventService struct {
	Options   []option.RequestOption
	Bytimes   *AnalyticsEventBytimeService
	Summaries *AnalyticsEventSummaryService
}

// NewAnalyticsEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsEventService(opts ...option.RequestOption) (r *AnalyticsEventService) {
	r = &AnalyticsEventService{}
	r.Options = opts
	r.Bytimes = NewAnalyticsEventBytimeService(opts...)
	r.Summaries = NewAnalyticsEventSummaryService(opts...)
	return
}

type Dimension string

const (
	DimensionEvent     Dimension = "event"
	DimensionAppID     Dimension = "appID"
	DimensionColoName  Dimension = "coloName"
	DimensionIPVersion Dimension = "ipVersion"
)

func (r Dimension) IsKnown() bool {
	switch r {
	case DimensionEvent, DimensionAppID, DimensionColoName, DimensionIPVersion:
		return true
	}
	return false
}