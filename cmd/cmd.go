package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(statCmd)
}

var rootCmd = &cobra.Command{
	Use: "fs command [options] arguments",
}

func Execute() {
	rootCmd.Execute()
}
