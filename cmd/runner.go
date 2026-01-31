package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

// Run executes a command and streams its output to the current terminal.
// Example: Run("git", "status")
func Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	// show the exact command (helpful for debugging)
	fmt.Printf(">> %s", name)
	for _, a := range args {
		fmt.Printf(" %s", a)
	}
	fmt.Println()

	// connect command input/output to this terminal
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
