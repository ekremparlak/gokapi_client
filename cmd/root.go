/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/ekremparlak/gokapi_client/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apiKey string
	apiURL string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gokapi_client",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config.InitConfig()

	rootCmd.PersistentFlags().StringVarP(&apiKey, "api_key", "k", viper.GetString("GOKAPI_KEY"), "Set the api key for the server")
	rootCmd.PersistentFlags().StringVarP(&apiURL, "api_url", "u", viper.GetString("GOKAPI_URL"), "Set the api url for the server")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
