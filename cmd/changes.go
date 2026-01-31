package cmd

import (
	"bytes"
	"os/exec"
	"strings"
)

// HasChanges returns true if there are uncommitted changes (tracked or untracked).
// It checks: git status --porcelain
func HasChanges() (bool, error) {
	c := exec.Command("git", "status", "--porcelain")
	var out bytes.Buffer
	c.Stdout = &out
	c.Stderr = &out

	if err := c.Run(); err != nil {
		return false, err
	}

	// If output is empty -> clean working tree
	return strings.TrimSpace(out.String()) != "", nil
}
