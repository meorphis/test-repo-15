// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/apiquery"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/tidwall/gjson"
)

// ConfigurationService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigurationService] method instead.
type ConfigurationService struct {
	Options []option.RequestOption
}

// NewConfigurationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigurationService(opts ...option.RequestOption) (r *ConfigurationService) {
	r = &ConfigurationService{}
	r.Options = opts
	return
}

// Set configuration properties
func (r *ConfigurationService) Update(ctx context.Context, zoneID string, body ConfigurationUpdateParams, opts ...option.RequestOption) (res *ConfigurationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/configuration", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve information about specific configuration properties
func (r *ConfigurationService) Get(ctx context.Context, zoneID string, query ConfigurationGetParams, opts ...option.RequestOption) (res *Configuration, err error) {
	var env ConfigurationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/configuration", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Configuration struct {
	AuthIDCharacteristics []ConfigurationAuthIDCharacteristic `json:"auth_id_characteristics,required"`
	JSON                  configurationJSON                   `json:"-"`
}

// configurationJSON contains the JSON metadata for the struct [Configuration]
type configurationJSON struct {
	AuthIDCharacteristics apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationJSON) RawJSON() string {
	return r.raw
}

// Auth ID Characteristic
type ConfigurationAuthIDCharacteristic struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name string `json:"name,required"`
	// The type of characteristic.
	Type  ConfigurationAuthIDCharacteristicsType `json:"type,required"`
	JSON  configurationAuthIDCharacteristicJSON  `json:"-"`
	union ConfigurationAuthIDCharacteristicsUnion
}

// configurationAuthIDCharacteristicJSON contains the JSON metadata for the struct
// [ConfigurationAuthIDCharacteristic]
type configurationAuthIDCharacteristicJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configurationAuthIDCharacteristicJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigurationAuthIDCharacteristic) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigurationAuthIDCharacteristic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigurationAuthIDCharacteristicsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [api_gateway.ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic],
// [api_gateway.ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim].
func (r ConfigurationAuthIDCharacteristic) AsUnion() ConfigurationAuthIDCharacteristicsUnion {
	return r.union
}

// Auth ID Characteristic
//
// Union satisfied by
// [api_gateway.ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic] or
// [api_gateway.ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim].
type ConfigurationAuthIDCharacteristicsUnion interface {
	implementsAPIGatewayConfigurationAuthIDCharacteristic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationAuthIDCharacteristicsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim{}),
		},
	)
}

// Auth ID Characteristic
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name string `json:"name,required"`
	// The type of characteristic.
	Type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType `json:"type,required"`
	JSON configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJSON `json:"-"`
}

// configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJSON contains the
// JSON metadata for the struct
// [ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic]
type configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic) implementsAPIGatewayConfigurationAuthIDCharacteristic() {
}

// The type of characteristic.
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType string

const (
	ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicTypeHeader ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType = "header"
	ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicTypeCookie ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType = "cookie"
)

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType) IsKnown() bool {
	switch r {
	case ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicTypeHeader, ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicTypeCookie:
		return true
	}
	return false
}

// Auth ID Characteristic extracted from JWT Token Claims
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim struct {
	// Claim location expressed as `$(token_config_id):$(json_path)`, where
	// `token_config_id` is the ID of the token configuration used in validating the
	// JWT, and `json_path` is a RFC 9535 JSONPath
	// (https://goessner.net/articles/JsonPath/,
	// https://www.rfc-editor.org/rfc/rfc9535.html). The JSONPath expression may be in
	// dot or bracket notation, may only specify literal keys or array indexes, and
	// must return a singleton value, which will be interpreted as a string.
	Name string `json:"name,required"`
	// The type of characteristic.
	Type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimType `json:"type,required"`
	JSON configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimJSON `json:"-"`
}

// configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimJSON
// contains the JSON metadata for the struct
// [ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim]
type configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim) implementsAPIGatewayConfigurationAuthIDCharacteristic() {
}

// The type of characteristic.
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimType string

const (
	ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimTypeJwt ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimType = "jwt"
)

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimType) IsKnown() bool {
	switch r {
	case ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimTypeJwt:
		return true
	}
	return false
}

// The type of characteristic.
type ConfigurationAuthIDCharacteristicsType string

const (
	ConfigurationAuthIDCharacteristicsTypeHeader ConfigurationAuthIDCharacteristicsType = "header"
	ConfigurationAuthIDCharacteristicsTypeCookie ConfigurationAuthIDCharacteristicsType = "cookie"
	ConfigurationAuthIDCharacteristicsTypeJwt    ConfigurationAuthIDCharacteristicsType = "jwt"
)

