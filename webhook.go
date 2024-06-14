// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/tidwall/gjson"
)

// WebhookService contains methods and other services that help with interacting
// with the meorphis-test-40 API.
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

// Create a new webhook to receive notifications from Bolt about various events,
// such as transaction status.
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve information for an existing webhook.
func (r *WebhookService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve information about all existing webhooks.
func (r *WebhookService) List(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing webhook. You will no longer receive notifications from Bolt
// about its events.
func (r *WebhookService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type WebhookNewResponse struct {
	Event WebhookNewResponseEvent `json:"event,required"`
	// The webhook's URL
	URL string `json:"url,required" format:"uri"`
	// The webhook's ID
	ID string `json:"id"`
	// The time at which the webhook was created
	CreatedAt time.Time              `json:"created_at" format:"date-time"`
	JSON      webhookNewResponseJSON `json:"-"`
}

// webhookNewResponseJSON contains the JSON metadata for the struct
// [WebhookNewResponse]
type webhookNewResponseJSON struct {
	Event       apijson.Field
	URL         apijson.Field
	ID          apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookNewResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookNewResponseEvent struct {
	Tag        WebhookNewResponseEventTag        `json:".tag,required"`
	EventGroup WebhookNewResponseEventEventGroup `json:"event_group"`
	EventList  interface{}                       `json:"event_list,required"`
	JSON       webhookNewResponseEventJSON       `json:"-"`
	union      WebhookNewResponseEventUnion
}

// webhookNewResponseEventJSON contains the JSON metadata for the struct
// [WebhookNewResponseEvent]
type webhookNewResponseEventJSON struct {
	Tag         apijson.Field
	EventGroup  apijson.Field
	EventList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r webhookNewResponseEventJSON) RawJSON() string {
	return r.raw
}

func (r *WebhookNewResponseEvent) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WebhookNewResponseEvent) AsUnion() WebhookNewResponseEventUnion {
	return r.union
}

// Union satisfied by [WebhookNewResponseEventEventGroup] or
// [WebhookNewResponseEventEventList].
type WebhookNewResponseEventUnion interface {
	implementsWebhookNewResponseEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WebhookNewResponseEventUnion)(nil)).Elem(),
		".tag",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(WebhookNewResponseEventEventGroup{}),
			DiscriminatorValue: "group",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(WebhookNewResponseEventEventList{}),
			DiscriminatorValue: "list",
		},
	)
}

type WebhookNewResponseEventEventGroup struct {
	Tag        WebhookNewResponseEventEventGroupTag        `json:".tag,required"`
	EventGroup WebhookNewResponseEventEventGroupEventGroup `json:"event_group,required"`
	JSON       webhookNewResponseEventEventGroupJSON       `json:"-"`
}

// webhookNewResponseEventEventGroupJSON contains the JSON metadata for the struct
// [WebhookNewResponseEventEventGroup]
type webhookNewResponseEventEventGroupJSON struct {
	Tag         apijson.Field
	EventGroup  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookNewResponseEventEventGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookNewResponseEventEventGroupJSON) RawJSON() string {
	return r.raw
}

func (r WebhookNewResponseEventEventGroup) implementsWebhookNewResponseEvent() {}

type WebhookNewResponseEventEventGroupTag string

const (
	WebhookNewResponseEventEventGroupTagGroup WebhookNewResponseEventEventGroupTag = "group"
)

func (r WebhookNewResponseEventEventGroupTag) IsKnown() bool {
	switch r {
	case WebhookNewResponseEventEventGroupTagGroup:
		return true
	}
	return false
}

type WebhookNewResponseEventEventGroupEventGroup string

const (
	WebhookNewResponseEventEventGroupEventGroupAll WebhookNewResponseEventEventGroupEventGroup = "all"
)

func (r WebhookNewResponseEventEventGroupEventGroup) IsKnown() bool {
	switch r {
	case WebhookNewResponseEventEventGroupEventGroupAll:
		return true
	}
	return false
}

type WebhookNewResponseEventEventList struct {
	Tag       WebhookNewResponseEventEventListTag         `json:".tag,required"`
	EventList []WebhookNewResponseEventEventListEventList `json:"event_list,required"`
	JSON      webhookNewResponseEventEventListJSON        `json:"-"`
}

