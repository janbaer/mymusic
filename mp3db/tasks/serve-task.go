package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/janbaer/mp3db/files"
	"github.com/janbaer/mp3db/model"
	"github.com/janbaer/mp3db/storage"
)

// ServeTask - This will provide an webserver with an restfull endpoint for searching via HTTP
type ServeTask struct {
	storage    storage.Storage
	fileAccess files.FileAccess
	id3Writer  files.ID3Writer
	port       int
}

// SearchOptions - defines in which fields should be searched
type SearchOptions struct {
	SearchArtist bool
	SearchTitle  bool
	SearchAlbum  bool
}

// NewServeTask - creates a new instance of the ServeTask
func NewServeTask(storage storage.Storage, fileAccess files.FileAccess, id3Writer files.ID3Writer, port int) *ServeTask {
	return &ServeTask{storage, fileAccess, id3Writer, port}
}

// Execute - Executes the taks and searches for the given search term
func (task *ServeTask) Execute() error {
	router := mux.NewRouter()

	router.HandleFunc("/songs/{id:[0-9]+}", task.handleGetSong).Methods("GET")
	router.HandleFunc("/songs/{id:[0-9]+}", task.handleDeleteSong).Methods("DELETE")
	router.HandleFunc("/songs/{id:[0-9]+}", task.handlePutSong).Methods("PUT")
	router.HandleFunc("/songs", task.handleGetSongs).Methods("GET")

	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "PUT", "DELETE", "OPTIONS"})

	listenAddress := fmt.Sprintf(":%d", task.port)
	fmt.Printf("MP3DB is waiting for search-requests on port %d\n", task.port)

	return http.ListenAndServe(listenAddress, handlers.CORS(allowedOrigins, allowedMethods)(router))
}

func (task *ServeTask) handleGetSongs(w http.ResponseWriter, r *http.Request) {
	searchTerm, searchOptions := getQueryParams(r.URL)
	searchQuery, values := storage.BuildSearchQuery(searchTerm, *searchOptions)

	var songs *[]model.Song

	if len(searchTerm) > 0 {
		songs, _ = task.storage.QuerySongs(searchQuery, values)
	} else {
		songs, _ = task.storage.QueryAll()
	}

	payload, _ := json.Marshal(songs)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(payload)
}

func (task *ServeTask) handleGetSong(w http.ResponseWriter, r *http.Request) {
	songID, _ := strconv.Atoi(mux.Vars(r)["id"])

	song, err := task.storage.QueryByID(songID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	payload, _ := json.Marshal(song)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(payload)
}

func (task *ServeTask) handleDeleteSong(w http.ResponseWriter, r *http.Request) {
	songID, _ := strconv.Atoi(mux.Vars(r)["id"])

	song, _ := task.storage.QueryByID(songID)
	if song != nil {
		if err := task.storage.Delete(song); err == nil {
			task.fileAccess.DeleteFile(song.FilePath)
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (task *ServeTask) handlePutSong(w http.ResponseWriter, r *http.Request) {
	songID, _ := strconv.Atoi(mux.Vars(r)["id"])

	songToUpdate := model.Song{}
	err := json.NewDecoder(r.Body).Decode(&songToUpdate)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	song, _ := task.storage.QueryByID(songID)
	if song == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	song.UpdateFrom(&songToUpdate)

	if err := task.storage.Update(song); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := task.id3Writer.Write(song.FilePath, song); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getQueryParams(requestURL *url.URL) (string, *model.SearchOptions) {
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
