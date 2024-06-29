package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCommand(cmdStr string) string {
	cmd := exec.Command("sh", "-c", cmdStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(output))
}
