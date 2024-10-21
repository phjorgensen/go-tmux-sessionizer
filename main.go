package main

import "log"

func main() {
	f := fzf{
		paths: []string{
			"~/",
			"~/Projects",
			"~/Documents",
			"~/Documents/notes",
		},
	}

	path, err := f.selectPath()
	if err != nil {
		log.Fatal(err)
	}

	t := tmuxSession{
		name: path.formatName(),
		path: path.path,
	}

	t.connect()
}
