package storage

import (
	"strings"

	"github.com/janbaer/mp3db/model"
)

// BuildSearchQuery - Builds the search query
func BuildSearchQuery(searchTerm string, searchOptions model.SearchOptions) (string, []string) {
	var sb strings.Builder
	var values []string

	if searchOptions.SearchArtist {
		sb.WriteString("Artist LIKE ?")
		values = append(values, searchTerm+"%")
	}

	if searchOptions.SearchAlbum {
		if sb.Len() > 0 {
			sb.WriteString(" OR ")
		}
		sb.WriteString("Album LIKE ?")
		values = append(values, searchTerm+"%")
	}

	if searchOptions.SearchTitle {
		if sb.Len() > 0 {
			sb.WriteString(" OR ")
		}
		sb.WriteString("Title LIKE ?")
		values = append(values, searchTerm+"%")
	}

	return sb.String(), values
}
