// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// DLPProfilePredefinedService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfilePredefinedService] method instead.
type DLPProfilePredefinedService struct {
	Options []option.RequestOption
}

// NewDLPProfilePredefinedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPProfilePredefinedService(opts ...option.RequestOption) (r *DLPProfilePredefinedService) {
	r = &DLPProfilePredefinedService{}
	r.Options = opts
	return
}

// Updates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *DLPProfilePredefinedService) Update(ctx context.Context, accountID string, profileID string, body DLPProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *PredefinedProfile, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches a predefined DLP profile.
func (r *DLPProfilePredefinedService) Get(ctx context.Context, accountID string, profileID string, opts ...option.RequestOption) (res *PredefinedProfile, err error) {
	var env DLPProfilePredefinedGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PredefinedProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// The entries for this profile.
	Entries []PredefinedProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// If true, scan images via OCR to determine if any text present matches filters.
	OCREnabled bool `json:"ocr_enabled"`
	// The type of the profile.
	Type PredefinedProfileType `json:"type"`
	JSON predefinedProfileJSON `json:"-"`
}

// predefinedProfileJSON contains the JSON metadata for the struct
// [PredefinedProfile]
type predefinedProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	OCREnabled        apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r PredefinedProfile) implementsZeroTrustProfile() {}

func (r PredefinedProfile) implementsZeroTrustDLPProfileGetResponse() {}

// A predefined entry that matches a profile
type PredefinedProfileEntry struct {
	// The ID for this entry
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                `json:"profile_id"`
	JSON      predefinedProfileEntryJSON `json:"-"`
}

// predefinedProfileEntryJSON contains the JSON metadata for the struct
// [PredefinedProfileEntry]
type predefinedProfileEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type PredefinedProfileType string

const (
	PredefinedProfileTypePredefined PredefinedProfileType = "predefined"
)

func (r PredefinedProfileType) IsKnown() bool {
	switch r {
	case PredefinedProfileTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateParams struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The entries for this profile.
	Entries param.Field[[]DLPProfilePredefinedUpdateParamsEntry] `json:"entries"`
	// If true, scan images via OCR to determine if any text present matches filters.
	OCREnabled param.Field[bool] `json:"ocr_enabled"`
}

func (r DLPProfilePredefinedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedUpdateParamsEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DLPProfilePredefinedUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   PredefinedProfile     `json:"result,required"`
	// Whether the API call was successful
	Success DLPProfilePredefinedGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpProfilePredefinedGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedGetResponseEnvelope]
type dlpProfilePredefinedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfilePredefinedGetResponseEnvelopeSuccess bool

const (
	DLPProfilePredefinedGetResponseEnvelopeSuccessTrue DLPProfilePredefinedGetResponseEnvelopeSuccess = true
)

func (r DLPProfilePredefinedGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
