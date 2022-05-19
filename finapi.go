package finapi

import (
	"context"
	"net/http"
	"time"
)

// ClientConfig holds the configuration values for the finAPI client.
type ClientConfig struct {
	Endpoint     string
	ClientID     string
	ClientSecret string
}

// Client bundles the different API services and holds the configuration.
type Client struct {
	*APIClient
	clientID     string
	clientSecret string
}

// NewClient creates a new finAPI client.
func NewClient(cfg ClientConfig) *Client {
	apiClient := NewAPIClient(&Configuration{
		Servers: ServerConfigurations{
			{
				URL: cfg.Endpoint,
			},
		},
		HTTPClient: &http.Client{
			Timeout: 55 * time.Second,
		},
	})
	return &Client{
		APIClient:    apiClient,
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
	}
}

// NewClientContext creates a new context for general client actions like creating users.
func (c *Client) NewClientContext() (context.Context, error) {
	token, _, err := c.AuthorizationApi.GetToken(context.Background()).
		GrantType("client_credentials").
		ClientId(c.clientID).
		ClientSecret(c.clientSecret).
		Execute()
	if err != nil {
		return nil, err
	}

	ctx := context.WithValue(context.Background(), ContextAccessToken, token.AccessToken)
	return ctx, nil
}

// NewUserContext creates a context for user based actions like connection new banks, gettig transactions etc.
func (c *Client) NewUserContext(finAPIUserID string, password string) (context.Context, error) {
	token, _, err := c.AuthorizationApi.GetToken(context.Background()).
		GrantType("password").
		ClientId(c.clientID).
		ClientSecret(c.clientSecret).
		Username(finAPIUserID).
		Password(password).
		Execute()
	if err != nil {
		return nil, err
	}

	ctx := context.WithValue(context.Background(), ContextAccessToken, token.AccessToken)
	return ctx, nil
}
