package main

import (
	"go-tmux-sessionizer/cmd"
	"go-tmux-sessionizer/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
