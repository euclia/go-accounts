package goAccounts

import (
	"net/http"
	"time"

	"github.com/euclia/go-accounts/models"
	"github.com/euclia/go-accounts/organization"
	"github.com/euclia/go-accounts/user"
)

const (
	// ClientVersion is used in User-Agent request header to provide server with API level.
	ClientVersion = "0.0.1"

	// httpClientTimeout is used to limit http.Client waiting time.
	httpClientTimeout = 15 * time.Second
)

// Client object
type Client models.Client

// InitClient creates a Jaqpot Go Client
func InitClient(baseURL string) *Client {
	if baseURL[len(baseURL)-1:] != "/" {
		baseURL = baseURL + "/"
	}
	return &Client{
		C: models.ClientProperties{
			BaseURL: baseURL,
			HTTPClient: &http.Client{
				Timeout: 100000000000,
			},
		},
	}
}

// IAccountsClient is the Accounts Client Interface
type IAccountsClient interface {
	// Get a User by ID.
	GetUser(userID string, authToken string) (user models.User, err error)

	// Get an Organization by ID.
	GetOrganization(orgID string, authToken string) (organization models.Organization, err error)
}

// GetUser is a method to get a User by ID.
func (client *Client) GetUser(userID string, authToken string) (respUser models.User, err error) {
	return user.GetUser(userID, authToken, client.C.BaseURL, client.C.HTTPClient)
}

// GetOrganization is a method to get an Organization by ID.
func (client *Client) GetOrganization(orgID string, authToken string) (respOrganization models.Organization, err error) {
	return organization.GetOrganization(orgID, authToken, client.C.BaseURL, client.C.HTTPClient)
}
