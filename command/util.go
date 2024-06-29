package command

import (
	"fmt"
	"os/exec"
	"strings"
)

func runCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command '%s' failed with error:\n%s", command, output)
	}
	return strings.TrimSpace(string(output)), nil
}
