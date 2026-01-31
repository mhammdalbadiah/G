package cmd

import "fmt"

func Clone(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: repository URL is required")
		fmt.Println("Ex   : G clone https://github.com/user/repo.git")
		return
	}

	repoURL := args[0]

	if len(args) >= 2 {
		folder := args[1]
		if err := Run("git", "clone", repoURL, folder); err != nil {
			fmt.Println("Error:", err)
		}
		return
	}

	if err := Run("git", "clone", repoURL); err != nil {
		fmt.Println("Error:", err)
	}
}
