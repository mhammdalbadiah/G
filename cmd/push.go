package cmd

import (
	"fmt"
	"strings"
)

func Push(args []string) {
	msg := parseMessage(args)
	if strings.TrimSpace(msg) == "" {
		fmt.Println("Error: commit message is required")
		fmt.Println("Ex : G push -m \"hello\"")
		fmt.Println("Or : G push \"hello\"")
		return
	}

	changed, err := HasChanges()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !changed {
		fmt.Println("Nothing to commit. Working tree clean ")
		return
	}

	// 1) add
	if err := Run("git", "add", "."); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 2) commit
	if err := Run("git", "commit", "-m", msg); err != nil {
		fmt.Println("Commit:", err)
		return
	}

	// 3) push
	if err := Run("git", "push"); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func parseMessage(args []string) string {
	if len(args) == 0 {
		return ""
	}

	if args[0] == "-m" {
		if len(args) < 2 {
			return ""
		}
		return strings.Join(args[1:], " ")
	}

	return strings.Join(args, " ")
}
