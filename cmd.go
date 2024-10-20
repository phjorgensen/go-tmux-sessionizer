package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type command interface {
	run()
	out() (string, error)
}

type stringCommand struct {
	command string
}

func (c stringCommand) run() {
	cmd := exec.Command("bash", "-c", c.command)
	bindOutput(cmd)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (c stringCommand) out() (string, error) {
	cmd := exec.Command("bash", "-c", c.command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return formatOutput(out), nil
}

type arrayCommand struct {
	name    string
	command []string
}

func (c arrayCommand) run() {
	cmd := exec.Command(c.name, c.command...)
	bindOutput(cmd)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (c arrayCommand) out() (string, error) {
	cmd := exec.Command(c.name, c.command...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return formatOutput(out), nil
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
