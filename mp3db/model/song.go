package model

// Song - defines the struct for inserting and querying songs into MP3 database
type Song struct {
	ID       int    `json:"id"`
	FilePath string `json:"filePath"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Album    string `json:"album"`
	Genre    string `json:"genre"`
}

// NewSong - creates a new song with using the given params
func NewSong(id int, filePath string, artist string, title string, album string, genre string) Song {
	return Song{
		ID:       id,
		FilePath: filePath,
		Artist:   artist,
		Title:    title,
		Album:    album,
		Genre:    genre,
	}
}

// UpdateFrom - Updates all fields from the target with the fields from the source song
func (song *Song) UpdateFrom(otherSong *Song) {
	song.Artist = otherSong.Artist
	song.Album = otherSong.Album
	song.Title = otherSong.Title
	song.Genre = otherSong.Genre
	song.FilePath = otherSong.FilePath
}

// TagsAreEqual - compares, if boths songs are equal except the id
func (song *Song) TagsAreEqual(otherSong *Song) bool {
	return song.Artist == otherSong.Artist &&
		song.Album == otherSong.Album &&
		song.Title == otherSong.Title &&
		song.Genre == otherSong.Genre &&
		song.FilePath == otherSong.FilePath
}