// webhookNewResponseEventEventListJSON contains the JSON metadata for the struct
// [WebhookNewResponseEventEventList]
type webhookNewResponseEventEventListJSON struct {
	Tag         apijson.Field
	EventList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookNewResponseEventEventList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookNewResponseEventEventListJSON) RawJSON() string {
	return r.raw
}

func (r WebhookNewResponseEventEventList) implementsWebhookNewResponseEvent() {}

type WebhookNewResponseEventEventListTag string

const (
	WebhookNewResponseEventEventListTagList WebhookNewResponseEventEventListTag = "list"
)

func (r WebhookNewResponseEventEventListTag) IsKnown() bool {
	switch r {
	case WebhookNewResponseEventEventListTagList:
		return true
	}
	return false
}

type WebhookNewResponseEventEventListEventList string

const (
	WebhookNewResponseEventEventListEventListPayment                WebhookNewResponseEventEventListEventList = "payment"
	WebhookNewResponseEventEventListEventListCredit                 WebhookNewResponseEventEventListEventList = "credit"
	WebhookNewResponseEventEventListEventListCapture                WebhookNewResponseEventEventListEventList = "capture"
	WebhookNewResponseEventEventListEventListVoid                   WebhookNewResponseEventEventListEventList = "void"
	WebhookNewResponseEventEventListEventListAuth                   WebhookNewResponseEventEventListEventList = "auth"
	WebhookNewResponseEventEventListEventListPending                WebhookNewResponseEventEventListEventList = "pending"
	WebhookNewResponseEventEventListEventListRejectedIrreversible   WebhookNewResponseEventEventListEventList = "rejected_irreversible"
	WebhookNewResponseEventEventListEventListRejectedReversible     WebhookNewResponseEventEventListEventList = "rejected_reversible"
	WebhookNewResponseEventEventListEventListFailedPayment          WebhookNewResponseEventEventListEventList = "failed_payment"
	WebhookNewResponseEventEventListEventListNewsletterSubscription WebhookNewResponseEventEventListEventList = "newsletter_subscription"
	WebhookNewResponseEventEventListEventListRiskInsights           WebhookNewResponseEventEventListEventList = "risk_insights"
	WebhookNewResponseEventEventListEventListCreditCardDeleted      WebhookNewResponseEventEventListEventList = "credit_card_deleted"
)

func (r WebhookNewResponseEventEventListEventList) IsKnown() bool {
	switch r {
	case WebhookNewResponseEventEventListEventListPayment, WebhookNewResponseEventEventListEventListCredit, WebhookNewResponseEventEventListEventListCapture, WebhookNewResponseEventEventListEventListVoid, WebhookNewResponseEventEventListEventListAuth, WebhookNewResponseEventEventListEventListPending, WebhookNewResponseEventEventListEventListRejectedIrreversible, WebhookNewResponseEventEventListEventListRejectedReversible, WebhookNewResponseEventEventListEventListFailedPayment, WebhookNewResponseEventEventListEventListNewsletterSubscription, WebhookNewResponseEventEventListEventListRiskInsights, WebhookNewResponseEventEventListEventListCreditCardDeleted:
		return true
	}
	return false
}

type WebhookNewResponseEventTag string

const (
	WebhookNewResponseEventTagGroup WebhookNewResponseEventTag = "group"
	WebhookNewResponseEventTagList  WebhookNewResponseEventTag = "list"
)

func (r WebhookNewResponseEventTag) IsKnown() bool {
	switch r {
	case WebhookNewResponseEventTagGroup, WebhookNewResponseEventTagList:
		return true
	}
	return false
}

type WebhookGetResponse struct {
	Event WebhookGetResponseEvent `json:"event,required"`
	// The webhook's URL
	URL string `json:"url,required" format:"uri"`
	// The webhook's ID
	ID string `json:"id"`
	// The time at which the webhook was created
	CreatedAt time.Time              `json:"created_at" format:"date-time"`
	JSON      webhookGetResponseJSON `json:"-"`
}

