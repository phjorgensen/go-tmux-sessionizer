package main

import (
	"strings"
)

type fzf struct {
	paths []string
}

func (fzf *fzf) addPath(path string) {
	fzf.paths = append(fzf.paths, path)
}

func (fzf *fzf) selectPath() (selectedPath, error) {
	cmd := stringCommand{
		command: "find " + strings.Join(fzf.paths, " ") + " -mindepth 1 -maxdepth 1 -type d | fzf",
	}

	path, err := cmd.out()
	if err != nil {
		return selectedPath{}, err
	}

	return createPath(path), nil
}
