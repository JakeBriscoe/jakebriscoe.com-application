package database

func InsertSingleTrack(track *Track) error {
	return DB.Create(track).Error
}

func InsertManyTracks(tracks []*Track) error {
	return DB.Create(tracks).Error
}

func RetrieveSingleTrack(id uint) (*Track, error) {
	var track Track
	err := DB.First(&track, id).Error
	if err != nil {
		return nil, err
	}
	return &track, nil
}

func DeleteSingleTrack(id uint) error {
	return DB.Delete(&Track{}, id).Error
}

func GetAllTracks() ([]*Track, error) {
	var tracks []*Track
	err := DB.Find(&tracks).Error
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

func InsertManyArtists(artist []*Artist) error {
	return DB.Create(artist).Error
}

func InsertManyAlbums(Album []*Album) error {
	return DB.Create(Album).Error
}

func InsertManyGenres(Genre []*Genre) error {
	return DB.Create(Genre).Error
}

func InsertManyImages(Image []*Image) error {
	return DB.Create(Image).Error
}

func AnyTrackExists() (bool, error) {
	var count int64
	DB.Model(&Track{}).Count(&count)

	return count == 0, nil
}

// ... similar functions for other models
