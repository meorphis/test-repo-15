// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"net/http"
	"os"

	"github.com/stainless-sdks/meorphis-test-40-go/accounts"
	"github.com/stainless-sdks/meorphis-test-40-go/acm"
	"github.com/stainless-sdks/meorphis-test-40-go/addressing"
	"github.com/stainless-sdks/meorphis-test-40-go/ai_gateway"
	"github.com/stainless-sdks/meorphis-test-40-go/alerting"
	"github.com/stainless-sdks/meorphis-test-40-go/api_gateway"
	"github.com/stainless-sdks/meorphis-test-40-go/argo"
	"github.com/stainless-sdks/meorphis-test-40-go/audit_logs"
	"github.com/stainless-sdks/meorphis-test-40-go/billing"
	"github.com/stainless-sdks/meorphis-test-40-go/bot_management"
	"github.com/stainless-sdks/meorphis-test-40-go/brand_protection"
	"github.com/stainless-sdks/meorphis-test-40-go/cache"
	"github.com/stainless-sdks/meorphis-test-40-go/calls"
	"github.com/stainless-sdks/meorphis-test-40-go/certificate_authorities"
	"github.com/stainless-sdks/meorphis-test-40-go/challenges"
	"github.com/stainless-sdks/meorphis-test-40-go/client_certificates"
	"github.com/stainless-sdks/meorphis-test-40-go/cloud_connector"
	"github.com/stainless-sdks/meorphis-test-40-go/cloudforce_one"
	"github.com/stainless-sdks/meorphis-test-40-go/custom_certificates"
	"github.com/stainless-sdks/meorphis-test-40-go/custom_hostnames"
	"github.com/stainless-sdks/meorphis-test-40-go/custom_nameservers"
	"github.com/stainless-sdks/meorphis-test-40-go/d1"
	"github.com/stainless-sdks/meorphis-test-40-go/dcv_delegation"
	"github.com/stainless-sdks/meorphis-test-40-go/diagnostics"
	"github.com/stainless-sdks/meorphis-test-40-go/dns"
	"github.com/stainless-sdks/meorphis-test-40-go/dnssec"
	"github.com/stainless-sdks/meorphis-test-40-go/durable_objects"
	"github.com/stainless-sdks/meorphis-test-40-go/email_routing"
	"github.com/stainless-sdks/meorphis-test-40-go/event_notifications"
	"github.com/stainless-sdks/meorphis-test-40-go/filters"
	"github.com/stainless-sdks/meorphis-test-40-go/firewall"
	"github.com/stainless-sdks/meorphis-test-40-go/healthchecks"
	"github.com/stainless-sdks/meorphis-test-40-go/hostnames"
	"github.com/stainless-sdks/meorphis-test-40-go/hyperdrive"
	"github.com/stainless-sdks/meorphis-test-40-go/iam"
	"github.com/stainless-sdks/meorphis-test-40-go/images"
	"github.com/stainless-sdks/meorphis-test-40-go/intel"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/ips"
	"github.com/stainless-sdks/meorphis-test-40-go/keyless_certificates"
	"github.com/stainless-sdks/meorphis-test-40-go/kv"
	"github.com/stainless-sdks/meorphis-test-40-go/load_balancers"
	"github.com/stainless-sdks/meorphis-test-40-go/logpush"
	"github.com/stainless-sdks/meorphis-test-40-go/logs"
	"github.com/stainless-sdks/meorphis-test-40-go/magic_network_monitoring"
	"github.com/stainless-sdks/meorphis-test-40-go/magic_transit"
	"github.com/stainless-sdks/meorphis-test-40-go/managed_headers"
	"github.com/stainless-sdks/meorphis-test-40-go/memberships"
	"github.com/stainless-sdks/meorphis-test-40-go/mtls_certificates"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/origin_ca_certificates"
	"github.com/stainless-sdks/meorphis-test-40-go/origin_post_quantum_encryption"
	"github.com/stainless-sdks/meorphis-test-40-go/origin_tls_client_auth"
	"github.com/stainless-sdks/meorphis-test-40-go/page_shield"
	"github.com/stainless-sdks/meorphis-test-40-go/pagerules"
	"github.com/stainless-sdks/meorphis-test-40-go/pages"
	"github.com/stainless-sdks/meorphis-test-40-go/pcaps"
	"github.com/stainless-sdks/meorphis-test-40-go/plans"
	"github.com/stainless-sdks/meorphis-test-40-go/queues"
	"github.com/stainless-sdks/meorphis-test-40-go/r2"
	"github.com/stainless-sdks/meorphis-test-40-go/radar"
	"github.com/stainless-sdks/meorphis-test-40-go/rate_limits"
	"github.com/stainless-sdks/meorphis-test-40-go/rate_plans"
	"github.com/stainless-sdks/meorphis-test-40-go/registrar"
	"github.com/stainless-sdks/meorphis-test-40-go/request_tracers"
	"github.com/stainless-sdks/meorphis-test-40-go/rules"
	"github.com/stainless-sdks/meorphis-test-40-go/rulesets"
	"github.com/stainless-sdks/meorphis-test-40-go/rum"
	"github.com/stainless-sdks/meorphis-test-40-go/secondary_dns"
	"github.com/stainless-sdks/meorphis-test-40-go/snippets"
	"github.com/stainless-sdks/meorphis-test-40-go/spectrum"
	"github.com/stainless-sdks/meorphis-test-40-go/speed"
	"github.com/stainless-sdks/meorphis-test-40-go/ssl"
	"github.com/stainless-sdks/meorphis-test-40-go/storage"
	"github.com/stainless-sdks/meorphis-test-40-go/stream"
	"github.com/stainless-sdks/meorphis-test-40-go/subscriptions"
	"github.com/stainless-sdks/meorphis-test-40-go/url_normalization"
	"github.com/stainless-sdks/meorphis-test-40-go/url_scanner"
	"github.com/stainless-sdks/meorphis-test-40-go/user"
	"github.com/stainless-sdks/meorphis-test-40-go/vectorize"
	"github.com/stainless-sdks/meorphis-test-40-go/waiting_rooms"
	"github.com/stainless-sdks/meorphis-test-40-go/warp_connector"
	"github.com/stainless-sdks/meorphis-test-40-go/web3"
	"github.com/stainless-sdks/meorphis-test-40-go/workers"
	"github.com/stainless-sdks/meorphis-test-40-go/workers_for_platforms"
	"github.com/stainless-sdks/meorphis-test-40-go/zero_trust"
	"github.com/stainless-sdks/meorphis-test-40-go/zones"
)

