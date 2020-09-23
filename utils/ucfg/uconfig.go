/*
Each package can register a dedicated section of viper config (with an optional callback), usually
in its init() function. For example, a package foo.bar can register its config namespace "foo.bar"
by `ucfg.Register("foo.bar", func(vp *viper.Viper){})`.

Call ucfg.Bootstrap() to propagate the config sections to the registered callbacks when the viper
config is ready (typically after config file has been loaded at the beginning of the main function).

See ucfg/example for a complete example.
*/

package ucfg

import (
	"sync"

	"github.com/spf13/viper"
)

// CallbackFn takes a viper subtree. A callback should panic on error.
type CallbackFn func(*viper.Viper)

// Registry of config callbacks.
type registry struct {
	sync.Mutex // guards the whole structure

	callbacks map[string]CallbackFn
	initQueue []string
	closed    bool
}

func newRegistry() *registry {
	return &registry{
		callbacks: make(map[string]CallbackFn),
		initQueue: nil,
		closed:    false,
	}
}

var reg = newRegistry() // singleton

// Registers a callback for a config namespace.
// - Each package having a dedicated config namespace should call this method once.
// - It will panic if the namespace is already registered.
// - `callback` can be nil, which means no callback for that namespace.
func Register(namespace string, callback CallbackFn) {
	reg.registerCallback(namespace, callback)
}

// Propagates the configs to the registered packages by calling the callbacks. This function is
// supposed to be called exactly once when the viper config is loaded.
//
// The callbacks are called in the order they are registered one by one in a single goroutine.
func Bootstrap() {
	reg.runCallbacks()
}

// Registers a callback. Panics if the key is already registered.
func (cfg *registry) registerCallback(key string, callback CallbackFn) {
	cfg.Lock()
	defer cfg.Unlock()

	if cfg.closed {
		panic("Cannot register more callbacks after closed!")
	}

	if _, ok := cfg.callbacks[key]; ok {
		panic(key + " is already registered")
	}

	cfg.callbacks[key] = callback
	cfg.initQueue = append(cfg.initQueue, key)
}

// Calls the registered callbacks in the order they are registered one by one in a single goroutine.
func (cfg *registry) runCallbacks() {
	cfg.Lock()
	defer cfg.Unlock()

	if cfg.closed {
		panic("Cannot run callbacks after closed!")
	}
	cfg.closed = true

	for _, key := range cfg.initQueue {
		callback := cfg.callbacks[key]
		if callback != nil {
			sub := viper.Sub(key)
			callback(sub)
		}
	}
}