// webhookGetResponseJSON contains the JSON metadata for the struct
// [WebhookGetResponse]
type webhookGetResponseJSON struct {
	Event       apijson.Field
	URL         apijson.Field
	ID          apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookGetResponseEvent struct {
	Tag        WebhookGetResponseEventTag        `json:".tag,required"`
	EventGroup WebhookGetResponseEventEventGroup `json:"event_group"`
	EventList  interface{}                       `json:"event_list,required"`
	JSON       webhookGetResponseEventJSON       `json:"-"`
	union      WebhookGetResponseEventUnion
}

// webhookGetResponseEventJSON contains the JSON metadata for the struct
// [WebhookGetResponseEvent]
type webhookGetResponseEventJSON struct {
	Tag         apijson.Field
	EventGroup  apijson.Field
	EventList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r webhookGetResponseEventJSON) RawJSON() string {
	return r.raw
}

func (r *WebhookGetResponseEvent) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WebhookGetResponseEvent) AsUnion() WebhookGetResponseEventUnion {
	return r.union
}

// Union satisfied by [WebhookGetResponseEventEventGroup] or
// [WebhookGetResponseEventEventList].
type WebhookGetResponseEventUnion interface {
	implementsWebhookGetResponseEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WebhookGetResponseEventUnion)(nil)).Elem(),
		".tag",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(WebhookGetResponseEventEventGroup{}),
			DiscriminatorValue: "group",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(WebhookGetResponseEventEventList{}),
			DiscriminatorValue: "list",
		},
	)
}

type WebhookGetResponseEventEventGroup struct {
	Tag        WebhookGetResponseEventEventGroupTag        `json:".tag,required"`
	EventGroup WebhookGetResponseEventEventGroupEventGroup `json:"event_group,required"`
	JSON       webhookGetResponseEventEventGroupJSON       `json:"-"`
}

// webhookGetResponseEventEventGroupJSON contains the JSON metadata for the struct
// [WebhookGetResponseEventEventGroup]
type webhookGetResponseEventEventGroupJSON struct {
	Tag         apijson.Field
	EventGroup  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetResponseEventEventGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetResponseEventEventGroupJSON) RawJSON() string {
	return r.raw
}

func (r WebhookGetResponseEventEventGroup) implementsWebhookGetResponseEvent() {}

type WebhookGetResponseEventEventGroupTag string

const (
	WebhookGetResponseEventEventGroupTagGroup WebhookGetResponseEventEventGroupTag = "group"
)

func (r WebhookGetResponseEventEventGroupTag) IsKnown() bool {
	switch r {
	case WebhookGetResponseEventEventGroupTagGroup:
		return true
	}
	return false
}

type WebhookGetResponseEventEventGroupEventGroup string

const (
	WebhookGetResponseEventEventGroupEventGroupAll WebhookGetResponseEventEventGroupEventGroup = "all"
)

func (r WebhookGetResponseEventEventGroupEventGroup) IsKnown() bool {
	switch r {
	case WebhookGetResponseEventEventGroupEventGroupAll:
		return true
	}
	return false
}

type WebhookGetResponseEventEventList struct {
	Tag       WebhookGetResponseEventEventListTag         `json:".tag,required"`
	EventList []WebhookGetResponseEventEventListEventList `json:"event_list,required"`
	JSON      webhookGetResponseEventEventListJSON        `json:"-"`
}

// webhookGetResponseEventEventListJSON contains the JSON metadata for the struct
// [WebhookGetResponseEventEventList]
type webhookGetResponseEventEventListJSON struct {
	Tag         apijson.Field
	EventList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetResponseEventEventList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetResponseEventEventListJSON) RawJSON() string {
	return r.raw
}

func (r WebhookGetResponseEventEventList) implementsWebhookGetResponseEvent() {}

type WebhookGetResponseEventEventListTag string

const (
	WebhookGetResponseEventEventListTagList WebhookGetResponseEventEventListTag = "list"
)

func (r WebhookGetResponseEventEventListTag) IsKnown() bool {
	switch r {
	case WebhookGetResponseEventEventListTagList:
		return true
	}
	return false
}

type WebhookGetResponseEventEventListEventList string

