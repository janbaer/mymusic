package cmd

import (
	"fmt"
	"log"

	"github.com/janbaer/mp3db/files"
	"github.com/janbaer/mp3db/logger"
	"github.com/janbaer/mp3db/tasks"
	"github.com/janbaer/mp3db/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// verifyImport represents the import command
var verifyImport = &cobra.Command{
	Use:   "verify-import ~/Music",
	Short: "Verifies all mp3 files that are under the given directory and returns all found duplicates",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		mp3dbServerAddress := viper.GetString("mp3db-server")

		var importDir string

		if len(args) == 0 {
			log.Fatal("No import directory passed as argument")
		} else {
			importDir = args[0]
		}

		task := tasks.NewVerifyImportTask(
			new(files.FileSystem),
			new(files.ID3TagReadWriter),
			web.NewMP3DbSearcher(mp3dbServerAddress),
			logger.NewVerifyImportLogLogger(),
		)

		stats, err := task.Execute(importDir)
		if err != nil {
			fmt.Printf("Error while verifying files from %s: %v", importDir, err)
		}

		if stats.FailedFilesCount > 0 {
			fmt.Printf("Verfication of files complete, but %d files could not be verfified.\n", stats.FailedFilesCount)
		} else {
			fmt.Printf("Verifications complete, %d files were identified as duplicated!\n", stats.DuplicatedFilesCount)
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyImport)
}
