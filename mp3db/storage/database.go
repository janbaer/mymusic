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
			genre text,
			length text
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
		"INSERT INTO Songs (filePath,artist,title,album,genre,length) VALUES(?, ?, ?, ?, ?, ?)",
		song.FilePath,
		song.Artist,
		song.Title,
		song.Album,
		song.Genre,
		song.Length,
	)
	return err
}

// Update the given Song into the songs table
func (database *Database) Update(song *model.Song) error {
	_, err := database.db.Exec(
		"UPDATE Songs SET artist=?, title=?, album=?, genre=?, filePath=?, length=? WHERE id=?",
		song.Artist, song.Title, song.Album, song.Genre, song.FilePath, song.Length, song.ID,
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
	return database.query("SELECT * FROM SONGS ORDER BY artist, album, title")
}

// QuerySongs - Executes a Query and returns the results
func (database *Database) QuerySongs(query string, queryValues []string) (*[]model.Song, error) {
	return database.query(
		fmt.Sprintf("SELECT * FROM SONGS WHERE %s ORDER BY artist, album, title", query),
		convertQueryValuesToArgs(queryValues)...,
	)
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

// FindDuplicates - Finds all duplicates in the database
func (database *Database) FindDuplicates() (*[]model.Song, error) {
	sql := `
		drop table if exists duplicate_songs;

		create temporary table duplicate_songs (artist text, title text);

		insert into duplicate_songs(artist, title)
  	select artist, title from songs 
		where artist != '' AND title != ''
		group by artist, title having count(*) > 1;
	`
	_, err := database.db.Exec(sql)
	if err != nil {
		return nil, err
	}

	sql = `
		select songs.id,  songs.filePath,  songs.artist, songs.title, songs.album, songs.genre, songs.length from songs join duplicate_songs 
		on songs.artist = duplicate_songs.artist AND songs.title = duplicate_songs.title
		order by songs.artist, songs.title
	`
	return database.query(sql)
}

func (database *Database) query(sql string, args ...interface{}) (*[]model.Song, error) {
	rows, err := database.db.Query(
		sql,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapRowsToSongs(rows)
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
	var filePath, artist, title, album, genre, length string

	err := row.Scan(&id, &filePath, &artist, &title, &album, &genre, &length)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	song := model.NewSong(id, filePath, artist, title, album, genre, length)

	return &song, nil
}

func convertQueryValuesToArgs(queryValues []string) []interface{} {
	args := make([]interface{}, len(queryValues))
	for index, value := range queryValues {
		args[index] = value
	}
	return args
}
