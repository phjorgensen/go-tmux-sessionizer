package exec

import (
	"log"
	"os/exec"
)

type ArrayCommand struct {
	Name    string
	Command []string
}

func (c ArrayCommand) Run() {
	cmd := exec.Command(c.Name, c.Command...)
	bindOutput(cmd)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (c ArrayCommand) Out() (string, error) {
	cmd := exec.Command(c.Name, c.Command...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return formatOutput(out), nil
}
