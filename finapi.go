package finapi

import (
	"context"

	"github.com/antihax/optional"
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
	apiClient := NewAPIClient(&Configuration{BasePath: cfg.Endpoint})
	return &Client{
		APIClient:    apiClient,
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
	}
}

// NewClientContext creates a new context for general client actions like creating users.
func (c *Client) NewClientContext() (context.Context, error) {
	token, _, err := c.AuthorizationApi.GetToken(context.Background(), "client_credentials", c.clientID, c.clientSecret, nil)
	if err != nil {
		return nil, err
	}

	ctx := context.WithValue(context.Background(), ContextAccessToken, token.AccessToken)
	return ctx, nil
}

// NewUserContext creates a context for user based actions like connection new banks, gettig transactions etc.
func (c *Client) NewUserContext(finAPIUserID string, password string) (context.Context, error) {
	opts := &AuthorizationApiGetTokenOpts{Username: optional.NewString(finAPIUserID), Password: optional.NewString(password)}
	token, _, err := c.AuthorizationApi.GetToken(context.Background(), "password", c.clientID, c.clientSecret, opts)
	if err != nil {
		return nil, err
	}

	ctx := context.WithValue(context.Background(), ContextAccessToken, token.AccessToken)
	return ctx, nil
}
