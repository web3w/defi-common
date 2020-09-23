package dexinit

import (
	"bytes"
	"github.com/gisvr/deif-common/utils/ucfg"
	"github.com/gisvr/deif-common/utils/ulog"
	"io/ioutil"
	"os"
	"path"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	envFlag       = flag.String("env", "", "testing|local|devnet|testnet|mainnet")
	configDirFlag = flag.String("config_dir", "",
		"Path to directory containing the config file. By default the prefix of cmdline arg0 is used.")
	dexcfgFileFlag = flag.String("dexcfg_file", "",
		"Path to the DEx global config file. By default '<config_dir>/dexcfg_<env>.yml' will be used.")

	log = ulog.NewLogger()
)

// Reads in config files and calls `ucfg.Bootstrap()`. To be called at the beginning of the main
// function.
//
// Assumptions:
// - Package "github.com/spf13/pflag" is used for defining cmdline flags.
//
// Examples:
//     # All by default, use ./mybin/dexcfg_mainnet.yml and ./mybin/abc_config_mainnet.yml
//     $ ./mybin/abc --env=mainnet
//
//     # Use /opt/dex/dexcfg_local.yml and ../cfg/abc_config_local_yml
//     $ ./mybin/abc --env=local --config_dir=../cfg --dexcfg_file=/opt/dex/dexcfg_local.yml
//
func DexReadConfigAndBootstrap() {
	flag.Parse()
	viper.BindPFlags(flag.CommandLine)

	checkFlags()

	if *configDirFlag == "" {
		*configDirFlag = path.Dir(os.Args[0])
	}

	if *dexcfgFileFlag == "" {
		*dexcfgFileFlag = path.Join(*configDirFlag, "dexcfg_"+*envFlag+".yml")
	}

	dexCfg, err := ioutil.ReadFile(*dexcfgFileFlag)
	if err != nil {
		log.Fatalw("failed to read dexcfg", "path", *dexcfgFileFlag, "err", err)
	}
	dexCfg = append(dexCfg, '\n') // in case dexcfg file does not have a ending line break

	binName := path.Base(os.Args[0])
	configPath := path.Join(*configDirFlag, binName) + "_config_" + *envFlag + ".yml"
	localCfg, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalw("failed to read config file", "path", configPath, "err", err)
	}

	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(append(dexCfg, localCfg...))); err != nil {
		log.Fatal("failed to parse config: ", err)
	}

	// dexcfg.env sanity check
	if *envFlag != viper.GetString("dexcfg.env") {
		log.Fatal("--env flag mismatches dexcfg.env")
	}

	ucfg.Bootstrap()
}

func checkFlags() {
	env := Env(*envFlag)
	if env == "" {
		log.Fatal("missing --env flag")
	}
	if env != Testing && env != Local && env != Devnet && env != Testnet && env != Mainnet {
		log.Fatal("--env flag must be one of testing|local|devnet|testnet|mainnet")
	}
}
