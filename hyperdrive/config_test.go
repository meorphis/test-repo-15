// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/hyperdrive"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestConfigNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Hyperdrive.Configs.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		hyperdrive.ConfigNewParams{
			Hyperdrive: hyperdrive.HyperdriveParam{
				Name: cloudflare.F("example-hyperdrive"),
				Origin: cloudflare.F(hyperdrive.ConfigurationParam{
					Database:       cloudflare.F("postgres"),
					Host:           cloudflare.F("database.example.com"),
					Scheme:         cloudflare.F(hyperdrive.ConfigurationSchemePostgres),
					User:           cloudflare.F("postgres"),
					AccessClientID: cloudflare.F("0123456789abcdef0123456789abcdef.access"),
					Port:           cloudflare.F(int64(5432)),
				}),
				Caching: cloudflare.F(hyperdrive.HyperdriveCachingParam{
					Disabled:             cloudflare.F(false),
					MaxAge:               cloudflare.F(int64(60)),
					StaleWhileRevalidate: cloudflare.F(int64(15)),
				}),
			},
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Hyperdrive.Configs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		hyperdrive.ConfigUpdateParams{
			Hyperdrive: hyperdrive.HyperdriveParam{
				Name: cloudflare.F("example-hyperdrive"),
				Origin: cloudflare.F(hyperdrive.ConfigurationParam{
					Database:       cloudflare.F("postgres"),
					Host:           cloudflare.F("database.example.com"),
					Scheme:         cloudflare.F(hyperdrive.ConfigurationSchemePostgres),
					User:           cloudflare.F("postgres"),
					AccessClientID: cloudflare.F("0123456789abcdef0123456789abcdef.access"),
					Port:           cloudflare.F(int64(5432)),
				}),
				Caching: cloudflare.F(hyperdrive.HyperdriveCachingParam{
					Disabled:             cloudflare.F(false),
					MaxAge:               cloudflare.F(int64(60)),
					StaleWhileRevalidate: cloudflare.F(int64(15)),
				}),
			},
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigList(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Hyperdrive.Configs.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Hyperdrive.Configs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigEditWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Hyperdrive.Configs.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		hyperdrive.ConfigEditParams{
			Hyperdrive: hyperdrive.HyperdriveParam{
				Caching: cloudflare.F(hyperdrive.HyperdriveCachingParam{
					Disabled:             cloudflare.F(false),
					MaxAge:               cloudflare.F(int64(60)),
					StaleWhileRevalidate: cloudflare.F(int64(15)),
				}),
				Name: cloudflare.F("example-hyperdrive"),
				Origin: cloudflare.F(hyperdrive.ConfigurationParam{
					Database:       cloudflare.F("postgres"),
					Host:           cloudflare.F("database.example.com"),
					Scheme:         cloudflare.F(hyperdrive.ConfigurationSchemePostgres),
					User:           cloudflare.F("postgres"),
					AccessClientID: cloudflare.F("0123456789abcdef0123456789abcdef.access"),
					Port:           cloudflare.F(int64(5432)),
				}),
			},
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Hyperdrive.Configs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
