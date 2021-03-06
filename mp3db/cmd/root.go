package cmd

import (
	"fmt"
	"os"

	"github.com/janbaer/mp3db/constants"
	"github.com/janbaer/mp3db/files"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "db",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mp3db.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/mp3db")
		viper.SetConfigName("mp3db-config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Unexpected error while reading config file %s: %v", viper.ConfigFileUsed(), err)
		os.Exit(1)
	}

	fmt.Println("Using config-file:", viper.ConfigFileUsed())

	dbPath := viper.GetString("database")
	if len(dbPath) == 0 {
		dbPath = constants.DefaultDbPath
		viper.Set("database", dbPath)
	}

	dbPath = files.ExpandHomeDirIfNeeded(dbPath)
	viper.Set("database", dbPath)

	importDir := viper.GetString("importDir")
	if len(importDir) > 0 {
		importDir = files.ExpandHomeDirIfNeeded(importDir)
		viper.Set("importDir", importDir)
	}

	port := viper.GetInt("port")
	if port == 0 {
		port = 8080
		viper.Set("port", port)
	}
}
