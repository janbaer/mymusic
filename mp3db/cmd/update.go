package cmd

import (
	"fmt"
	"path"

	"github.com/janbaer/mp3db/files"
	"github.com/janbaer/mp3db/logger"
	"github.com/janbaer/mp3db/storage"
	"github.com/janbaer/mp3db/tasks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the import command
var updateCmd = &cobra.Command{
	Use:   "update ~/Music",
	Short: "Updates the database with all mp3 files that are under the given directory",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		databaseFilePath := viper.GetString("database")
		dataDirectoryPath := path.Dir(databaseFilePath)

		database, err := storage.OpenDatabase(databaseFilePath)
		if err != nil {
			fmt.Printf("Unexpected error while opening the database: %v\n", err)
			return
		}

		defer database.Close()

		var importDir string

		if len(args) == 0 {
			importDir = viper.GetString("importDir")
		} else {
			importDir = args[0]
		}

		task := tasks.NewUpdateTask(
			database,
			new(files.FileSystem),
			new(files.RealMP3MetadataReader),
			logger.NewUpdateLogLogger(dataDirectoryPath),
		)

		updateStats, err := task.Execute(importDir)
		if err != nil {
			fmt.Printf("Error while importing from %s: %v", importDir, err)
		}

		if updateStats.FailedFilesCount > 0 {
			fmt.Printf("Update complete, but %d files could not be imported. All failed files was logged to file update.log!\n", updateStats.FailedFilesCount)
		} else {
			fmt.Printf("Update complete, %d files were imported, %d were updated!", updateStats.ImportedFilesCount, updateStats.UpdatedFilesCount)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
