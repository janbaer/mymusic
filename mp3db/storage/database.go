package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/janbaer/mp3db/model"
)

// Database - provides access to the MP3 database
type Database struct {
	db *sql.DB
}

type dbRow interface {
	Scan(dest ...interface{}) error
}

// CreateDatabase - creates and open a database
func CreateDatabase(dbPath string) (*Database, error) {
	os.Remove(dbPath)

	fmt.Printf("Creating new database %s\n", dbPath)

	db, err := openDb(dbPath)
	if err != nil {
		return nil, fmt.Errorf("Unexpected error while opening database %s: %v", dbPath, err)
	}

	sql := `
		create table Songs (
			id integer primary key autoincrement,
			filePath text not null,
			artist text,
			title text,
			album text,
			genre text
		)
	`

	_, err = db.Exec(sql)
	if err != nil {
		return nil, fmt.Errorf("Unexpected error while creating table music %s", err)
	}

	return &Database{db}, nil
}

// OpenDatabase - opens an existing storage with the given dbPath
func OpenDatabase(dbPath string) (*Database, error) {
	fmt.Printf("Using database %s\n", dbPath)

	db, err := openDb(dbPath)
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

// Close - closes the database
func (database *Database) Close() {
	if database.db != nil {
		database.db.Close()
		database.db = nil
	}
}

// Insert the given Song into the songs table
func (database *Database) Insert(song *model.Song) error {
	_, err := database.db.Exec(
		"INSERT INTO Songs (filePath,artist,title,album,genre) VALUES(?, ?, ?, ?, ?)",
		song.FilePath,
		song.Artist,
		song.Title,
		song.Album,
		song.Genre,
	)
	return err
}

// Update the given Song into the songs table
func (database *Database) Update(song *model.Song) error {
	_, err := database.db.Exec(
		"UPDATE Songs SET artist=?, title=?, album=?, genre=?, filePath=? WHERE id=?",
		song.Artist, song.Title, song.Album, song.Genre, song.FilePath, song.ID,
	)
	return err
}

// Delete - Deletes the given song from the database
func (database *Database) Delete(song *model.Song) error {
	_, err := database.db.Exec("DELETE FROM Songs WHERE id=?", song.ID)
	return err
}

// QueryAll - Returns all the Songs we actually have in the database
func (database *Database) QueryAll() (*[]model.Song, error) {
	rows, err := database.db.Query("SELECT * FROM SONGS ORDER BY artist, album, title")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapRowsToSongs(rows)
}

// QuerySongs - Executes a Query and returns the results
func (database *Database) QuerySongs(query string, queryValues []string) (*[]model.Song, error) {
	rows, err := database.db.Query(
		fmt.Sprintf("SELECT * FROM SONGS WHERE %s ORDER BY artist, album, title", query),
		convertQueryValuesToArgs(queryValues)...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapRowsToSongs(rows)
}

// QueryByID - returns the song that is matched to the given id
func (database *Database) QueryByID(id int) (*model.Song, error) {
	row := database.db.QueryRow("SELECT * FROM SONGS WHERE id=?", id)
	return mapRowToSong(row)
}

// QueryFilePath - returns the first song that is matched to the given filePath
func (database *Database) QueryFilePath(filePath string) (*model.Song, error) {
	row := database.db.QueryRow("SELECT * FROM SONGS WHERE FilePath=?", filePath)
	return mapRowToSong(row)
}

func openDb(dbPath string) (*sql.DB, error) {
	return sql.Open("sqlite3", dbPath)
}

func mapRowsToSongs(rows *sql.Rows) (*[]model.Song, error) {
	// It's necessary to initialize here an empty array instead of an empty array
	// otherwise we would later return null instead of an empty array
	// when we try to marshal it to JSON
	songs := make([]model.Song, 0)

	for rows.Next() {
		song, _ := mapRowToSong(rows)
		songs = append(songs, *song)
	}

	return &songs, nil
}

func mapRowToSong(row dbRow) (*model.Song, error) {
	var id int
	var filePath, artist, title, album, genre string

	err := row.Scan(&id, &filePath, &artist, &title, &album, &genre)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	song := model.NewSong(id, filePath, artist, title, album, genre)

	return &song, nil
}

func convertQueryValuesToArgs(queryValues []string) []interface{} {
	args := make([]interface{}, len(queryValues))
	for index, value := range queryValues {
		args[index] = value
	}
	return args
}
