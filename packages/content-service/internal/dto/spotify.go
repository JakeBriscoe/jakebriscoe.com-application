package dto

type SpotifyTrackRecommendationResponse struct {
	Tracks []SpotifyTrack `json:"tracks"`
}

type SpotifyTrack struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Artists    []SpotifyArtist `json:"artists"`
	Album      SpotifyAlbum    `json:"album"`
	Explicit   bool            `json:"explicit"`
	Popularity int             `json:"popularity"`
	PreviewUrl string          `json:"preview_url"`
}

type SpotifyArtist struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Popularity int      `json:"popularity"`
	Genres     []string `json:"genres"`
	Images     []Image  `json:"images"`
}

type SpotifyAlbum struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Artists    []SpotifyArtist `json:"artists"`
	Images     []Image         `json:"images"`
	Popularity int             `json:"popularity"`
	Genres     []string        `json:"genres"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