const (
	WebhookGetResponseEventEventListEventListPayment                WebhookGetResponseEventEventListEventList = "payment"
	WebhookGetResponseEventEventListEventListCredit                 WebhookGetResponseEventEventListEventList = "credit"
	WebhookGetResponseEventEventListEventListCapture                WebhookGetResponseEventEventListEventList = "capture"
	WebhookGetResponseEventEventListEventListVoid                   WebhookGetResponseEventEventListEventList = "void"
	WebhookGetResponseEventEventListEventListAuth                   WebhookGetResponseEventEventListEventList = "auth"
	WebhookGetResponseEventEventListEventListPending                WebhookGetResponseEventEventListEventList = "pending"
	WebhookGetResponseEventEventListEventListRejectedIrreversible   WebhookGetResponseEventEventListEventList = "rejected_irreversible"
	WebhookGetResponseEventEventListEventListRejectedReversible     WebhookGetResponseEventEventListEventList = "rejected_reversible"
	WebhookGetResponseEventEventListEventListFailedPayment          WebhookGetResponseEventEventListEventList = "failed_payment"
	WebhookGetResponseEventEventListEventListNewsletterSubscription WebhookGetResponseEventEventListEventList = "newsletter_subscription"
	WebhookGetResponseEventEventListEventListRiskInsights           WebhookGetResponseEventEventListEventList = "risk_insights"
	WebhookGetResponseEventEventListEventListCreditCardDeleted      WebhookGetResponseEventEventListEventList = "credit_card_deleted"
)

func (r WebhookGetResponseEventEventListEventList) IsKnown() bool {
	switch r {
	case WebhookGetResponseEventEventListEventListPayment, WebhookGetResponseEventEventListEventListCredit, WebhookGetResponseEventEventListEventListCapture, WebhookGetResponseEventEventListEventListVoid, WebhookGetResponseEventEventListEventListAuth, WebhookGetResponseEventEventListEventListPending, WebhookGetResponseEventEventListEventListRejectedIrreversible, WebhookGetResponseEventEventListEventListRejectedReversible, WebhookGetResponseEventEventListEventListFailedPayment, WebhookGetResponseEventEventListEventListNewsletterSubscription, WebhookGetResponseEventEventListEventListRiskInsights, WebhookGetResponseEventEventListEventListCreditCardDeleted:
		return true
	}
	return false
}

type WebhookGetResponseEventTag string

const (
	WebhookGetResponseEventTagGroup WebhookGetResponseEventTag = "group"
	WebhookGetResponseEventTagList  WebhookGetResponseEventTag = "list"
)

func (r WebhookGetResponseEventTag) IsKnown() bool {
	switch r {
	case WebhookGetResponseEventTagGroup, WebhookGetResponseEventTagList:
		return true
	}
	return false
}

type WebhookListResponse struct {
	Webhooks []WebhookListResponseWebhook `json:"webhooks"`
	JSON     webhookListResponseJSON      `json:"-"`
}

// webhookListResponseJSON contains the JSON metadata for the struct
// [WebhookListResponse]
type webhookListResponseJSON struct {
	Webhooks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookListResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookListResponseWebhook struct {
	Event WebhookListResponseWebhooksEvent `json:"event,required"`
	// The webhook's URL
	URL string `json:"url,required" format:"uri"`
	// The webhook's ID
	ID string `json:"id"`
	// The time at which the webhook was created
	CreatedAt time.Time                      `json:"created_at" format:"date-time"`
	JSON      webhookListResponseWebhookJSON `json:"-"`
}

// webhookListResponseWebhookJSON contains the JSON metadata for the struct
// [WebhookListResponseWebhook]
type webhookListResponseWebhookJSON struct {
	Event       apijson.Field
	URL         apijson.Field
	ID          apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookListResponseWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookListResponseWebhookJSON) RawJSON() string {
	return r.raw
}

type WebhookListResponseWebhooksEvent struct {
	Tag        WebhookListResponseWebhooksEventTag        `json:".tag,required"`
	EventGroup WebhookListResponseWebhooksEventEventGroup `json:"event_group"`
	EventList  interface{}                                `json:"event_list,required"`
	JSON       webhookListResponseWebhooksEventJSON       `json:"-"`
	union      WebhookListResponseWebhooksEventUnion
}

// webhookListResponseWebhooksEventJSON contains the JSON metadata for the struct
// [WebhookListResponseWebhooksEvent]
type webhookListResponseWebhooksEventJSON struct {
	Tag         apijson.Field
	EventGroup  apijson.Field
	EventList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r webhookListResponseWebhooksEventJSON) RawJSON() string {
	return r.raw
}

func (r *WebhookListResponseWebhooksEvent) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WebhookListResponseWebhooksEvent) AsUnion() WebhookListResponseWebhooksEventUnion {
	return r.union
}

