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

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import ~/Music",
	Short: "Imports the all mp3 files that are under the given directory into the mp3db",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		databaseFilePath := viper.GetString("database")
		dataDirectoryPath := path.Dir(databaseFilePath)

		database, err := storage.CreateDatabase(databaseFilePath)
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

		task := tasks.NewImportTask(
			database,
			new(files.FileSystem),
			new(files.RealMP3MetadataReader),
			logger.NewImportLogLogger(dataDirectoryPath),
		)

		importStats, err := task.Execute(importDir)
		if err != nil {
			fmt.Printf("Error while importing from %s: %v", importDir, err)
		}

		if importStats.FailedFilesCount > 0 {
			fmt.Printf("Import complete, but %d files could not be imported. All failed files was logged to file import.log!\n", importStats.FailedFilesCount)
		} else {
			fmt.Printf("Import complete, %d files were imported!", importStats.ImportedFilesCount)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
