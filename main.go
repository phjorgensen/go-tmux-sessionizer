package main

func main() {
	initConfig()

	f := fzf{
		paths: getPaths(),
	}

	path, err := f.selectPath()
	if err != nil {
		panic(err)
	}

	t := tmuxSession{
		name: path.formatName(),
		path: path.path,
	}

	t.connect()
}
