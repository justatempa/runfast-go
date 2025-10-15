package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "runfast-go",
	Short: "runfast-go service",
	Long:  `runfast-go is a service`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(tokenCmd)
}
