// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// WebhookService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

// Creates a webhook notification.
func (r *WebhookService) Update(ctx context.Context, accountID string, body WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponseUnion, err error) {
	var env WebhookUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a webhook.
func (r *WebhookService) Delete(ctx context.Context, accountID string, opts ...option.RequestOption) (res *WebhookDeleteResponseUnion, err error) {
	var env WebhookDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a list of webhooks.
func (r *WebhookService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *WebhookGetResponseUnion, err error) {
	var env WebhookGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [stream.WebhookUpdateResponseUnknown] or
// [shared.UnionString].
type WebhookUpdateResponseUnion interface {
	ImplementsStreamWebhookUpdateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WebhookUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [stream.WebhookDeleteResponseUnknown] or
// [shared.UnionString].
type WebhookDeleteResponseUnion interface {
	ImplementsStreamWebhookDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WebhookDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [stream.WebhookGetResponseUnknown] or [shared.UnionString].
type WebhookGetResponseUnion interface {
	ImplementsStreamWebhookGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WebhookGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WebhookUpdateParams struct {
	// The URL where webhooks will be sent.
	NotificationURL param.Field[string] `json:"notificationUrl,required" format:"uri"`
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success WebhookUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  WebhookUpdateResponseUnion           `json:"result"`
	JSON    webhookUpdateResponseEnvelopeJSON    `json:"-"`
}

// webhookUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WebhookUpdateResponseEnvelope]
type webhookUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WebhookUpdateResponseEnvelopeSuccess bool

const (
	WebhookUpdateResponseEnvelopeSuccessTrue WebhookUpdateResponseEnvelopeSuccess = true
)

func (r WebhookUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WebhookUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WebhookDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success WebhookDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  WebhookDeleteResponseUnion           `json:"result"`
	JSON    webhookDeleteResponseEnvelopeJSON    `json:"-"`
}

// webhookDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WebhookDeleteResponseEnvelope]
type webhookDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WebhookDeleteResponseEnvelopeSuccess bool

const (
	WebhookDeleteResponseEnvelopeSuccessTrue WebhookDeleteResponseEnvelopeSuccess = true
)

func (r WebhookDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WebhookDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WebhookGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success WebhookGetResponseEnvelopeSuccess `json:"success,required"`
	Result  WebhookGetResponseUnion           `json:"result"`
	JSON    webhookGetResponseEnvelopeJSON    `json:"-"`
}

// webhookGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WebhookGetResponseEnvelope]
type webhookGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WebhookGetResponseEnvelopeSuccess bool

const (
	WebhookGetResponseEnvelopeSuccessTrue WebhookGetResponseEnvelopeSuccess = true
)

func (r WebhookGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WebhookGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
