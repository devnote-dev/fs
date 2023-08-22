package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "fs command [options] arguments",
}

func Execute() {
	rootCmd.Execute()
}
