package cmd

import (
	"fmt"
	"strings"

	"github.com/janbaer/mp3db/model"
	"github.com/janbaer/mp3db/storage"
	"github.com/janbaer/mp3db/tasks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	searchArtist bool
	searchTitle  bool
	searchAlbum  bool
)

var searchCmd = &cobra.Command{
	Use:   `search "Michael Jackson" --artist`,
	Short: "Searches for the given search term in the database",
	Long:  `When you don't pass any of the flags the search will be processed for the artist, title and album.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		database, err := storage.OpenDatabase(viper.GetString("database"))
		if err != nil {
			fmt.Printf("Unexpected error while opening the database: %v\n", err)
			return
		}

		defer database.Close()

		searchTerm := args[0]

		task := tasks.NewSearchTask(database)

		if !searchArtist && !searchTitle && !searchAlbum {
			searchArtist = true
			searchTitle = true
			searchAlbum = true
		}

		songs, err := task.Execute(searchTerm, model.SearchOptions{searchArtist, searchTitle, searchAlbum})
		if err != nil {
			fmt.Printf("Unexpected while searching: %v\n", err)
			return
		}

		printSearchResult(searchTerm, songs)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVarP(&searchArtist, "artist", "a", false, "Searches only in the artist field")
	searchCmd.Flags().BoolVarP(&searchTitle, "title", "t", false, "Searches only in the title field")
	searchCmd.Flags().BoolVarP(&searchAlbum, "album", "b", false, "Searches only in the album field")
}

func printSearchResult(searchTerm string, songs *[]model.Song) {
	fmt.Printf("%d songs matched to the given search term: '%s'\n", len(*songs), searchTerm)
	fmt.Println(strings.Repeat("-", 50))

	for _, song := range *songs {
		fmt.Printf("%s\t%s\t%s\t%s\n", song.Artist, song.Title, song.Album, song.FilePath)
	}
}
