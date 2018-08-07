package cmd

import (
	"fmt"

	"github.com/janbaer/mp3db/files"
	"github.com/janbaer/mp3db/storage"
	"github.com/janbaer/mp3db/tasks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Port - This Flag defines optionally the port
var Port int

// versionCmd represents the version command
var serveCmd = &cobra.Command{
	Use:   "serve --port=8080",
	Short: "Run a webserver which is providing a restapi for querying the MP3DB via http",
	Long:  "The param port is optional, default is 8080",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := storage.OpenDatabase(viper.GetString("database"))
		if err != nil {
			fmt.Printf("Unexpected error while opening the database: %v\n", err)
			return
		}

		defer database.Close()

		port := viper.GetInt("port")
		if Port > 0 {
			port = Port
		}

		task := tasks.NewServeTask(database, new(files.FileSystem), port)

		err = task.Execute()
		if err != nil {
			fmt.Printf("Unexpected while starting web server: %v\n", err)
			return
		}

	},
}

func init() {
	serveCmd.Flags().IntVarP(&Port, "port", "p", 0, "Defines optionally the port")
	rootCmd.AddCommand(serveCmd)
}
