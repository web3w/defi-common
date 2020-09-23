package redis

import (
	"git.bibox.com/dextop/common.git/utils/ucfg"
	"git.bibox.com/dextop/common.git/utils/ulog"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/jsonpb"
	"github.com/spf13/viper"
)

var (
	redisCli        *redis.Client
	JsonPbMarshaler = jsonpb.Marshaler{EmitDefaults: true}
)

type redisConfig struct {
	Addr     string `mapstructure:"endpoint"`
	Password string `mapstructure:"password"`
}

func init() {
	ulog.Infoln("Redis register to get config.")
	ucfg.Register("redis", initRedisClient)
}

func initRedisClient(vp *viper.Viper) {
	if vp == nil {
		panic("Invalid DfeStore config")
	}

	var cfg redisConfig
	err := vp.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	redisCli = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
	})
	_, err = redisCli.Ping().Result()
	if err != nil {
		ulog.Panicln("Redis init fail!", cfg.Addr, cfg.Password)
	}
	ulog.Infoln("Redis init success")
}

/**
 * 将数据存入到缓存中。
 */
func RedisSet(k, v string) bool {
	_, err := redisCli.Set(k, v, 0).Result()
	if err != nil {
		ulog.Error(err)
		return false
	}
	return true
}

func RedisDel(k string) bool {
	_, err := redisCli.Del(k).Result()
	if err != nil {
		ulog.Error(err)
		return false
	}
	return true
}

func RedisLPush(k, v string) bool {
	_, err := redisCli.LPush(k, v).Result()
	if err != nil {
		ulog.Error(err)
		return false
	}
	return true
}

func RedisLTrim(k string, start, stop int64) bool {
	_, err := redisCli.LTrim(k, start, stop).Result()
	if err != nil {
		ulog.Error(err)
		return false
	}
	return true
}
