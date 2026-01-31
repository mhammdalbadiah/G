package main

import (
	"G/cmd"
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		cmd.Help()
		return
	}

	command := args[1]

	switch command {

	case "push":
		cmd.Push(args[2:])

	case "pull":
		cmd.Pull()

	case "clone":
		cmd.Clone(args[2:])

	case "status":
		cmd.Status()

	case "help":
		cmd.Help()

	default:
		fmt.Println("####################################")
		fmt.Println("	 Unknown command", command)
		fmt.Println("####################################")
		cmd.Help()
	}

}
