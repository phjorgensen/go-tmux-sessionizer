package exec

import (
	"log"
	"os/exec"
)

type StringCommand struct {
	Command string
}

func (c StringCommand) Run() {
	cmd := exec.Command("bash", "-c", c.Command)
	bindOutput(cmd)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (c StringCommand) Out() (string, error) {
	cmd := exec.Command("bash", "-c", c.Command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return formatOutput(out), nil
}
