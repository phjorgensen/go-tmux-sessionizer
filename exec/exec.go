package exec

import (
	"os"
	"os/exec"
	"strings"
)

type command interface {
	Run()
	Out() (string, error)
}

func bindOutput(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}

func formatOutput(output []byte) string {
	outString := string(output)
	outString = strings.TrimSuffix(outString, "\n")
	return outString
}
