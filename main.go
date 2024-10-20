package main

func main() {
	f := fzf{}

	f.addSearchPath("~/")
	f.addSearchPath("~/Projects")
	f.addSearchPath("~/Documents")
	f.addSearchPath("~/Documents/notes")

	f.selectPath()

	t := tmux{
		sessionName: f.selectedPath.formatName(),
		sessionPath: f.selectedPath.path,
	}

	t.connect()
}
