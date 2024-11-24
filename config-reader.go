package main

import (
	"log"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config/tmux-sessionizer")

	viper.SetDefault("paths", []string{"~/"})

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func getPaths() []string {
	return viper.GetStringSlice("paths")
}
