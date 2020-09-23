package ulog

import (
	"github.com/gisvr/defi-common/utils/ucfg"
	"github.com/spf13/viper"
)

// By default only log to console.
var pkgCfg = pkgConfig{
	Console:      true,
	DefaultLevel: "info",
}

var defaultLogger *ulogger

func init() {
	defaultLogger = &ulogger{
		log:   newProductionLogger(pkgCfg.DefaultLevel),
		level: pkgCfg.DefaultLevel,
	}
	ucfg.Register("ulog", configureDefault)
}

func configureDefault(vp *viper.Viper) {
	// config file may not include `ulog` key
	// in case, user may be OK with sole `console` logger.
	if vp == nil {
		return
	}

	var cfg pkgConfig

	if err := vp.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	if err := cfg.Validate(); err != nil {
		panic(err)
	}
	pkgCfg = cfg

	defaultLogger = &ulogger{
		log:   newProductionLogger(cfg.DefaultLevel),
		level: pkgCfg.DefaultLevel,
	}
}
