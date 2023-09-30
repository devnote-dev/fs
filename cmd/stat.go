package cmd

import (
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

var statCmd = &cobra.Command{
	Use:  "stat path",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		info, err := os.Lstat(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "file not found")
			return
		}

		name := info.Name()
		var tname string

		if info.Mode()&os.ModeSymlink != 0 {
			tname = "symlink "
			if tg, err := os.Readlink(info.Name()); err == nil {
				name = name + " -> " + tg
			}
			info, _ = os.Stat(info.Name())
		}

		if info.IsDir() {
			tname = tname + "directory"
		} else if info.Mode()&0111 == 0111 {
			tname = tname + "executable"
		} else if info.Mode().IsRegular() {
			tname = tname + "file"
		} else {
			if tname == "" {
				tname = "unknown file"
			} else {
				tname = "unknown " + tname
			}
		}

		fmt.Printf("   name: %s\n", name)
		fmt.Printf("   type: %s\n", tname)
		fmt.Printf("   size: %s (%d)\n", humanize.IBytes(uint64(info.Size())), info.Size())
		fmt.Printf("  perms: %v (0%d)\n", info.Mode(), uint32(info.Mode()))
		fmt.Printf("modtime: %v\n", info.ModTime())
	},
}
