package main

func main() {
	conf := getConfig()

	f := fzf{
		paths: conf.Paths,
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
