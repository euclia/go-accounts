package organization

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/euclia/go-accounts/models"
)

const (
	organizationPath = "/organizations"
)

// GetOrganization is a method to get an Organization by its ID.
func GetOrganization(orgID string, authToken string, baseURL string, HTTPClient *http.Client) (org models.Organization, err error) {
	var endpoint = baseURL + organizationPath + "/" + orgID

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	// req.URL.RawQuery = q.Encode()
	resp, err := HTTPClient.Do(req)

	var respOrg models.Organization

	if err != nil {
		fmt.Printf(err.Error())
		return respOrg, err
	}

	if resp.StatusCode != 200 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return respOrg, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respOrg)
	return respOrg, err
}
