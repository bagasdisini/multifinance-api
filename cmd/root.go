package cmd

import (
	"fmt"
	"github.com/bagasdisini/multifinance-api/cmd/api"
	"github.com/bagasdisini/multifinance-api/cmd/seed"
	"github.com/bagasdisini/multifinance-api/version"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Start the API server",
	Run: func(cmd *cobra.Command, args []string) {
		api.RunServer()
	},
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database",
	Run: func(cmd *cobra.Command, args []string) {
		seed.RunSeed()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Multifinance API version:", version.Version)
	},
}

func Execute() {
	rootCmd.AddCommand(seedCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