// Union satisfied by [WebhookListResponseWebhooksEventEventGroup] or
// [WebhookListResponseWebhooksEventEventList].
type WebhookListResponseWebhooksEventUnion interface {
	implementsWebhookListResponseWebhooksEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WebhookListResponseWebhooksEventUnion)(nil)).Elem(),
		".tag",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(WebhookListResponseWebhooksEventEventGroup{}),
			DiscriminatorValue: "group",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(WebhookListResponseWebhooksEventEventList{}),
			DiscriminatorValue: "list",
		},
	)
}

type WebhookListResponseWebhooksEventEventGroup struct {
	Tag        WebhookListResponseWebhooksEventEventGroupTag        `json:".tag,required"`
	EventGroup WebhookListResponseWebhooksEventEventGroupEventGroup `json:"event_group,required"`
	JSON       webhookListResponseWebhooksEventEventGroupJSON       `json:"-"`
}

// webhookListResponseWebhooksEventEventGroupJSON contains the JSON metadata for
// the struct [WebhookListResponseWebhooksEventEventGroup]
type webhookListResponseWebhooksEventEventGroupJSON struct {
	Tag         apijson.Field
	EventGroup  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookListResponseWebhooksEventEventGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookListResponseWebhooksEventEventGroupJSON) RawJSON() string {
	return r.raw
}

func (r WebhookListResponseWebhooksEventEventGroup) implementsWebhookListResponseWebhooksEvent() {}

type WebhookListResponseWebhooksEventEventGroupTag string

const (
	WebhookListResponseWebhooksEventEventGroupTagGroup WebhookListResponseWebhooksEventEventGroupTag = "group"
)

func (r WebhookListResponseWebhooksEventEventGroupTag) IsKnown() bool {
	switch r {
	case WebhookListResponseWebhooksEventEventGroupTagGroup:
		return true
	}
	return false
}

type WebhookListResponseWebhooksEventEventGroupEventGroup string

const (
	WebhookListResponseWebhooksEventEventGroupEventGroupAll WebhookListResponseWebhooksEventEventGroupEventGroup = "all"
)

func (r WebhookListResponseWebhooksEventEventGroupEventGroup) IsKnown() bool {
	switch r {
	case WebhookListResponseWebhooksEventEventGroupEventGroupAll:
		return true
	}
	return false
}

type WebhookListResponseWebhooksEventEventList struct {
	Tag       WebhookListResponseWebhooksEventEventListTag         `json:".tag,required"`
	EventList []WebhookListResponseWebhooksEventEventListEventList `json:"event_list,required"`
	JSON      webhookListResponseWebhooksEventEventListJSON        `json:"-"`
}

// webhookListResponseWebhooksEventEventListJSON contains the JSON metadata for the
// struct [WebhookListResponseWebhooksEventEventList]
type webhookListResponseWebhooksEventEventListJSON struct {
	Tag         apijson.Field
	EventList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookListResponseWebhooksEventEventList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookListResponseWebhooksEventEventListJSON) RawJSON() string {
	return r.raw
}

func (r WebhookListResponseWebhooksEventEventList) implementsWebhookListResponseWebhooksEvent() {}

type WebhookListResponseWebhooksEventEventListTag string

const (
	WebhookListResponseWebhooksEventEventListTagList WebhookListResponseWebhooksEventEventListTag = "list"
)

func (r WebhookListResponseWebhooksEventEventListTag) IsKnown() bool {
	switch r {
	case WebhookListResponseWebhooksEventEventListTagList:
		return true
	}
	return false
}

type WebhookListResponseWebhooksEventEventListEventList string

