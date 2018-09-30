package storage_test

import (
	"testing"

	"github.com/janbaer/mp3db/model"
	"github.com/janbaer/mp3db/storage"
	"github.com/stretchr/testify/assert"
)

func TestBuildSearchQuery_when_user_entered_artist_and_title(t *testing.T) {
	searchOptions := model.SearchOptions{}
	searchTerm, values := storage.BuildSearchQuery("4 strings - believe", searchOptions)

	assert.Equal(t, searchTerm, "Artist LIKE ? AND Title LIKE ?")
	assert.Equal(t, len(values), 2)
	assert.Equal(t, "4 strings%", values[0])
	assert.Equal(t, "believe%", values[1])
}
