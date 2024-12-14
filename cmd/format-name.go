package cmd

import (
	"path/filepath"
	"strings"
)

func formatName(path string) string {
	name := filepath.Base(path)
	return strings.Replace(name, ".", "_", -1)
}
