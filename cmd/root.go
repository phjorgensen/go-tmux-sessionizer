package cmd

import (
	"fmt"
	"log"
	"os"

	"go-tmux-sessionizer/config"
	"go-tmux-sessionizer/fzf"
	"go-tmux-sessionizer/path"
	"go-tmux-sessionizer/tmux"

	"github.com/spf13/cobra"
)

var userLicense string

var rootCmd = &cobra.Command{
	Use:   "go-tmux-sessionizer",
	Short: "CLI utility written in Go that makes it easy to handle tmux sessions.",
	Long:  "This is a CLI utility written in Go that makes it easy to handle tmux sessions. Created by phjorgensen and inspired by ThePrimeagen.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a directory to use for a session",
	Long:  "This will search in the provided paths. If no paths are provided, it will use the paths in the config. It will use the defaults if no config is defined.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println(args)

		f := fzf.Fzf{
			Paths: config.GetPaths(),
		}

		path, err := f.SelectPath()
		if err != nil {
			log.Fatal(err)
		}

		t := tmux.Session{
			Name: path.FormatName(),
			Path: path.Path,
		}

		t.Connect()
	},
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a specific session.",
	Long:  "Open a specific session.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Too few arguments")
		}

		if len(args) > 1 {
			log.Fatal("Too many arguments")
		}

		firstArg := args[0]
		if firstArg == "" {
			log.Fatal("No path passed")
		}

		p := path.SelectedPath{
			Path: firstArg,
		}
		fmt.Println(p)

		t := tmux.Session{
			Name: p.FormatName(),
			Path: p.Path,
		}
		fmt.Println(t)

		t.Connect()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(openCmd)
}