// Client creates a struct with services and top level methods that help with
// interacting with the testcloudflare API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                     []option.RequestOption
	Accounts                    *accounts.AccountService
	OriginCACertificates        *origin_ca_certificates.OriginCACertificateService
	IPs                         *ips.IPService
	Memberships                 *memberships.MembershipService
	User                        *user.UserService
	Zones                       *zones.ZoneService
	LoadBalancers               *load_balancers.LoadBalancerService
	Cache                       *cache.CacheService
	SSL                         *ssl.SSLService
	Subscriptions               *subscriptions.SubscriptionService
	ACM                         *acm.ACMService
	Argo                        *argo.ArgoService
	Plans                       *plans.PlanService
	RatePlans                   *rate_plans.RatePlanService
	CertificateAuthorities      *certificate_authorities.CertificateAuthorityService
	ClientCertificates          *client_certificates.ClientCertificateService
	CustomCertificates          *custom_certificates.CustomCertificateService
	CustomHostnames             *custom_hostnames.CustomHostnameService
	CustomNameservers           *custom_nameservers.CustomNameserverService
	DNS                         *dns.DNSService
	DNSSEC                      *dnssec.DNSSECService
	EmailRouting                *email_routing.EmailRoutingService
	Filters                     *filters.FilterService
	Firewall                    *firewall.FirewallService
	Healthchecks                *healthchecks.HealthcheckService
	KeylessCertificates         *keyless_certificates.KeylessCertificateService
	Logpush                     *logpush.LogpushService
	Logs                        *logs.LogService
	OriginTLSClientAuth         *origin_tls_client_auth.OriginTLSClientAuthService
	Pagerules                   *pagerules.PageruleService
	RateLimits                  *rate_limits.RateLimitService
	SecondaryDNS                *secondary_dns.SecondaryDNSService
	WaitingRooms                *waiting_rooms.WaitingRoomService
	Web3                        *web3.Web3Service
	Workers                     *workers.WorkerService
	KV                          *kv.KVService
	DurableObjects              *durable_objects.DurableObjectService
	Queues                      *queues.QueueService
	APIGateway                  *api_gateway.APIGatewayService
	ManagedHeaders              *managed_headers.ManagedHeaderService
	PageShield                  *page_shield.PageShieldService
	Rulesets                    *rulesets.RulesetService
	URLNormalization            *url_normalization.URLNormalizationService
	Spectrum                    *spectrum.SpectrumService
	Addressing                  *addressing.AddressingService
	AuditLogs                   *audit_logs.AuditLogService
	Billing                     *billing.BillingService
	BrandProtection             *brand_protection.BrandProtectionService
	Diagnostics                 *diagnostics.DiagnosticService
	Images                      *images.ImageService
	Intel                       *intel.IntelService
	MagicTransit                *magic_transit.MagicTransitService
	MagicNetworkMonitoring      *magic_network_monitoring.MagicNetworkMonitoringService
	MTLSCertificates            *mtls_certificates.MTLSCertificateService
	Pages                       *pages.PageService
	PCAPs                       *pcaps.PCAPService
	Registrar                   *registrar.RegistrarService
	RequestTracers              *request_tracers.RequestTracerService
	Rules                       *rules.RuleService
	Storage                     *storage.StorageService
	Stream                      *stream.StreamService
	Alerting                    *alerting.AlertingService
	D1                          *d1.D1Service
	R2                          *r2.R2Service
	WARPConnector               *warp_connector.WARPConnectorService
	WorkersForPlatforms         *workers_for_platforms.WorkersForPlatformService
	ZeroTrust                   *zero_trust.ZeroTrustService
	Challenges                  *challenges.ChallengeService
	Hyperdrive                  *hyperdrive.HyperdriveService
	RUM                         *rum.RUMService
	Vectorize                   *vectorize.VectorizeService
	URLScanner                  *url_scanner.URLScannerService
	Radar                       *radar.RadarService
	BotManagement               *bot_management.BotManagementService
	OriginPostQuantumEncryption *origin_post_quantum_encryption.OriginPostQuantumEncryptionService
	Speed                       *speed.SpeedService
	DCVDelegation               *dcv_delegation.DCVDelegationService
	Hostnames                   *hostnames.HostnameService
	Snippets                    *snippets.SnippetService
	Calls                       *calls.CallService
	CloudforceOne               *cloudforce_one.CloudforceOneService
	EventNotifications          *event_notifications.EventNotificationService
	AIGateway                   *ai_gateway.AIGatewayService
	IAM                         *iam.IAMService
	CloudConnector              *cloud_connector.CloudConnectorService
}

