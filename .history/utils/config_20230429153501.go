package utils

import ()


type Config struct {
	DBDriver string  `mapstructure:"DB_DRIVER"`
	DBSource string  `mapstructure:"DB_SOURCE"`
	ServerAddress string string `mapstructure:"SERVER_ADDRESS"`

}
