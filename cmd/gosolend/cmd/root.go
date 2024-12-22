package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	keypairPath string
	rpcURL      string
	apiKey      string
)

var rootCmd = &cobra.Command{
	Use:   "gosolend",
	Short: "A CLI for interacting with Solend Protocol",
	Long: `gosolend is a command line interface for interacting with the Solend Protocol
on the Solana blockchain. It provides commands for managing lending positions,
viewing market data, and more.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	rootCmd.PersistentFlags().StringVar(&keypairPath, "keypair", "", "path to keypair file")
	rootCmd.PersistentFlags().StringVar(&rpcURL, "rpc-url", "", "RPC server URL")
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "API key for Solend")

	// Bind flags to viper
	viper.BindPFlag("keypair", rootCmd.PersistentFlags().Lookup("keypair"))
	viper.BindPFlag("rpc-url", rootCmd.PersistentFlags().Lookup("rpc-url"))
	viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))
}

func initConfig() {
	// Get working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Explicitly set the config file path
	configPath := filepath.Join(wd, "config.yaml")
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// Read environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GOSOLEND")

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())

	// Debug: Print actual values being read
	fmt.Printf("Loaded values:\n")
	fmt.Printf("keypair: %s\n", viper.GetString("keypair"))
	fmt.Printf("rpc-url: %s\n", viper.GetString("rpc-url"))
	fmt.Printf("api-key: %s\n", viper.GetString("api-key"))
}
