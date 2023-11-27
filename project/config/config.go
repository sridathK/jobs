package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig   AppConfig
	DbConfig    DbConfig
	RedisConfig RedisConfig
	Keys        Keys
}

type AppConfig struct {
	Host string `env:"APP_HOST"`
	Port string `env:"APP_PORT,required=true"`
}

type DbConfig struct {
	DbAdress string `env:"DB_DSN,required=true"`
}
type RedisConfig struct {
	Adress   string `env:"Redis_Adress,required=true"`
	Password string `env:"Redis_Password,required=true"`
	DB       string `env:"Redis_DB,required=true"`
}

type Keys struct {
	PrivateKey string `env:"PrivateKey,required=true"`
	PublicKey  string `env:"PublicKey,required=true"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Panic(err)
	}
}

func GetConfig() Config {
	return cfg
}
