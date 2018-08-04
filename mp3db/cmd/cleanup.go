package cmd

import (
	"fmt"

	"github.com/janbaer/mp3db/files"
	"github.com/janbaer/mp3db/storage"
	"github.com/janbaer/mp3db/tasks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup the database with deleting songs that are no longer existing in the file system",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := storage.OpenDatabase(viper.GetString("database"))
		if err != nil {
			fmt.Printf("Unexpected error while cleaning up the database: %v\n", err)
			return
		}

		defer database.Close()

		task := tasks.NewCleanupTask(database, new(files.FileSystem))

		cleanupStats, err := task.Execute()
		if err != nil {
			fmt.Printf("Unexpected error while cleaning up the database: %v\n", err)
			return
		}

		fmt.Printf(
			"Cleanup finished, %d songs checked and %d songs has been deleted",
			cleanupStats.TotalCountOfSongs,
			cleanupStats.TotalCountOfDeletedSongs,
		)
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
