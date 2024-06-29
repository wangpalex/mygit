package command

import (
	"fmt"
	"os"
)

type SwitchCommand struct {
	branch string
}

func NewSwitchCommand(args []string) SwitchCommand {
	if len(args) < 1 {
		fmt.Printf("Usage: mygit switch <branch>")
		os.Exit(1)
	}
	return SwitchCommand{branch: args[0]}
}

func (c SwitchCommand) Exec() {
	runCommand("git stash")
	runCommand(fmt.Sprintf("git checkout %s", c.branch))
	runCommand("git stash pop")
}
