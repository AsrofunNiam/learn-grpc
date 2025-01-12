package configuration

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	// connection to db
	HttpPort      string `mapstructure:"HTTP_PORT"`
	PortDB        string `mapstructure:"PORT_DB"`
	Host          string `mapstructure:"HOST_DB"`
	Password      string `mapstructure:"PASSWORD_DB"`
	User          string `mapstructure:"USER_DB"`
	Db            string `mapstructure:"DATABASE_DB"`
	EncryptionKey string `mapstructure:"ENCRYPTION_KEY"`

	// connection to redis
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`

	// connection to gRPC
	GRPCPort string `mapstructure:"GRPC_PORT"`
}

func LoadConfig() (config Configuration, err error) {
	viper.SetConfigFile("./configuration/.env")
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
