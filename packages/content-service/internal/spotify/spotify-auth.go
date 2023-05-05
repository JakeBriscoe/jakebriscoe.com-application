package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	clientID     = "CLIENT_ID"
	clientSecret = "CLIENT_SECRET"
	token        AuthResponse
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func GetSpotifyAuthToken() (string, error) {
	// TODO: expires in is just 1 hour, so need to also store issue time
	if token.AccessToken != "" && token.ExpiresIn > int(time.Now().Unix()) {
		return token.AccessToken, nil
	}

	authOptions := url.Values{}
	authOptions.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(
		"POST",
		"https://accounts.spotify.com/api/token",
		strings.NewReader(authOptions.Encode()),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var authResponse AuthResponse
	if err := json.NewDecoder(res.Body).Decode(&authResponse); err != nil {
		panic(err)
	}

	fmt.Printf("Access token: %s", authResponse.AccessToken)

	token = authResponse

	return authResponse.AccessToken, nil
}
