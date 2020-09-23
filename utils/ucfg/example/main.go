package main

import (
	"github.com/gisvr/defi-common/utils/ucfg"
	"github.com/gisvr/defi-common/utils/ucfg/example/foo"
	"github.com/spf13/viper"
)

// Since package `foo` depends on `bar`, `ucfg` will bootstrap `bar` before `foo`.
func main() {

	// config file setup
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// add current working directory and load config file
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	ucfg.Bootstrap()

	foo.Work()
}
