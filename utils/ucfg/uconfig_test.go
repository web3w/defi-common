package ucfg

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestDoubleRegister(t *testing.T) {
	fn := func(vp *viper.Viper) { fmt.Println(vp) }

	doubleRegister := func() {
		Register("key1", fn)
		// double `Register` should induce a panic
		Register("key1", fn)
	}

	assert.PanicsWithValue(t, fmt.Sprintf("%s is already registered", "key1"), doubleRegister)
}

func TestBootstrap(t *testing.T) {
	var list []string
	fn2 := func(vp *viper.Viper) { list = append(list, "fn2") }
	fn3 := func(vp *viper.Viper) { list = append(list, "fn3") }
	fn4 := func(vp *viper.Viper) { list = append(list, "fn4") }

	Register("key2", fn2)
	Register("key3", fn3)
	Register("key4", fn4)
	Bootstrap()

	// verify that the callbacks are executed in order
	assert.Equal(t, []string{"fn2", "fn3", "fn4"}, list)

	// double invocation of `Bootstrap` should induce a panic
	assert.PanicsWithValue(t, "Cannot run callbacks after closed!", func() { Bootstrap() })

	// `Register` when registry is closed should induce a panic
	assert.PanicsWithValue(t, "Cannot register more callbacks after closed!",
		func() { Register("key99", nil) })
}
