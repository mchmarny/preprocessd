package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	meta "cloud.google.com/go/compute/metadata"
)

var (
	logger  = log.New(os.Stdout, "[EPS] ", 0)
	project = mustEnvVar("PROJECT", "none")
	topic   = mustEnvVar("TOPIC", "processedevents")
	que     *queue
)

func main() {

	if project == "none" {
		mc := meta.NewClient(&http.Client{Transport: userAgentTransport{
			userAgent: "buttons",
			base:      http.DefaultTransport,
		}})
		p, err := mc.ProjectID()
		if err != nil {
			logger.Fatalf("Error creating metadata client: %v", err)
		}
		project = p
	}

	q, err := newQueue(context.Background(), project, topic)
	if err != nil {
		logger.Fatalf("Error creating pubsub client: %v", err)
	}
	que = q

	http.HandleFunc("/", requestHandler)
	port := fmt.Sprintf(":%s", mustEnvVar("PORT", "8080"))
	if err := http.ListenAndServe(port, nil); err != nil {
		logger.Fatal(err)
	}

}

func mustEnvVar(key, fallbackValue string) string {

	if val, ok := os.LookupEnv(key); ok {
		logger.Printf("%s: %s", key, val)
		return strings.TrimSpace(val)
	}

	if fallbackValue == "" {
		logger.Fatalf("Required envvar not set: %s", key)
	}

	logger.Printf("%s: %s (not set, using default)", key, fallbackValue)
	return fallbackValue
}

// GCP Metadata
// https://godoc.org/cloud.google.com/go/compute/metadata#example-NewClient
type userAgentTransport struct {
	userAgent string
	base      http.RoundTripper
}

func (t userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent)
	return t.base.RoundTrip(req)
}
