// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package speed

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// AvailabilityService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAvailabilityService] method instead.
type AvailabilityService struct {
	Options []option.RequestOption
}

// NewAvailabilityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAvailabilityService(opts ...option.RequestOption) (r *AvailabilityService) {
	r = &AvailabilityService{}
	r.Options = opts
	return
}

// Retrieves quota for all plans, as well as the current zone quota.
func (r *AvailabilityService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *Availability, err error) {
	var env AvailabilityListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/availabilities", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Availability struct {
	Quota          AvailabilityQuota `json:"quota"`
	Regions        []LabeledRegion   `json:"regions"`
	RegionsPerPlan interface{}       `json:"regionsPerPlan"`
	JSON           availabilityJSON  `json:"-"`
}

// availabilityJSON contains the JSON metadata for the struct [Availability]
type availabilityJSON struct {
	Quota          apijson.Field
	Regions        apijson.Field
	RegionsPerPlan apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Availability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityJSON) RawJSON() string {
	return r.raw
}

type AvailabilityQuota struct {
	// Cloudflare plan.
	Plan string `json:"plan"`
	// The number of tests available per plan.
	QuotasPerPlan map[string]float64 `json:"quotasPerPlan"`
	// The number of remaining schedules available.
	RemainingSchedules float64 `json:"remainingSchedules"`
	// The number of remaining tests available.
	RemainingTests float64 `json:"remainingTests"`
	// The number of schedules available per plan.
	ScheduleQuotasPerPlan map[string]float64    `json:"scheduleQuotasPerPlan"`
	JSON                  availabilityQuotaJSON `json:"-"`
}

// availabilityQuotaJSON contains the JSON metadata for the struct
// [AvailabilityQuota]
type availabilityQuotaJSON struct {
	Plan                  apijson.Field
	QuotasPerPlan         apijson.Field
	RemainingSchedules    apijson.Field
	RemainingTests        apijson.Field
	ScheduleQuotasPerPlan apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AvailabilityQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityQuotaJSON) RawJSON() string {
	return r.raw
}

type AvailabilityListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                                 `json:"success,required"`
	Result  Availability                         `json:"result"`
	JSON    availabilityListResponseEnvelopeJSON `json:"-"`
}

// availabilityListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AvailabilityListResponseEnvelope]
type availabilityListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
