package web

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/janbaer/mp3db/model"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

// MP3DbSearch - defines the interface to search for
type MP3DbSearch interface {
	Search(song *model.Song) (bool, error)
}

// MP3DbSearcher - implements the Interface MP3DbSearch
type MP3DbSearcher struct {
	MP3DbServerAddress string
}

// NewMP3DbSearcher - Creates a new MP3DbSearcher
func NewMP3DbSearcher(mp3dbServerAddress string) *MP3DbSearcher {
	return &MP3DbSearcher{MP3DbServerAddress: mp3dbServerAddress}
}

// Search - Searches if the given song exists in the MP3 database
func (mp3DbSearcher MP3DbSearcher) Search(song *model.Song) (bool, error) {
	query := fmt.Sprintf("artist=%s&title=%s", url.QueryEscape(song.Artist), url.QueryEscape(song.Title))
	requestURL := fmt.Sprintf("%s/songs?%s", mp3DbSearcher.MP3DbServerAddress, query)
	response, err := http.Get(requestURL)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var songs []model.Song
	err = decoder.Decode(&songs)
	if err != nil {
		return false, err
	}

	if len(songs) == 0 {
		return false, nil
	}

	return true, nil
}
