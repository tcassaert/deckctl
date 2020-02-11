/*
Copyright © 2020 Thomas Cassaert <tcassaert@inuits.eu>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tcassaert/deckctl/lib"
)

var cfgFile string
var username string
var password string
var endpoint string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "deckctl",
	Short: "Manage NextCloud Deck on the CLI",
	Long: `Deckctl is a CLI application to manage the NextCloud Deck app. You can list, create, delete and modify boards, stacks and cards.

  deckctl list cards --board myboard --stack todo
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.deckctl.yaml)")
	rootCmd.PersistentFlags().StringVarP(&username, "user", "u", "", "username")
	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password")
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	rootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "", "Base URL")
	viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint"))

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".deckctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".deckctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// NewHTTPClient initialization
func NewHTTPClient() lib.Client {
	username := viper.GetString("user")
	password := viper.GetString("password")
	endpoint := viper.GetString("endpoint")
	if endpoint == "" {
		log.Fatal("Please set an endpoint.")
	}
	c := lib.Client{
		Username: username,
		Password: password,
		Endpoint: endpoint,
	}
	return c
}
