package common

/**
 * MySQL配置项对象。
 */
type MySQLConfig struct {
	Endpoint string `mapstructure:"endpoint"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

/**
 * Redis配置项对象。
 */
type RedisConfig struct {
	Endpoint string `mapstructure:"endpoint"`
	Password string `mapstructure:"password"`
}
