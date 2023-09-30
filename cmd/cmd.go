package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

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

func getFileType(info os.FileInfo, sym bool) string {
	var t string
	if sym {
		t = "symlink "
	}

	if info.IsDir() {
		return t + "directory"
	} else if info.Mode()&0111 == 0111 {
		return t + "executable"
	} else if info.Mode().IsRegular() {
		return t + "file"
	} else {
		if t == "" {
			return "unknown file"
		} else {
			return "unknown " + t
		}
	}
}
