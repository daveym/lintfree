package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd - The base command for the application.
var RootCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint - pull back articles from your Pocket datastore and get rid of the fluff.",
	Long: `Lint provides you with a way of seeing what you've saved to read later. Often we forget what we've stored, 
	and lint aims to help you remember whats stored, but also provide a quick way to purge items you wont read, or retag 
	items for easier retrieval later.

	Run 'lint COMMAND --help' for more information on a command.`,
}

// Execute - Adds all defined commands
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	viper.Get("consumerkey")
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	viper.SetConfigName("lint")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Please make sure that lint.yaml exists in the working directory.")
		os.Exit(-1)
	}
}
