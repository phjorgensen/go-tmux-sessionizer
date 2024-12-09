package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func Init() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(configDir + "/tmux-sessionizer")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()

	viper.SetDefault("paths", []string{"~/"})

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func GetPaths() []string {
	return viper.GetStringSlice("paths")
}
