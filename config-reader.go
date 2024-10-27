package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configSession struct {
	name string
	path string
}

type config struct {
	Paths           []string
	StartupSessions []configSession
}

func getConfig() config {
	found, conf := getPathsFromArgs()
	if found {
		return conf
	}

	found, conf = getPathsFromConfig()
	if found {
		return conf
	}

	return getDefaultConfig()
}

func getPathsFromArgs() (bool, config) {
	args := os.Args[1:]

	if len(args) == 0 {
		return false, config{}
	}

	return true, config{
		Paths: args,
	}
}

func getPathsFromConfig() (bool, config) {
	conf := readConfigFile()

	if len(conf.Paths) == 0 {
		return false, config{}
	}

	return true, config{
		Paths: conf.Paths,
	}
}

func readConfigFile() config {
	data, err := os.ReadFile("/home/phj/.config/tmux-sessionizer/config.json")
	if err != nil {
		fmt.Println("Could not find config file at ~/.config/tmux-sessionizer/config.json. Using default config.")
		return getDefaultConfig()
	}

	var conf config

	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println("Could not parse config file at ~/.config/tmux-sessionizer/config.json. Using default config.")
		return getDefaultConfig()
	}

	return conf
}

func getDefaultConfig() config {
	return config{
		Paths: []string{},
	}
}
