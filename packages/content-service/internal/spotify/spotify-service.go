package spotify

import (
	"content-service/internal/dto"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Move the seed to the caller
func GetTacksFromSeed() ([]dto.SpotifyTrack, error) {

	return getSampleResponse(), nil

	baseURL := "https://api.spotify.com/v1"
	endpoint := "/recommendations"

	// max of 5 total seeds
	queryParams := url.Values{}
	queryParams.Set("seed_artists", getArtistSeeds(5))
	// queryParams.Set("seed_genres", "")
	// queryParams.Set("seed_tracks", "")
	url := baseURL + endpoint + "?" + queryParams.Encode()

	accessToken, err := GetSpotifyAuthToken()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	bearer := "Bearer " + accessToken
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	log.Println("making request to spotify")
	if err != nil {
		log.Fatalln("error", err)
	}
	defer resp.Body.Close()

	// ---------------------------------------------------------------------------------
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal("error", err)
	// }

	// // Log the raw JSON response
	// log.Println("response body:", string(body))

	// // Write the response body to a file
	// if err := ioutil.WriteFile("sample_response.json", body, os.ModePerm); err != nil {
	// 	log.Fatal("error", err)
	// }
	// ---------------------------------------------------------------------------------

	var respBody dto.SpotifyTrackRecommendationResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		log.Fatal("error", err)
	}

	tracks := respBody.Tracks

	return tracks, nil
}

func getArtistSeeds(numberOfArtists int) string {
	artistSeeds := []string{
		"2RQXRUsr4IW1f3mKyKsy4B", // Noah Kahan
		"5KNNVgR6LBIABRIomyCwKJ", // Dermot Kennedy
		"2NjfBq1NflQcKSeiDooVjY", // Tones and I
		"1caoBfXJrbKCwIaTzGkyHn", // Six60
		"4GNC7GD6oZMSxPGyXy4MNB", // Lewis Capaldi
	}

	// Randomly select 5 artists from the array
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(artistSeeds), func(i, j int) {
		artistSeeds[i], artistSeeds[j] = artistSeeds[j], artistSeeds[i]
	})
	artistSeeds = artistSeeds[:numberOfArtists]

	// Join the artist seeds into a single string with commas
	return strings.Join(artistSeeds, ",")
}

func getSampleResponse() []dto.SpotifyTrack {
	data, err := ioutil.ReadFile("sample_response.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the response into a Go struct
	var response []dto.SpotifyTrack
	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
