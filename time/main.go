package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// I can't believe this doesn't exist in Windows/Command Prompt already
// and for some reason Powershell's "cmdlet" thing is just poor.

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "missing command input\n")
		os.Exit(1)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	st := time.Now()
	_ = cmd.Run()
	tk := time.Now().Sub(st)

	fmt.Printf("\n%s\n", tk)
}
