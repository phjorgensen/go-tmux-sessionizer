package cmd

import (
	"fmt"
	"log"
	"os"

	"go-tmux-sessionizer/config"
	"go-tmux-sessionizer/fzf"
	"go-tmux-sessionizer/tmux"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(configPathCmd)
}

var rootCmd = &cobra.Command{
	Use:   "go-tmux-sessionizer",
	Short: "CLI utility written in Go that makes it easy to handle tmux sessions",
	Long:  "This is a CLI utility written in Go that makes it easy to handle tmux sessions. Created by phjorgensen and inspired by ThePrimeagen.",
	Run: func(cmd *cobra.Command, args []string) {
		searchCmd.Run(cmd, args)
	},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a directory to use for a session",
	Long:  "This will search in the provided paths. If no paths are provided, it will use the paths in the config. It will use the defaults if no config is defined.",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := fzf.Open(config.GetPaths())
		if err != nil {
			log.Fatal(err)
		}

		name := formatName(path)

		if !tmux.HasSession(name) {
			tmux.CreateSession(name, path)
		}

		tmux.Connect(name)
	},
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a specific session",
	Long:  "Open a specific session.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Too few arguments")
		}

		if len(args) > 1 {
			log.Fatal("Too many arguments")
		}

		path := args[0]
		if path == "" {
			log.Fatal("No path passed")
		}

		name := formatName(path)

		if !tmux.HasSession(name) {
			tmux.CreateSession(name, path)
		}

		tmux.Connect(name)
	},
}

var configPathCmd = &cobra.Command{
	Use:   "get-config",
	Short: "Get the path to the config file",
	Long:  "This will list out the path to the config that is used.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.ConfigFileUsed())
	},
}
