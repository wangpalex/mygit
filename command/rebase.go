package command

import (
	"flag"
	"fmt"
	"os"
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
	// Fetch the latest changes from the remote
	if _, err := runCommand("git fetch origin"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Rebase the current branch onto the specified branch
	if _, err := runCommand(fmt.Sprintf("git rebase origin/%s", c.branch)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if c.noSquash {
		fmt.Println("Rebase complete")
		return
	}
	// Do squash
	// Get the first commit hash and message of the current branch
	firstCommitHash, err := runCommand("git rev-list --reverse HEAD | head -n 1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Reset the branch to the first commit, preserving the working directory
	if _, err := runCommand(fmt.Sprintf("git reset --soft %s", firstCommitHash)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if c.commitMsg != "" {
		// Amend the first commit with all changes and the final commit message
		if _, err := runCommand(fmt.Sprintf("git commit --amend -m \"%s\"", c.commitMsg)); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("Rebase and squash complete")
}
