// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/queues"
)

func TestMessageAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Queues.Messages.Ack(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		queues.MessageAckParams{
			Acks: cloudflare.F([]queues.MessageAckParamsAck{{
				LeaseID: cloudflare.F("eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIn0..Q8p21d7dceR6vUfwftONdQ.JVqZgAS-Zk7MqmqccYtTHeeMElNHaOMigeWdb8LyMOg.T2_HV99CYzGaQuhTyW8RsgbnpTRZHRM6N7UoSaAKeK0"),
			}, {
				LeaseID: cloudflare.F("eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIn0..Q8p21d7dceR6vUfwftONdQ.JVqZgAS-Zk7MqmqccYtTHeeMElNHaOMigeWdb8LyMOg.T2_HV99CYzGaQuhTyW8RsgbnpTRZHRM6N7UoSaAKeK0"),
			}, {
				LeaseID: cloudflare.F("eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIn0..Q8p21d7dceR6vUfwftONdQ.JVqZgAS-Zk7MqmqccYtTHeeMElNHaOMigeWdb8LyMOg.T2_HV99CYzGaQuhTyW8RsgbnpTRZHRM6N7UoSaAKeK0"),
			}}),
			Retries: cloudflare.F([]queues.MessageAckParamsRetry{{
				DelaySeconds: cloudflare.F(10.000000),
				LeaseID:      cloudflare.F("eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIn0..Q8p21d7dceR6vUfwftONdQ.JVqZgAS-Zk7MqmqccYtTHeeMElNHaOMigeWdb8LyMOg.T2_HV99CYzGaQuhTyW8RsgbnpTRZHRM6N7UoSaAKeK0"),
			}, {
				DelaySeconds: cloudflare.F(10.000000),
				LeaseID:      cloudflare.F("eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIn0..Q8p21d7dceR6vUfwftONdQ.JVqZgAS-Zk7MqmqccYtTHeeMElNHaOMigeWdb8LyMOg.T2_HV99CYzGaQuhTyW8RsgbnpTRZHRM6N7UoSaAKeK0"),
			}, {
				DelaySeconds: cloudflare.F(10.000000),
				LeaseID:      cloudflare.F("eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIn0..Q8p21d7dceR6vUfwftONdQ.JVqZgAS-Zk7MqmqccYtTHeeMElNHaOMigeWdb8LyMOg.T2_HV99CYzGaQuhTyW8RsgbnpTRZHRM6N7UoSaAKeK0"),
			}}),
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

func TestMessagePullWithOptionalParams(t *testing.T) {
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
	_, err := client.Queues.Messages.Pull(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		queues.MessagePullParams{
			BatchSize:           cloudflare.F(10.000000),
			VisibilityTimeoutMs: cloudflare.F(1000.000000),
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
