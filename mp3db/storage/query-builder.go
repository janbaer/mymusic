package storage

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/janbaer/mp3db/model"
)

var r regexp.Regexp

func init() {
	r = *regexp.MustCompile(`(?P<Artist>.+)\s-\s(?P<Title>.+)`)
}

// BuildSearchQuery - Builds the search query
func BuildSearchQuery(searchTerm string, searchOptions model.SearchOptions) (string, []string) {
	var sb strings.Builder
	var values []string

	if r.MatchString(searchTerm) {
		return parseSearchTerm(searchTerm)
	}

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

func parseSearchTerm(searchTerm string) (string, []string) {
	var sb strings.Builder
	var values []string

	strings := r.FindStringSubmatch(searchTerm)

	fmt.Printf("Result %v\n", strings)

	sb.WriteString("Artist LIKE ?")
	values = append(values, strings[1]+"%")

	sb.WriteString(" AND ")

	sb.WriteString("Title LIKE ?")
	values = append(values, strings[2]+"%")
	return sb.String(), values
}
