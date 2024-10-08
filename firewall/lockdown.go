// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

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
	"github.com/stainless-sdks/meorphis-test-40-go/internal/pagination"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// LockdownService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLockdownService] method instead.
type LockdownService struct {
	Options []option.RequestOption
}

// NewLockdownService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLockdownService(opts ...option.RequestOption) (r *LockdownService) {
	r = &LockdownService{}
	r.Options = opts
	return
}

// Creates a new Zone Lockdown rule.
func (r *LockdownService) New(ctx context.Context, zoneIdentifier string, body LockdownNewParams, opts ...option.RequestOption) (res *Lockdown, err error) {
	var env LockdownNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Zone Lockdown rule.
func (r *LockdownService) Update(ctx context.Context, zoneIdentifier string, id string, body LockdownUpdateParams, opts ...option.RequestOption) (res *Lockdown, err error) {
	var env LockdownUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *LockdownService) List(ctx context.Context, zoneIdentifier string, query LockdownListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Lockdown], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *LockdownService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query LockdownListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Lockdown] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing Zone Lockdown rule.
func (r *LockdownService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *LockdownDeleteResponse, err error) {
	var env LockdownDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a Zone Lockdown rule.
func (r *LockdownService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *Lockdown, err error) {
	var env LockdownGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
type Configuration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target ConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string            `json:"value"`
	JSON  configurationJSON `json:"-"`
	union ConfigurationUnion
}

// configurationJSON contains the JSON metadata for the struct [Configuration]
type configurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configurationJSON) RawJSON() string {
	return r.raw
}

func (r *Configuration) UnmarshalJSON(data []byte) (err error) {
	*r = Configuration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigurationUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [firewall.LockdownIPConfiguration],
// [firewall.LockdownCIDRConfiguration].
func (r Configuration) AsUnion() ConfigurationUnion {
	return r.union
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by [firewall.LockdownIPConfiguration] or
// [firewall.LockdownCIDRConfiguration].
type ConfigurationUnion interface {
	implementsFirewallConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownCIDRConfiguration{}),
		},
	)
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type ConfigurationTarget string

const (
	ConfigurationTargetIP      ConfigurationTarget = "ip"
	ConfigurationTargetIPRange ConfigurationTarget = "ip_range"
)

func (r ConfigurationTarget) IsKnown() bool {
	switch r {
	case ConfigurationTargetIP, ConfigurationTargetIPRange:
		return true
	}
	return false
}

type Lockdown struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations Configuration `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []LockdownURL `json:"urls,required"`
	JSON lockdownJSON  `json:"-"`
}

// lockdownJSON contains the JSON metadata for the struct [Lockdown]
type lockdownJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Lockdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownJSON) RawJSON() string {
	return r.raw
}

type LockdownCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target LockdownCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                        `json:"value"`
	JSON  lockdownCIDRConfigurationJSON `json:"-"`
}

// lockdownCIDRConfigurationJSON contains the JSON metadata for the struct
// [LockdownCIDRConfiguration]
type lockdownCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownCIDRConfiguration) implementsFirewallConfiguration() {}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type LockdownCIDRConfigurationTarget string

const (
	LockdownCIDRConfigurationTargetIPRange LockdownCIDRConfigurationTarget = "ip_range"
)

func (r LockdownCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type LockdownIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LockdownIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                      `json:"value"`
	JSON  lockdownIPConfigurationJSON `json:"-"`
}

// lockdownIPConfigurationJSON contains the JSON metadata for the struct
// [LockdownIPConfiguration]
type lockdownIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownIPConfiguration) implementsFirewallConfiguration() {}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LockdownIPConfigurationTarget string

const (
	LockdownIPConfigurationTargetIP LockdownIPConfigurationTarget = "ip"
)

func (r LockdownIPConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownIPConfigurationTargetIP:
		return true
	}
	return false
}

type LockdownURL = string

type LockdownDeleteResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID   string                     `json:"id"`
	JSON lockdownDeleteResponseJSON `json:"-"`
}

// lockdownDeleteResponseJSON contains the JSON metadata for the struct
// [LockdownDeleteResponse]
type lockdownDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type LockdownNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r LockdownNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LockdownNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Lockdown              `json:"result,required"`
	// Whether the API call was successful
	Success LockdownNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownNewResponseEnvelopeJSON    `json:"-"`
}

// lockdownNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownNewResponseEnvelope]
type lockdownNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LockdownNewResponseEnvelopeSuccess bool

const (
	LockdownNewResponseEnvelopeSuccessTrue LockdownNewResponseEnvelopeSuccess = true
)

func (r LockdownNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LockdownUpdateParams struct {
	Body interface{} `json:"body,required"`
}

func (r LockdownUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LockdownUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Lockdown              `json:"result,required"`
	// Whether the API call was successful
	Success LockdownUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownUpdateResponseEnvelopeJSON    `json:"-"`
}

// lockdownUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownUpdateResponseEnvelope]
type lockdownUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LockdownUpdateResponseEnvelopeSuccess bool

const (
	LockdownUpdateResponseEnvelopeSuccessTrue LockdownUpdateResponseEnvelopeSuccess = true
)

func (r LockdownUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LockdownListParams struct {
	// The timestamp of when the rule was created.
	CreatedOn param.Field[time.Time] `query:"created_on" format:"date-time"`
	// A string to search for in the description of existing rules.
	Description param.Field[string] `query:"description"`
	// A string to search for in the description of existing rules.
	DescriptionSearch param.Field[string] `query:"description_search"`
	// A single IP address to search for in existing rules.
	IP param.Field[string] `query:"ip"`
	// A single IP address range to search for in existing rules.
	IPRangeSearch param.Field[string] `query:"ip_range_search"`
	// A single IP address to search for in existing rules.
	IPSearch param.Field[string] `query:"ip_search"`
	// The timestamp of when the rule was last modified.
	ModifiedOn param.Field[time.Time] `query:"modified_on" format:"date-time"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
	// The priority of the rule to control the processing order. A lower number
	// indicates higher priority. If not provided, any rules with a configured priority
	// will be processed before rules without a priority.
	Priority param.Field[float64] `query:"priority"`
	// A single URI to search for in the list of URLs of existing rules.
	URISearch param.Field[string] `query:"uri_search"`
}

// URLQuery serializes [LockdownListParams]'s query parameters as `url.Values`.
func (r LockdownListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LockdownDeleteResponseEnvelope struct {
	Result LockdownDeleteResponse             `json:"result"`
	JSON   lockdownDeleteResponseEnvelopeJSON `json:"-"`
}

// lockdownDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownDeleteResponseEnvelope]
type lockdownDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LockdownGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Lockdown              `json:"result,required"`
	// Whether the API call was successful
	Success LockdownGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownGetResponseEnvelopeJSON    `json:"-"`
}

// lockdownGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownGetResponseEnvelope]
type lockdownGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LockdownGetResponseEnvelopeSuccess bool

const (
	LockdownGetResponseEnvelopeSuccessTrue LockdownGetResponseEnvelopeSuccess = true
)

func (r LockdownGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
