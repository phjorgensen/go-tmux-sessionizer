package path

import (
	"errors"
	"path/filepath"
	"strings"
)

type SelectedPath struct {
	name string
	path string

	nameSetManually bool
}

func (s *SelectedPath) SetName(name string) {
	if name == "" {
		s.name = formatName(s.path)
	} else {
		s.name = name
		s.nameSetManually = true
	}
}

func (s SelectedPath) GetName() (string, error) {
	if s.name == "" {
		return "", errors.New("No name is set")
	} else {
		return s.name, nil
	}
}

func (s *SelectedPath) SetPath(path string) error {
	if path == "" {
		return errors.New("Path was not provided.")
	}

	s.path = path

	if !s.nameSetManually {
		s.SetName(formatName(path))
	}

	return nil
}

func (s SelectedPath) GetPath() string {
	return s.path
}

func CreateSelectedPath(name string, path string) SelectedPath {
	p := SelectedPath{}
	p.SetName(name)
	p.SetPath(path)
	return p
}

func formatName(path string) string {
	name := filepath.Base(path)
	return strings.Replace(name, ".", "_", -1)
}
