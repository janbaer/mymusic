package model

// SearchOptions - defines in which fields should be searched
type SearchOptions struct {
	SearchArtist bool
	SearchTitle  bool
	SearchAlbum  bool
}

// NewSearchOptions - Creates new SearchOptions
func NewSearchOptions(searchArtist bool, searchTitle bool, searchAlbum bool) *SearchOptions {
	return &SearchOptions{
		SearchArtist: searchArtist,
		SearchTitle:  searchTitle,
		SearchAlbum:  searchAlbum,
	}
}