const (
	WebhookListResponseWebhooksEventEventListEventListPayment                WebhookListResponseWebhooksEventEventListEventList = "payment"
	WebhookListResponseWebhooksEventEventListEventListCredit                 WebhookListResponseWebhooksEventEventListEventList = "credit"
	WebhookListResponseWebhooksEventEventListEventListCapture                WebhookListResponseWebhooksEventEventListEventList = "capture"
	WebhookListResponseWebhooksEventEventListEventListVoid                   WebhookListResponseWebhooksEventEventListEventList = "void"
	WebhookListResponseWebhooksEventEventListEventListAuth                   WebhookListResponseWebhooksEventEventListEventList = "auth"
	WebhookListResponseWebhooksEventEventListEventListPending                WebhookListResponseWebhooksEventEventListEventList = "pending"
	WebhookListResponseWebhooksEventEventListEventListRejectedIrreversible   WebhookListResponseWebhooksEventEventListEventList = "rejected_irreversible"
	WebhookListResponseWebhooksEventEventListEventListRejectedReversible     WebhookListResponseWebhooksEventEventListEventList = "rejected_reversible"
	WebhookListResponseWebhooksEventEventListEventListFailedPayment          WebhookListResponseWebhooksEventEventListEventList = "failed_payment"
	WebhookListResponseWebhooksEventEventListEventListNewsletterSubscription WebhookListResponseWebhooksEventEventListEventList = "newsletter_subscription"
	WebhookListResponseWebhooksEventEventListEventListRiskInsights           WebhookListResponseWebhooksEventEventListEventList = "risk_insights"
	WebhookListResponseWebhooksEventEventListEventListCreditCardDeleted      WebhookListResponseWebhooksEventEventListEventList = "credit_card_deleted"
)

func (r WebhookListResponseWebhooksEventEventListEventList) IsKnown() bool {
	switch r {
	case WebhookListResponseWebhooksEventEventListEventListPayment, WebhookListResponseWebhooksEventEventListEventListCredit, WebhookListResponseWebhooksEventEventListEventListCapture, WebhookListResponseWebhooksEventEventListEventListVoid, WebhookListResponseWebhooksEventEventListEventListAuth, WebhookListResponseWebhooksEventEventListEventListPending, WebhookListResponseWebhooksEventEventListEventListRejectedIrreversible, WebhookListResponseWebhooksEventEventListEventListRejectedReversible, WebhookListResponseWebhooksEventEventListEventListFailedPayment, WebhookListResponseWebhooksEventEventListEventListNewsletterSubscription, WebhookListResponseWebhooksEventEventListEventListRiskInsights, WebhookListResponseWebhooksEventEventListEventListCreditCardDeleted:
		return true
	}
	return false
}

type WebhookListResponseWebhooksEventTag string

const (
	WebhookListResponseWebhooksEventTagGroup WebhookListResponseWebhooksEventTag = "group"
	WebhookListResponseWebhooksEventTagList  WebhookListResponseWebhooksEventTag = "list"
)

func (r WebhookListResponseWebhooksEventTag) IsKnown() bool {
	switch r {
	case WebhookListResponseWebhooksEventTagGroup, WebhookListResponseWebhooksEventTagList:
		return true
	}
	return false
}

