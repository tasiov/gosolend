package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var reservesCmd = &cobra.Command{
    Use:   "reserves",
    Short: "Get information about Solend reserves",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Using RPC URL: %s\n", viper.GetString("rpc-url"))
        // Implement reserve fetching logic here
    },
}

func init() {
    rootCmd.AddCommand(reservesCmd)
}

