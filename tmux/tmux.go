package tmux

import (
	"go-tmux-sessionizer/exec"
	"os"
)

func Connect(name string) {
	if isConnected() {
		switchToSession(name)
	} else {
		attachToSession(name)
	}
}

func HasSession(name string) bool {
	cmd := exec.ArrayCommand{
		Name:    "tmux",
		Command: []string{"has-session", "-t=" + name},
	}

	_, err := cmd.Out()

	if err != nil {
		return false
	} else {
		return true
	}
}

func CreateSession(name string, path string) {
	cmd := exec.ArrayCommand{
		Name:    "tmux",
		Command: []string{"new-session", "-ds", name, "-c", path},
	}

	cmd.Run()
}

func isConnected() bool {
	tmuxEnv := os.Getenv("TMUX")
	return tmuxEnv != ""
}

func attachToSession(name string) {
	cmd := exec.ArrayCommand{
		Name:    "tmux",
		Command: []string{"a", "-t", name},
	}

	cmd.Run()
}

func switchToSession(name string) {
	cmd := exec.ArrayCommand{
		Name:    "tmux",
		Command: []string{"switch-client", "-t", name},
	}

	cmd.Run()
}
