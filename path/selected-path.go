package path

import (
	"path/filepath"
	"strings"
)

type SelectedPath struct {
	Path string
}

func (s SelectedPath) FormatName() string {
	name := filepath.Base(s.Path)
	return strings.Replace(name, ".", "_", -1)
}

func CreatePath(path string) SelectedPath {
	return SelectedPath{
		Path: path,
	}
}
