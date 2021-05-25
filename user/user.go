package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/euclia/go-accounts/models"
)

const (
	userPath = "users"
)

// GetUser is a method to get User by its ID.
func GetUser(userID string, authToken string, baseURL string, HTTPClient *http.Client) (user models.User, err error) {
	var endpoint = baseURL + userPath

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	q := req.URL.Query()
	q.Add("id", userID)
	q.Add("min", "")
	q.Add("max", "")
	q.Add("email", "")

	// req.URL.RawQuery = q.Encode()
	resp, err := HTTPClient.Do(req)

	var respUser models.User

	if err != nil {
		fmt.Printf(err.Error())
		return respUser, err
	}

	if resp.StatusCode != 200 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return respUser, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respUser)
	return respUser, err
}
