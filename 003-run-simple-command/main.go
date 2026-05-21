package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "Hello from Go!")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	cmdtop := exec.Command("top", "-b", "-n", "1") // NOTE: this command should NOT be interactive
	cmdtop.Stdout = os.Stdout
	cmdtop.Stderr = os.Stderr
	if err := cmdtop.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)

	}
}
