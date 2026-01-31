package cmd

import "fmt"

func Status() {
	fmt.Println("Running: git status")
	if err := Run("git", "status"); err != nil {
		fmt.Println("Error:", err)
	}
}
