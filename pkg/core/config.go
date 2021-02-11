package core

import (
	"fmt"
	"net/url"
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	DiscoveryURL *url.URL
	ClientID     string
	ClientSecret string

	EventsServiceURL *url.URL

	IMAPServerURL *url.URL
	Username      string
	Password      string
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.DiscoveryURL, validation.NotNil),
		validation.Field(&c.ClientID, validation.Required),
		validation.Field(&c.ClientSecret, validation.Required),
		validation.Field(&c.EventsServiceURL, validation.NotNil),
		validation.Field(&c.IMAPServerURL, validation.NotNil),
		validation.Field(&c.Username, validation.Required),
		validation.Field(&c.Password, validation.Required),
	)
}

func LoadConfig() (Config, error) {
	cfg := Config{
		Username: os.Getenv("IMAP_USERNAME"),
		Password: os.Getenv("IMAP_PASSWORD"),
	}

	if rawDiscoveryURL := os.Getenv("DISCOVERY_URL"); rawDiscoveryURL != "" {
		discoveryURL, err := url.Parse(rawDiscoveryURL)
		if err != nil {
			return cfg, fmt.Errorf("error parsing DISCOVERY_URL: %w", err)
		}

		cfg.DiscoveryURL = discoveryURL
	}

	cfg.ClientID = os.Getenv("CLIENT_ID")
	cfg.ClientSecret = os.Getenv("CLIENT_SECRET")

	if rawEventsServiceURL := os.Getenv("EVENTS_SERVICE_URL"); rawEventsServiceURL != "" {
		eventsServiceURL, err := url.Parse(rawEventsServiceURL)
		if err != nil {
			return cfg, fmt.Errorf("error parsing EVENTS_SERVICE_URL: %w", err)
		}

		cfg.EventsServiceURL = eventsServiceURL
	}

	if rawIMAPServerURL := os.Getenv("IMAP_SERVER_URL"); rawIMAPServerURL != "" {
		IMAPServerURL, err := url.Parse(rawIMAPServerURL)
		if err != nil {
			return cfg, fmt.Errorf("parsing IMAP_SERVER_URL: %w", err)
		}

		cfg.IMAPServerURL = IMAPServerURL
	}

	return cfg, nil
}
