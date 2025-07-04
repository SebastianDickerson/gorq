package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gorq",
	Short: "gorq is a CLI tool for managing queue tasks and workers",
	Long: `gorq is a CLI tool for managing queue tasks and workers. 
				It provides commands to create, list, and manage tasks
				and workers in a queue system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gorq! Use 'gorq --help' to see available commands.")
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
