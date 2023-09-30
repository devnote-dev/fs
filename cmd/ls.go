package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

func init() {
	lsCmd.Flags().BoolP("all", "a", false, "")
	lsCmd.Flags().BoolP("dirs", "d", false, "")
	lsCmd.Flags().BoolP("files", "f", false, "")
	lsCmd.Flags().BoolP("links", "l", false, "")
	lsCmd.Flags().BoolP("size", "s", false, "")
	lsCmd.Flags().BoolP("type", "t", false, "")
}

type entry struct {
	path  string
	tname string
	size  int64
}

var lsCmd = &cobra.Command{
	Use:  "ls [-a | --all] [-d | --dirs] [-f | --files] [-l | --links] [-s | --size] [-t | --type] [dir]",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var dir string
		if len(args) == 1 {
			dir = args[0]
		} else {
			dir, _ = os.Getwd()
		}

		info, err := os.Stat(dir)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintln(os.Stderr, "path not found")
				return
			}

			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}

		if !info.IsDir() {
			fmt.Fprintln(os.Stderr, "path is not a directory")
			return
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}

		dirs, _ := cmd.Flags().GetBool("dirs")
		files, _ := cmd.Flags().GetBool("files")
		links, _ := cmd.Flags().GetBool("links")
		size, _ := cmd.Flags().GetBool("size")
		_type, _ := cmd.Flags().GetBool("type")

		all, _ := cmd.Flags().GetBool("all")
		if all {
			dirs = true
			files = true
			links = true
			size = true
			_type = true
		}

		var res []entry
		maxw := 0

		for _, e := range entries {
			var fsize int64
			if len(e.Name()) > maxw {
				maxw = len(e.Name())
			}

			if size && !e.IsDir() {
				if i, err := os.Stat(e.Name()); err == nil {
					fsize = i.Size()
				}
			}

			if dirs && e.IsDir() {
				res = append(res, entry{e.Name(), "dir", 0})
			}

			if files && e.Type().IsRegular() {
				res = append(res, entry{e.Name(), "file", fsize})
			}

			if links && e.Type()&os.ModeSymlink != 0 {
				res = append(res, entry{e.Name(), "symlink", fsize})
			}
		}

		if len(res) == 0 {
			return
		}

		var b strings.Builder

		for _, e := range res {
			b.WriteString(e.path)

			if _type {
				b.WriteString(strings.Repeat(" ", maxw+3-len(e.path)))
				b.WriteString(e.tname)
			}

			if size {
				b.WriteString(strings.Repeat(" ", 10-len(e.tname)))
				b.WriteString(humanize.IBytes(uint64(e.size)))
			}

			b.WriteRune('\n')
		}

		fmt.Print(b.String())
	},
}
