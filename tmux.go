package main

import (
	"os"
)

type tmux struct {
	sessionPath string
	sessionName string
}

func (tmux *tmux) connect() {
	if !tmux.isRunning() || !tmux.hasSession() {
		tmux.createSession()
	}

	if tmux.isConnected() {
		tmux.switchToSession()
	} else {
		tmux.attachToSession()
	}
}

func (tmux *tmux) hasSession() bool {
	cmd := arrayCommand{
		name:    "tmux",
		command: []string{"has-session", "-t=" + tmux.sessionName},
	}

	_, err := cmd.out()

	if err != nil {
		return false
	} else {
		return true
	}
}

func (tmux *tmux) isRunning() bool {
	cmd := arrayCommand{
		name:    "pgrep",
		command: []string{"tmux"},
	}

	processes, err := cmd.out()
	if err != nil {
		return false
	}

	return processes != ""
}

func (tmux *tmux) isConnected() bool {
	tmuxEnv := os.Getenv("TMUX")
	return tmuxEnv != ""
}

func (tmux *tmux) createSession() {
	cmd := arrayCommand{
		name:    "tmux",
		command: []string{"new-session", "-ds", tmux.sessionName, "-c", tmux.sessionPath},
	}

	cmd.run()
}

func (tmux *tmux) attachToSession() {
	cmd := arrayCommand{
		name:    "tmux",
		command: []string{"a", "-t", tmux.sessionName},
	}

	cmd.run()
}

func (tmux *tmux) switchToSession() {
	cmd := arrayCommand{
		name:    "tmux",
		command: []string{"switch-client", "-t", tmux.sessionName},
	}

	cmd.run()
}
