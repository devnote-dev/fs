package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

	var b strings.Builder
	b.WriteRune('\n')

	if hr := tk.Hours(); hr >= 1.0 {
		fmt.Fprintf(&b, "%1fhrs ", hr)
	}

	if min := tk.Minutes(); min >= 1.0 && min < 60.0 {
		fmt.Fprintf(&b, "%3.fmins ", min)
	}

	if sc := tk.Seconds(); sc >= 1.0 && sc < 60.0 {
		fmt.Fprintf(&b, "%3.fsecs ", sc)
	}

	if ms := tk.Milliseconds(); ms >= 1 && ms < 1000 {
		fmt.Fprintf(&b, "%dms ", ms)
	}

	b.WriteString("taken")
	fmt.Println(b.String())
}
