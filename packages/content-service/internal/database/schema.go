package database

import (
	"gorm.io/gorm"
)

// can get further track details from /audio-analysis
type Track struct {
	gorm.Model
	Name       string
	PreviewURL string
	Popularity int
	IsExplicit bool
	AlbumID    uint
	Album      Album
	Artists    []*Artist `gorm:"many2many:track_artists;"`
	SpotifyID  string    `gorm:"unique"`
}

type Artist struct {
	gorm.Model
	Name string
	// Followers  int
	Popularity int
	Albums     []*Album `gorm:"many2many:album_artists;"`
	Tracks     []*Track `gorm:"many2many:track_artists;"`
	Genres     []*Genre `gorm:"many2many:artist_genres;"`
	Images     []*Image `gorm:"many2many:artist_images;"`
	SpotifyID  string   `gorm:"unique"`
}

type Album struct {
	gorm.Model
	Name       string
	Popularity int
	Artists    []*Artist `gorm:"many2many:album_artists;"`
	Tracks     []*Track  `gorm:"many2many:album_tracks;"`
	Genres     []*Genre  `gorm:"many2many:album_genres;"`
	Images     []*Image  `gorm:"many2many:album_images;"`
	SpotifyID  string    `gorm:"unique"`
}

type Image struct {
	URL    string
	Width  int
	Height int
}

type Genre struct {
	gorm.Model
	Name    string    `gorm:"unique"`
	Albums  []*Album  `gorm:"many2many:album_genres;"`
	Artists []*Artist `gorm:"many2many:artist_genres;"`
}

// type Playlist struct {
//     gorm.Model
// }
