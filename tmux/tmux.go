package tmux

import (
	"go-tmux-sessionizer/exec"
	"os"
)

type Session struct {
	Path string
	Name string
}

func (s *Session) Connect() {
	if !s.isRunning() || !s.hasSession() {
		s.createSession()
	}

	if s.isConnected() {
		s.switchToSession()
	} else {
		s.attachToSession()
	}
}

func (s *Session) hasSession() bool {
	cmd := exec.ArrayCommand{
		Name:    "tmux",
		Command: []string{"has-session", "-t=" + s.Name},
	}

	_, err := cmd.Out()

	if err != nil {
		return false
	} else {
		return true
	}
}

func (s *Session) isRunning() bool {
	cmd := exec.ArrayCommand{
		Name:    "pgrep",
		Command: []string{"tmux"},
	}

	processes, err := cmd.Out()
	if err != nil {
		return false
	}

	return processes != ""
}

func (s *Session) isConnected() bool {
	tmuxEnv := os.Getenv("TMUX")
	return tmuxEnv != ""
}

func (s *Session) createSession() {
	cmd := exec.ArrayCommand{
		Name:    "tmux",
		Command: []string{"new-session", "-ds", s.Name, "-c", s.Path},
	}

	cmd.Run()
}

func (s *Session) attachToSession() {
	cmd := exec.ArrayCommand{
		Name:    "tmux",
		Command: []string{"a", "-t", s.Name},
	}

	cmd.Run()
}

func (s *Session) switchToSession() {
	cmd := exec.ArrayCommand{
		Name:    "tmux",
		Command: []string{"switch-client", "-t", s.Name},
	}

	cmd.Run()
}
