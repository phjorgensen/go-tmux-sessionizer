package fzf

import (
	"go-tmux-sessionizer/exec"
	"strings"
)

func Open(paths []string) (string, error) {
	cmd := exec.StringCommand{
		Command: "find " + strings.Join(paths, " ") + " -mindepth 1 -maxdepth 1 -type d | fzf",
	}

	path, err := cmd.Out()
	if err != nil {
		return "", err
	}

	return path, nil
}
