package main

import (
	"os"
)

type tmuxSession struct {
	path string
	name string
}

func (s *tmuxSession) connect() {
	if !s.isRunning() || !s.hasSession() {
		s.createSession()
	}

	if s.isConnected() {
		s.switchToSession()
	} else {
		s.attachToSession()
	}
}

func (s *tmuxSession) hasSession() bool {
	cmd := arrayCommand{
		name:    "tmux",
		command: []string{"has-session", "-t=" + s.name},
	}

	_, err := cmd.out()

	if err != nil {
		return false
	} else {
		return true
	}
}

func (s *tmuxSession) isRunning() bool {
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

func (s *tmuxSession) isConnected() bool {
	tmuxEnv := os.Getenv("TMUX")
	return tmuxEnv != ""
}

func (s *tmuxSession) createSession() {
	cmd := arrayCommand{
		name:    "tmux",
		command: []string{"new-session", "-ds", s.name, "-c", s.path},
	}

	cmd.run()
}

func (s *tmuxSession) attachToSession() {
	cmd := arrayCommand{
		name:    "tmux",
		command: []string{"a", "-t", s.name},
	}

	cmd.run()
}

func (s *tmuxSession) switchToSession() {
	cmd := arrayCommand{
		name:    "tmux",
		command: []string{"switch-client", "-t", s.name},
	}

	cmd.run()
}
