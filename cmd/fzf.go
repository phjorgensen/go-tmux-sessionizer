package cmd

import (
	"go-tmux-sessionizer/exec"
	"go-tmux-sessionizer/path"
	"strings"
)

func openFzf(paths []string) (path.SelectedPath, error) {
	cmd := exec.StringCommand{
		Command: "find " + strings.Join(paths, " ") + " -mindepth 1 -maxdepth 1 -type d | fzf",
	}

	selectedPath, err := cmd.Out()
	if err != nil {
		return path.SelectedPath{}, err
	}

	return path.CreateSelectedPath("", selectedPath), nil
}
