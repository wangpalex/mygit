package command

import (
	"flag"
	"fmt"
)

type RebaseCommand struct {
	branch    string
	noSquash  bool
	commitMsg string
}

func NewRebaseCommand(args []string) RebaseCommand {
	parser := flag.NewFlagSet("rebase", flag.ExitOnError)
	branch := parser.String("b", "master", "Branch to rebase onto (default is 'master')")
	noSquash := parser.Bool("no_squash", false, "Do not squash commits, just rebase")
	commitMessage := parser.String("m", "", "Commit message for the squashed commit")
	parser.Parse(args)

	return RebaseCommand{
		noSquash:  *noSquash,
		branch:    *branch,
		commitMsg: *commitMessage,
	}
}

func (c RebaseCommand) Exec() {
	fmt.Println("Fetching origin")
	runCommand("git fetch origin")

	fmt.Printf("Rebasing to origin/%s\n", c.branch)
	runCommand(fmt.Sprintf("git rebase origin/%s", c.branch))

	if c.noSquash {
		fmt.Println("Rebase complete")
		return
	}
	fmt.Println("Squashing commits")

	commitMsg := c.commitMsg
	if commitMsg == "" {
		// Use first commit's message
		firstCommit := runCommand(fmt.Sprintf("git rev-list --reverse HEAD ^origin/%s | head -n 1", c.branch))
		commitMsg = runCommand(fmt.Sprintf("git log --format=%%B -n 1 %s", firstCommit))
	}

	resetCommit := runCommand(fmt.Sprintf("git rev-list origin/%s --max-count=1", c.branch))
	runCommand(fmt.Sprintf("git reset --soft %s", resetCommit))
	runCommand(fmt.Sprintf("git commit -m \"%s\"", commitMsg))

	fmt.Println("Rebase and squash complete")
}