// NewClient generates a new client with the default option read from the
// environment (CLOUDFLARE_API_TOKEN, CLOUDFLARE_API_KEY, CLOUDFLARE_EMAIL,
// CLOUDFLARE_API_USER_SERVICE_KEY). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_TOKEN"); ok {
		defaults = append(defaults, option.WithAPIToken(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_EMAIL"); ok {
		defaults = append(defaults, option.WithAPIEmail(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_USER_SERVICE_KEY"); ok {
		defaults = append(defaults, option.WithUserServiceKey(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Accounts = accounts.NewAccountService(opts...)
	r.OriginCACertificates = origin_ca_certificates.NewOriginCACertificateService(opts...)
	r.IPs = ips.NewIPService(opts...)
	r.Memberships = memberships.NewMembershipService(opts...)
	r.User = user.NewUserService(opts...)
	r.Zones = zones.NewZoneService(opts...)
	r.LoadBalancers = load_balancers.NewLoadBalancerService(opts...)
	r.Cache = cache.NewCacheService(opts...)
	r.SSL = ssl.NewSSLService(opts...)
	r.Subscriptions = subscriptions.NewSubscriptionService(opts...)
	r.ACM = acm.NewACMService(opts...)
	r.Argo = argo.NewArgoService(opts...)
	r.Plans = plans.NewPlanService(opts...)
	r.RatePlans = rate_plans.NewRatePlanService(opts...)
	r.CertificateAuthorities = certificate_authorities.NewCertificateAuthorityService(opts...)
	r.ClientCertificates = client_certificates.NewClientCertificateService(opts...)
	r.CustomCertificates = custom_certificates.NewCustomCertificateService(opts...)
	r.CustomHostnames = custom_hostnames.NewCustomHostnameService(opts...)
	r.CustomNameservers = custom_nameservers.NewCustomNameserverService(opts...)
	r.DNS = dns.NewDNSService(opts...)
	r.DNSSEC = dnssec.NewDNSSECService(opts...)
	r.EmailRouting = email_routing.NewEmailRoutingService(opts...)
	r.Filters = filters.NewFilterService(opts...)
	r.Firewall = firewall.NewFirewallService(opts...)
	r.Healthchecks = healthchecks.NewHealthcheckService(opts...)
	r.KeylessCertificates = keyless_certificates.NewKeylessCertificateService(opts...)
	r.Logpush = logpush.NewLogpushService(opts...)
	r.Logs = logs.NewLogService(opts...)
	r.OriginTLSClientAuth = origin_tls_client_auth.NewOriginTLSClientAuthService(opts...)
	r.Pagerules = pagerules.NewPageruleService(opts...)
	r.RateLimits = rate_limits.NewRateLimitService(opts...)
	r.SecondaryDNS = secondary_dns.NewSecondaryDNSService(opts...)
	r.WaitingRooms = waiting_rooms.NewWaitingRoomService(opts...)
	r.Web3 = web3.NewWeb3Service(opts...)
	r.Workers = workers.NewWorkerService(opts...)
	r.KV = kv.NewKVService(opts...)
	r.DurableObjects = durable_objects.NewDurableObjectService(opts...)
	r.Queues = queues.NewQueueService(opts...)
	r.APIGateway = api_gateway.NewAPIGatewayService(opts...)
	r.ManagedHeaders = managed_headers.NewManagedHeaderService(opts...)
	r.PageShield = page_shield.NewPageShieldService(opts...)
	r.Rulesets = rulesets.NewRulesetService(opts...)
	r.URLNormalization = url_normalization.NewURLNormalizationService(opts...)
	r.Spectrum = spectrum.NewSpectrumService(opts...)
	r.Addressing = addressing.NewAddressingService(opts...)
	r.AuditLogs = audit_logs.NewAuditLogService(opts...)
	r.Billing = billing.NewBillingService(opts...)
	r.BrandProtection = brand_protection.NewBrandProtectionService(opts...)
	r.Diagnostics = diagnostics.NewDiagnosticService(opts...)
	r.Images = images.NewImageService(opts...)
	r.Intel = intel.NewIntelService(opts...)
	r.MagicTransit = magic_transit.NewMagicTransitService(opts...)
	r.MagicNetworkMonitoring = magic_network_monitoring.NewMagicNetworkMonitoringService(opts...)
	r.MTLSCertificates = mtls_certificates.NewMTLSCertificateService(opts...)
	r.Pages = pages.NewPageService(opts...)
	r.PCAPs = pcaps.NewPCAPService(opts...)
	r.Registrar = registrar.NewRegistrarService(opts...)
	r.RequestTracers = request_tracers.NewRequestTracerService(opts...)
	r.Rules = rules.NewRuleService(opts...)
	r.Storage = storage.NewStorageService(opts...)
	r.Stream = stream.NewStreamService(opts...)
	r.Alerting = alerting.NewAlertingService(opts...)
	r.D1 = d1.NewD1Service(opts...)
	r.R2 = r2.NewR2Service(opts...)
	r.WARPConnector = warp_connector.NewWARPConnectorService(opts...)
	r.WorkersForPlatforms = workers_for_platforms.NewWorkersForPlatformService(opts...)
	r.ZeroTrust = zero_trust.NewZeroTrustService(opts...)
	r.Challenges = challenges.NewChallengeService(opts...)
	r.Hyperdrive = hyperdrive.NewHyperdriveService(opts...)
	r.RUM = rum.NewRUMService(opts...)
	r.Vectorize = vectorize.NewVectorizeService(opts...)
	r.URLScanner = url_scanner.NewURLScannerService(opts...)
	r.Radar = radar.NewRadarService(opts...)
	r.BotManagement = bot_management.NewBotManagementService(opts...)
	r.OriginPostQuantumEncryption = origin_post_quantum_encryption.NewOriginPostQuantumEncryptionService(opts...)
	r.Speed = speed.NewSpeedService(opts...)
	r.DCVDelegation = dcv_delegation.NewDCVDelegationService(opts...)
	r.Hostnames = hostnames.NewHostnameService(opts...)
	r.Snippets = snippets.NewSnippetService(opts...)
	r.Calls = calls.NewCallService(opts...)
	r.CloudforceOne = cloudforce_one.NewCloudforceOneService(opts...)
	r.EventNotifications = event_notifications.NewEventNotificationService(opts...)
	r.AIGateway = ai_gateway.NewAIGatewayService(opts...)
	r.IAM = iam.NewIAMService(opts...)
	r.CloudConnector = cloud_connector.NewCloudConnectorService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = append(r.Options, opts...)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
