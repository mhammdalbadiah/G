package cmd

import "fmt"

func Pull() {
	if err := Run("git", "pull"); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
