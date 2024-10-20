package main

import (
	"strings"
)

type fzf struct {
	searchPaths  []string
	selectedPath selectedPath
}

func (fzf *fzf) addSearchPath(searchPath string) {
	fzf.searchPaths = append(fzf.searchPaths, searchPath)
}

func (fzf *fzf) selectPath() {
	cmd := stringCommand{
		command: "find " + strings.Join(fzf.searchPaths, " ") + " -mindepth 1 -maxdepth 1 -type d | fzf",
	}

	path, err := cmd.out()
	if err != nil {
		return
	}

	fzf.selectedPath = createPath(path)
}
