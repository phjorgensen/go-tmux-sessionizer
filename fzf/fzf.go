package fzf

import (
	"go-tmux-sessionizer/exec"
	"go-tmux-sessionizer/path"
	"strings"
)

type Fzf struct {
	Paths []string
}

func (fzf *Fzf) AddPath(path string) {
	fzf.Paths = append(fzf.Paths, path)
}

func (fzf *Fzf) SelectPath() (path.SelectedPath, error) {
	cmd := exec.StringCommand{
		Command: "find " + strings.Join(fzf.Paths, " ") + " -mindepth 1 -maxdepth 1 -type d | fzf",
	}

	res, err := cmd.Out()
	if err != nil {
		return path.SelectedPath{}, err
	}

	return path.CreatePath(res), nil
}