type WebhookNewParams struct {
	Event param.Field[WebhookNewParamsEventUnion] `json:"event,required"`
	// The webhook's URL
	URL param.Field[string] `json:"url,required" format:"uri"`
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookNewParamsEvent struct {
	Tag        param.Field[WebhookNewParamsEventTag]        `json:".tag,required"`
	EventGroup param.Field[WebhookNewParamsEventEventGroup] `json:"event_group"`
	EventList  param.Field[interface{}]                     `json:"event_list,required"`
}

func (r WebhookNewParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookNewParamsEvent) implementsWebhookNewParamsEventUnion() {}

// Satisfied by [WebhookNewParamsEventEventGroup],
// [WebhookNewParamsEventEventList], [WebhookNewParamsEvent].
type WebhookNewParamsEventUnion interface {
	implementsWebhookNewParamsEventUnion()
}

type WebhookNewParamsEventEventGroup struct {
	Tag        param.Field[WebhookNewParamsEventEventGroupTag]        `json:".tag,required"`
	EventGroup param.Field[WebhookNewParamsEventEventGroupEventGroup] `json:"event_group,required"`
}

func (r WebhookNewParamsEventEventGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookNewParamsEventEventGroup) implementsWebhookNewParamsEventUnion() {}

type WebhookNewParamsEventEventGroupTag string

const (
	WebhookNewParamsEventEventGroupTagGroup WebhookNewParamsEventEventGroupTag = "group"
)

func (r WebhookNewParamsEventEventGroupTag) IsKnown() bool {
	switch r {
	case WebhookNewParamsEventEventGroupTagGroup:
		return true
	}
	return false
}

type WebhookNewParamsEventEventGroupEventGroup string

const (
	WebhookNewParamsEventEventGroupEventGroupAll WebhookNewParamsEventEventGroupEventGroup = "all"
)

func (r WebhookNewParamsEventEventGroupEventGroup) IsKnown() bool {
	switch r {
	case WebhookNewParamsEventEventGroupEventGroupAll:
		return true
	}
	return false
}

type WebhookNewParamsEventEventList struct {
	Tag       param.Field[WebhookNewParamsEventEventListTag]         `json:".tag,required"`
	EventList param.Field[[]WebhookNewParamsEventEventListEventList] `json:"event_list,required"`
}

func (r WebhookNewParamsEventEventList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookNewParamsEventEventList) implementsWebhookNewParamsEventUnion() {}

type WebhookNewParamsEventEventListTag string

const (
	WebhookNewParamsEventEventListTagList WebhookNewParamsEventEventListTag = "list"
)

func (r WebhookNewParamsEventEventListTag) IsKnown() bool {
	switch r {
	case WebhookNewParamsEventEventListTagList:
		return true
	}
	return false
}

type WebhookNewParamsEventEventListEventList string

const (
	WebhookNewParamsEventEventListEventListPayment                WebhookNewParamsEventEventListEventList = "payment"
	WebhookNewParamsEventEventListEventListCredit                 WebhookNewParamsEventEventListEventList = "credit"
	WebhookNewParamsEventEventListEventListCapture                WebhookNewParamsEventEventListEventList = "capture"
	WebhookNewParamsEventEventListEventListVoid                   WebhookNewParamsEventEventListEventList = "void"
	WebhookNewParamsEventEventListEventListAuth                   WebhookNewParamsEventEventListEventList = "auth"
	WebhookNewParamsEventEventListEventListPending                WebhookNewParamsEventEventListEventList = "pending"
	WebhookNewParamsEventEventListEventListRejectedIrreversible   WebhookNewParamsEventEventListEventList = "rejected_irreversible"
	WebhookNewParamsEventEventListEventListRejectedReversible     WebhookNewParamsEventEventListEventList = "rejected_reversible"
	WebhookNewParamsEventEventListEventListFailedPayment          WebhookNewParamsEventEventListEventList = "failed_payment"
	WebhookNewParamsEventEventListEventListNewsletterSubscription WebhookNewParamsEventEventListEventList = "newsletter_subscription"
	WebhookNewParamsEventEventListEventListRiskInsights           WebhookNewParamsEventEventListEventList = "risk_insights"
	WebhookNewParamsEventEventListEventListCreditCardDeleted      WebhookNewParamsEventEventListEventList = "credit_card_deleted"
)

func (r WebhookNewParamsEventEventListEventList) IsKnown() bool {
	switch r {
	case WebhookNewParamsEventEventListEventListPayment, WebhookNewParamsEventEventListEventListCredit, WebhookNewParamsEventEventListEventListCapture, WebhookNewParamsEventEventListEventListVoid, WebhookNewParamsEventEventListEventListAuth, WebhookNewParamsEventEventListEventListPending, WebhookNewParamsEventEventListEventListRejectedIrreversible, WebhookNewParamsEventEventListEventListRejectedReversible, WebhookNewParamsEventEventListEventListFailedPayment, WebhookNewParamsEventEventListEventListNewsletterSubscription, WebhookNewParamsEventEventListEventListRiskInsights, WebhookNewParamsEventEventListEventListCreditCardDeleted:
		return true
	}
	return false
}

type WebhookNewParamsEventTag string

const (
	WebhookNewParamsEventTagGroup WebhookNewParamsEventTag = "group"
	WebhookNewParamsEventTagList  WebhookNewParamsEventTag = "list"
)

func (r WebhookNewParamsEventTag) IsKnown() bool {
	switch r {
	case WebhookNewParamsEventTagGroup, WebhookNewParamsEventTagList:
		return true
	}
	return false
}

type WebhookListParams struct {
	XPublishableKey param.Field[string] `header:"X-Publishable-Key,required"`
}