func (r ConfigurationAuthIDCharacteristicsType) IsKnown() bool {
	switch r {
	case ConfigurationAuthIDCharacteristicsTypeHeader, ConfigurationAuthIDCharacteristicsTypeCookie, ConfigurationAuthIDCharacteristicsTypeJwt:
		return true
	}
	return false
}

type ConfigurationParam struct {
	AuthIDCharacteristics param.Field[[]ConfigurationAuthIDCharacteristicsUnionParam] `json:"auth_id_characteristics,required"`
}

func (r ConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Auth ID Characteristic
type ConfigurationAuthIDCharacteristicParam struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name param.Field[string] `json:"name,required"`
	// The type of characteristic.
	Type param.Field[ConfigurationAuthIDCharacteristicsType] `json:"type,required"`
}

func (r ConfigurationAuthIDCharacteristicParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationAuthIDCharacteristicParam) implementsAPIGatewayConfigurationAuthIDCharacteristicsUnionParam() {
}

// Auth ID Characteristic
//
// Satisfied by
// [api_gateway.ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicParam],
// [api_gateway.ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimParam],
// [ConfigurationAuthIDCharacteristicParam].
type ConfigurationAuthIDCharacteristicsUnionParam interface {
	implementsAPIGatewayConfigurationAuthIDCharacteristicsUnionParam()
}

// Auth ID Characteristic
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicParam struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name param.Field[string] `json:"name,required"`
	// The type of characteristic.
	Type param.Field[ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType] `json:"type,required"`
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicParam) implementsAPIGatewayConfigurationAuthIDCharacteristicsUnionParam() {
}

// Auth ID Characteristic extracted from JWT Token Claims
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimParam struct {
	// Claim location expressed as `$(token_config_id):$(json_path)`, where
	// `token_config_id` is the ID of the token configuration used in validating the
	// JWT, and `json_path` is a RFC 9535 JSONPath
	// (https://goessner.net/articles/JsonPath/,
	// https://www.rfc-editor.org/rfc/rfc9535.html). The JSONPath expression may be in
	// dot or bracket notation, may only specify literal keys or array indexes, and
	// must return a singleton value, which will be interpreted as a string.
	Name param.Field[string] `json:"name,required"`
	// The type of characteristic.
	Type param.Field[ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimType] `json:"type,required"`
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimParam) implementsAPIGatewayConfigurationAuthIDCharacteristicsUnionParam() {
}

type ConfigurationUpdateResponse struct {
	Errors   Message `json:"errors,required"`
	Messages Message `json:"messages,required"`
	// Whether the API call was successful
	Success ConfigurationUpdateResponseSuccess `json:"success,required"`
	JSON    configurationUpdateResponseJSON    `json:"-"`
}

// configurationUpdateResponseJSON contains the JSON metadata for the struct
// [ConfigurationUpdateResponse]
type configurationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigurationUpdateResponseSuccess bool

const (
	ConfigurationUpdateResponseSuccessTrue ConfigurationUpdateResponseSuccess = true
)

func (r ConfigurationUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ConfigurationUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ConfigurationUpdateParams struct {
	Configuration ConfigurationParam `json:"configuration,required"`
}

func (r ConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Configuration)
}

type ConfigurationGetParams struct {
	// Requests information about certain properties.
	Properties param.Field[[]ConfigurationGetParamsProperty] `query:"properties"`
}

// URLQuery serializes [ConfigurationGetParams]'s query parameters as `url.Values`.
func (r ConfigurationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ConfigurationGetParamsProperty string

const (
	ConfigurationGetParamsPropertyAuthIDCharacteristics ConfigurationGetParamsProperty = "auth_id_characteristics"
)

func (r ConfigurationGetParamsProperty) IsKnown() bool {
	switch r {
	case ConfigurationGetParamsPropertyAuthIDCharacteristics:
		return true
	}
	return false
}

type ConfigurationGetResponseEnvelope struct {
	Errors   Message       `json:"errors,required"`
	Messages Message       `json:"messages,required"`
	Result   Configuration `json:"result,required"`
	// Whether the API call was successful
	Success ConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    configurationGetResponseEnvelopeJSON    `json:"-"`
}

// configurationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigurationGetResponseEnvelope]
type configurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigurationGetResponseEnvelopeSuccess bool

const (
	ConfigurationGetResponseEnvelopeSuccessTrue ConfigurationGetResponseEnvelopeSuccess = true
)

func (r ConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
