// Package ulog provides fast, structured, leveled logging.
// ulog depends on ucfg for initialization.
//
// It is easy to create a development logger which only writes to `stdout`, or `stderr`.
// 		logger := ulog.NewLogger()
//
// Config file should be loaded via `viper` before creating a production logger.
//
// example `config.yml`:
// 		---
// 		ulog:
//			defaultLevel: "info"
//			console: true
//			name: myApp
//			rotation:
//			  maxSize: 32 # 32Mb
//
//    viper.SetConfigName("config")
//    viper.SetConfigType("yaml")
//    viper.ReadInConfig()
// 	  ucfg.Bootstrap()
//    logger := ulog.NewLogger()
//
// In order to have custom log level, a new logger can created as follows:
//    txLogger := ulog.NewLoggerWithLevel("debug")
//
// If no config file with proper `ulog` options is given, a development logger will be created
// and messages will be written to `stdout`, `stderr`
//
// check out `example` dir for complete code.
package ulog
