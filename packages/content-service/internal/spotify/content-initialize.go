package spotify

import (
	"content-service/internal/database"
	"log"
	"reflect"
)

func InitializeContent() error {
	log.Print("initing content")

	spotifyDtoTracks, err := GetTacksFromSeed()
	if err != nil {
		return err
	}

	// Initialize empty lists for albums, artists, images, and genres to be inserted into the db
	var albums []*database.Album
	var artists []*database.Artist
	var images []*database.Image
	var genres []*database.Genre
	var tracks []*database.Track

	for _, spotifyDtoTrack := range spotifyDtoTracks {

		album := findOrCreateStruct(map[string]interface{}{
			"SpotifyID": spotifyDtoTrack.Album.ID,
		}, &albums, func() interface{} {
			albumToCreate := &database.Album{
				Name:       spotifyDtoTrack.Album.Name,
				Popularity: spotifyDtoTrack.Album.Popularity,
				Genres:     make([]*database.Genre, len(spotifyDtoTrack.Album.Genres)),
				Images:     make([]*database.Image, len(spotifyDtoTrack.Album.Images)),
				Tracks:     []*database.Track{},
				Artists:    []*database.Artist{},
			}
			for i, genreName := range spotifyDtoTrack.Album.Genres {
				genre := findOrCreateStruct(map[string]interface{}{
					"Name": genreName,
				}, &genres, func() interface{} {
					return &database.Genre{}
				}).(*database.Genre)
				albumToCreate.Genres[i] = genre
			}
			for i, spotifyDtoImage := range spotifyDtoTrack.Album.Images {
				image := findOrCreateStruct(map[string]interface{}{
					"URL":    spotifyDtoImage.URL,
					"Width":  spotifyDtoImage.Width,
					"Height": spotifyDtoImage.Height,
				}, &images, func() interface{} {
					return &database.Image{}
				}).(*database.Image)
				albumToCreate.Images[i] = image
			}
			return albumToCreate
		}).(*database.Album)

		track := &database.Track{
			SpotifyID:  spotifyDtoTrack.ID,
			Name:       spotifyDtoTrack.Name,
			Popularity: spotifyDtoTrack.Popularity,
			IsExplicit: spotifyDtoTrack.Explicit,
			PreviewURL: spotifyDtoTrack.PreviewUrl,
			Album:      *album,
			Artists:    make([]*database.Artist, len(spotifyDtoTrack.Artists)),
		}

		album.Tracks = append(album.Tracks, track)

		// Loop through the artists for this track
		for i, artistDto := range spotifyDtoTrack.Artists {

			artist := findOrCreateStruct(map[string]interface{}{
				"SpotifyID": artistDto.ID,
			}, &artists, func() interface{} {
				artistToCreate := &database.Artist{
					Name:       artistDto.Name,
					Popularity: artistDto.Popularity,
					Genres:     make([]*database.Genre, len(artistDto.Genres)),
					Images:     make([]*database.Image, len(artistDto.Images)),
				}
				for i, genreName := range artistDto.Genres {
					genre := findOrCreateStruct(map[string]interface{}{
						"Name": genreName,
					}, &genres, func() interface{} {
						return &database.Genre{}
					}).(*database.Genre)
					artistToCreate.Genres[i] = genre
				}
				for i, spotifyDtoImage := range artistDto.Images {
					image := findOrCreateStruct(map[string]interface{}{
						"URL":    spotifyDtoImage.URL,
						"Width":  spotifyDtoImage.Width,
						"Height": spotifyDtoImage.Height,
					}, &images, func() interface{} {
						return &database.Image{}
					}).(*database.Image)
					artistToCreate.Images[i] = image
				}
				return artistToCreate
			}).(*database.Artist)

			track.Artists[i] = artist

			artistFoundInAlbum := false
			for _, a := range album.Artists {
				if a.SpotifyID == artist.SpotifyID {
					artistFoundInAlbum = true
					break
				}
			}
			if !artistFoundInAlbum {
				album.Artists = append(album.Artists, artist)
			}
		}

		tracks = append(tracks, track)
	}

	// Use the database to insert albums, artists, images, and genres
	if err := database.InsertManyAlbums(albums); err != nil {
		return err
	}
	if err := database.InsertManyArtists(artists); err != nil {
		return err
	}
	if err := database.InsertManyImages(images); err != nil {
		return err
	}
	if err := database.InsertManyGenres(genres); err != nil {
		return err
	}
	if err := database.InsertManyTracks(tracks); err != nil {
		return err
	}

	return nil
}

func findOrCreateStruct(fields map[string]interface{}, structArrPtr interface{}, newStructFunc func() interface{}) interface{} {
	arrVal := reflect.ValueOf(structArrPtr).Elem()
	if arrVal.Kind() != reflect.Slice {
		panic("findOrCreateStruct: structArr is not a slice")
	}

	for i := 0; i < arrVal.Len(); i++ {
		structVal := arrVal.Index(i)
		found := true
		for fieldName, fieldValue := range fields {
			field := structVal.Elem().FieldByName(fieldName)
			if !field.IsValid() || !reflect.DeepEqual(field.Interface(), fieldValue) {
				found = false
				break
			}
		}
		if found {
			return structVal.Interface()
		}
	}

	newStructVal := reflect.ValueOf(newStructFunc())
	// sets the new struct ID fields to what was being searched for
	for fieldName, fieldValue := range fields {
		field := newStructVal.Elem().FieldByName(fieldName)
		if field.IsValid() {
			field.Set(reflect.ValueOf(fieldValue))
		}
	}

	newArrVal := reflect.Append(arrVal, newStructVal)
	arrVal.Set(newArrVal)

	return newStructVal.Interface()
}
