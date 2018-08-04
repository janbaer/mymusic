package tasks

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/janbaer/mp3db/model"
	"github.com/janbaer/mp3db/storage"
)

// ServeTask - This will provide an webserver with an restfull endpoint for searching via HTTP
type ServeTask struct {
	storage storage.Storage
	port    int
}

// SearchOptions - defines in which fields should be searched
type SearchOptions struct {
	SearchArtist bool
	SearchTitle  bool
	SearchAlbum  bool
}

// NewServeTask - creates a new instance of the ServeTask
func NewServeTask(storage storage.Storage, port int) *ServeTask {
	return &ServeTask{storage, port}
}

func allowCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

// Execute - Executes the taks and searches for the given search term
func (task *ServeTask) Execute() error {
	http.HandleFunc("/songs", func(w http.ResponseWriter, r *http.Request) {
		searchTerm, searchOptions := getQueryParams(r.RequestURI)
		searchQuery, values := storage.BuildSearchQuery(searchTerm, *searchOptions)

		var songs *[]model.Song

		if len(searchTerm) > 0 {
			songs, _ = task.storage.QuerySongs(searchQuery, values)
		} else {
			songs, _ = task.storage.QueryAll()
		}

		allowCors(w)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(songs); err != nil {
			log.Println(err)
		}
	})

	listenAddress := fmt.Sprintf(":%d", task.port)
	fmt.Printf("MP3DB is waiting for search-requests on port %d\n", task.port)

	return http.ListenAndServe(listenAddress, nil)
}

func getQueryParams(requestURI string) (string, *model.SearchOptions) {
	requestURL, _ := url.Parse(requestURI)

	query := requestURL.Query()

	if searchTerms, exists := query["q"]; exists {
		_, searchArtist := query["artist"]
		_, searchAlbum := query["album"]
		_, searchTitle := query["title"]

		if !(searchArtist || searchAlbum || searchTitle) {
			searchArtist = true
			searchAlbum = true
			searchTitle = true
		}

		return searchTerms[0], model.NewSearchOptions(searchArtist, searchTitle, searchAlbum)
	}

	return "", model.NewSearchOptions(true, true, true)
}
