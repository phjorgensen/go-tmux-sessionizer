package main

import (
	"path/filepath"
	"strings"
)

type selectedPath struct {
	path string
}

func (s selectedPath) formatName() string {
	name := filepath.Base(s.path)
	return strings.Replace(name, ".", "_", -1)
}

func createPath(path string) selectedPath {
	return selectedPath{
		path: path,
	}
}
