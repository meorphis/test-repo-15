// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// SettingSchemaValidationService contains methods and other services that help
// with interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingSchemaValidationService] method instead.
type SettingSchemaValidationService struct {
	Options []option.RequestOption
}

// NewSettingSchemaValidationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingSchemaValidationService(opts ...option.RequestOption) (r *SettingSchemaValidationService) {
	r = &SettingSchemaValidationService{}
	r.Options = opts
	return
}

// Updates zone level schema validation settings on the zone
func (r *SettingSchemaValidationService) Update(ctx context.Context, zoneID string, body SettingSchemaValidationUpdateParams, opts ...option.RequestOption) (res *Settings, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Updates zone level schema validation settings on the zone
func (r *SettingSchemaValidationService) Edit(ctx context.Context, zoneID string, body SettingSchemaValidationEditParams, opts ...option.RequestOption) (res *Settings, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieves zone level schema validation settings currently set on the zone
func (r *SettingSchemaValidationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *Settings, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SettingSchemaValidationUpdateParams struct {
	// The default mitigation action used when there is no mitigation action defined on
	// the operation
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for
	// the request when there is no mitigation action defined on the operation
	ValidationDefaultMitigationAction param.Field[SettingSchemaValidationUpdateParamsValidationDefaultMitigationAction] `json:"validation_default_mitigation_action,required"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	// - `null` indicates that no override is in place
	//
	// To clear any override, use the special value `disable_override` or `null`
	ValidationOverrideMitigationAction param.Field[SettingSchemaValidationUpdateParamsValidationOverrideMitigationAction] `json:"validation_override_mitigation_action"`
}

func (r SettingSchemaValidationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default mitigation action used when there is no mitigation action defined on
// the operation
//
// Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
//
// A special value of of `none` will skip running schema validation entirely for
// the request when there is no mitigation action defined on the operation
type SettingSchemaValidationUpdateParamsValidationDefaultMitigationAction string

const (
	SettingSchemaValidationUpdateParamsValidationDefaultMitigationActionNone  SettingSchemaValidationUpdateParamsValidationDefaultMitigationAction = "none"
	SettingSchemaValidationUpdateParamsValidationDefaultMitigationActionLog   SettingSchemaValidationUpdateParamsValidationDefaultMitigationAction = "log"
	SettingSchemaValidationUpdateParamsValidationDefaultMitigationActionBlock SettingSchemaValidationUpdateParamsValidationDefaultMitigationAction = "block"
)

func (r SettingSchemaValidationUpdateParamsValidationDefaultMitigationAction) IsKnown() bool {
	switch r {
	case SettingSchemaValidationUpdateParamsValidationDefaultMitigationActionNone, SettingSchemaValidationUpdateParamsValidationDefaultMitigationActionLog, SettingSchemaValidationUpdateParamsValidationDefaultMitigationActionBlock:
		return true
	}
	return false
}

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
// - `null` indicates that no override is in place
//
// To clear any override, use the special value `disable_override` or `null`
type SettingSchemaValidationUpdateParamsValidationOverrideMitigationAction string

const (
	SettingSchemaValidationUpdateParamsValidationOverrideMitigationActionNone            SettingSchemaValidationUpdateParamsValidationOverrideMitigationAction = "none"
	SettingSchemaValidationUpdateParamsValidationOverrideMitigationActionDisableOverride SettingSchemaValidationUpdateParamsValidationOverrideMitigationAction = "disable_override"
)

func (r SettingSchemaValidationUpdateParamsValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case SettingSchemaValidationUpdateParamsValidationOverrideMitigationActionNone, SettingSchemaValidationUpdateParamsValidationOverrideMitigationActionDisableOverride:
		return true
	}
	return false
}

type SettingSchemaValidationEditParams struct {
	// The default mitigation action used when there is no mitigation action defined on
	// the operation Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for
	// the request when there is no mitigation action defined on the operation
	//
	// `null` will have no effect.
	ValidationDefaultMitigationAction param.Field[SettingSchemaValidationEditParamsValidationDefaultMitigationAction] `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	//
	// To clear any override, use the special value `disable_override`
	//
	// `null` will have no effect.
	ValidationOverrideMitigationAction param.Field[SettingSchemaValidationEditParamsValidationOverrideMitigationAction] `json:"validation_override_mitigation_action"`
}

func (r SettingSchemaValidationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default mitigation action used when there is no mitigation action defined on
// the operation Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
//
// A special value of of `none` will skip running schema validation entirely for
// the request when there is no mitigation action defined on the operation
//
// `null` will have no effect.
type SettingSchemaValidationEditParamsValidationDefaultMitigationAction string

const (
	SettingSchemaValidationEditParamsValidationDefaultMitigationActionNone  SettingSchemaValidationEditParamsValidationDefaultMitigationAction = "none"
	SettingSchemaValidationEditParamsValidationDefaultMitigationActionLog   SettingSchemaValidationEditParamsValidationDefaultMitigationAction = "log"
	SettingSchemaValidationEditParamsValidationDefaultMitigationActionBlock SettingSchemaValidationEditParamsValidationDefaultMitigationAction = "block"
)

func (r SettingSchemaValidationEditParamsValidationDefaultMitigationAction) IsKnown() bool {
	switch r {
	case SettingSchemaValidationEditParamsValidationDefaultMitigationActionNone, SettingSchemaValidationEditParamsValidationDefaultMitigationActionLog, SettingSchemaValidationEditParamsValidationDefaultMitigationActionBlock:
		return true
	}
	return false
}

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
//
// To clear any override, use the special value `disable_override`
//
// `null` will have no effect.
type SettingSchemaValidationEditParamsValidationOverrideMitigationAction string

const (
	SettingSchemaValidationEditParamsValidationOverrideMitigationActionNone            SettingSchemaValidationEditParamsValidationOverrideMitigationAction = "none"
	SettingSchemaValidationEditParamsValidationOverrideMitigationActionDisableOverride SettingSchemaValidationEditParamsValidationOverrideMitigationAction = "disable_override"
)

func (r SettingSchemaValidationEditParamsValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case SettingSchemaValidationEditParamsValidationOverrideMitigationActionNone, SettingSchemaValidationEditParamsValidationOverrideMitigationActionDisableOverride:
		return true
	}
	return false
}
