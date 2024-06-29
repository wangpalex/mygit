package main

import (
	"fmt"
	"mygit/command"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mygit <command> [<args>]")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "commit_all":
		command.NewCommitAllCommand(os.Args[2:]).Exec()
	case "rebase":
		command.NewRebaseCommand(os.Args[2:]).Exec()
	case "switch":
		command.NewSwitchCommand(os.Args[2:]).Exec()
	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Usage: mygit <command> [<args>]")
		os.Exit(1)
	}
}
