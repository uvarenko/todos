package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "app commands",
}

var start = &cobra.Command{
	Use:   "start",
	Short: "start the app",
}

func init() {
	rootCmd.AddCommand(start)
}

func Execute(run func(cmd *cobra.Command, args []string)) error {
	start.Run = run
	return rootCmd.Execute()
}
