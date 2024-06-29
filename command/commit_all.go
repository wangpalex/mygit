package command

import (
	"flag"
	"fmt"
	"os"
)

type CommitAllCommand struct {
	commitMsg string
}

func NewCommitAllCommand(args []string) CommitAllCommand {
	parser := flag.NewFlagSet("commit_all", flag.ExitOnError)
	commitMsg := parser.String("m", "", "Commit message")
	parser.Parse(args)
	if *commitMsg == "" {
		fmt.Println("Error: -m <commit msg> flag is required")
		flag.Usage()
		os.Exit(1)
	}

	return CommitAllCommand{commitMsg: *commitMsg}
}

func (c CommitAllCommand) Exec() {
	runCommand("git add .")
	runCommand(fmt.Sprintf("git commit -m \"%s\"", c.commitMsg))
}
