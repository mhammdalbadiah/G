package cmd

import "fmt"

// Help prints usage instructions for the G CLI.
func Help() {
	fmt.Println("G - Simple Git Task Runner")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  G help")
	fmt.Println("  G status")
	fmt.Println("  G pull")
	fmt.Println("  G clone <repo_url> [folder]")
	fmt.Println("  G push <message>")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println(`  G clone https://github.com/user/repo.git`)
	fmt.Println(`  G clone https://github.com/user/repo.git myfolder`)
	fmt.Println(`  G status`)
	fmt.Println(`  G pull`)
	fmt.Println(`  G push "update: fix login bug"`)
	fmt.Println()
}
