/*
Package cmd contains a set of commands for the console
*/
package cmd

import (
	"log"
	"netshare/internal/config"
	"netshare/internal/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var cfg config.Config
var serverHost string
var serverPort int
var serverShareDir string
var serverType string

var rootCmd = NewRootCmd()

// NewRootCmd is a constructor for rootcmd.
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "netshare -p 8182 -s localhost -d data -t web",
		Short: "Sharing files over http and may be other protocols.",
		Long: `Currently, only the web Protocol is supported.
	However, the program architecture is designed for different file transfer protocols.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Server host:", cfg.Host)
			log.Println("Server port:", cfg.Port)
			log.Println("Server share dir:", cfg.ShareDir)
			log.Println("Server type:", cfg.Type)

			srv, err := server.Factory(cfg.Type, cfg.Host, cfg.Port, cfg.ShareDir)
			if err != nil {
				panic(err)
			}

			go func() {
				quitChannel := make(chan os.Signal, 1)
				signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)
				<-quitChannel
				log.Println("Close server by signal")
				srv.Stop()
			}()

			srv.Start()
		},
	}
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
	rootCmd.Flags().StringVarP(&serverHost, "host", "s", "", "Setup server host")
	rootCmd.Flags().IntVarP(&serverPort, "serverPort", "p", 0, "Setup server port")
	rootCmd.Flags().StringVarP(&serverShareDir, "dir", "d", "", "Setup server share dir")
	rootCmd.Flags().StringVarP(&serverType, "type", "t", "web", "Setup server type")
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
		viper.SetConfigName(".netshare")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())

		if err := viper.Unmarshal(&cfg); err != nil {
			panic(err)
		}
	}

	log.Println("Config host: ", cfg.Host)
	log.Println("Config port", cfg.Port)
	log.Println("Config share dir: ", cfg.ShareDir)
	log.Println("Config server type: ", cfg.Type)

	if serverHost != "" {
		cfg.Host = serverHost
		log.Println("Setup host from arguments: ", cfg.Host)
	}

	if serverPort != 0 {
		cfg.Port = serverPort
		log.Println("Setup port from arguments: ", cfg.Port)
	}

	if serverShareDir != "" {
		cfg.ShareDir = serverShareDir
		log.Println("Setup share dir from arguments", cfg.ShareDir)
	}

	if serverType != "" {
		cfg.Type = serverType
		log.Println("Setup share dir from arguments", cfg.Type)
	}
}
