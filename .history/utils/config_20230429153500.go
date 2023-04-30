package utils

import ()


type Config struct {
	DBDriver string  `mapstructure:"DB_DRIVER"`
	DBSource string string `mapstructure:"DB_SOURCE"`
	ServerAddress string string `mapstructure:"SERVER_ADDRESS"`

}
