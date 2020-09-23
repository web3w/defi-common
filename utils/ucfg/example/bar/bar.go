package bar

import (
	"fmt"
	"git.bibox.com/dextop/common.git/utils/ucfg"
	"github.com/spf13/viper"
)

type Bar struct {
	Addr     string `mapstructure:"addr"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

var bar Bar

func init() {
	ucfg.Register("example.bar", initBar)
}

func initBar(vp *viper.Viper) {
	err := vp.Unmarshal(&bar)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("init Bar done")
	fmt.Printf("%+v\n", bar)
}

func Work() {

}
