/*
Package cmd contains a set of commands for the console
*/
package cmd

import (
	"log"
	"os"
	"webshare/internal/config"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var cfg config.Config
var serverHost string
var serverPort int

var rootCmd = &cobra.Command{
	Use:   "webshare",
	Short: "Sharing files over http",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Run root cmd")
		log.Println("Server host:", cfg.Host)
		log.Println("Server serverPort:", cfg.Port)
		log.Println(4)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.webshare.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&serverHost, "host", "s", "", "Set up server host")
	rootCmd.Flags().IntVarP(&serverPort, "serverPort", "p", 0, "Set up serverPort")
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
			log.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".webshare" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".webshare")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())

		if err := viper.Unmarshal(&cfg); err != nil {
			panic(err)
		}
	}

	if serverHost != "" {
		cfg.Host = serverHost
	}

	if serverPort != 0 {
		cfg.Port = serverPort
	}
}
