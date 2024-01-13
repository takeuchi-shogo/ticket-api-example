package config

import (
	"log"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type ServerConfig struct {
	AppEnvironment string `mapstructure:"API_ENV"`
	AppName        string `mapstructure:"APP_NAME"`
	TokenTTL       int    `mapstructure:"TOKEN_TTL"`
	JwtAuthKey     string `mapstructure:"JWT_PRIVATE_KEY"`
	// DB             struct {
	// 	Host     string `mapstructure:"DB_HOST"`
	// 	UserName string `mapstructure:"DB_USER"`
	// 	Password string `mapstructure:"DB_PASS"`
	// 	DBName   string `mapstructure:"DB_NAME"`
	// }
	DBHost          string `mapstructure:"DB_HOST"`
	DBUserName      string `mapstructure:"DB_USER"`
	DBPassword      string `mapstructure:"DB_PASS"`
	DBName          string `mapstructure:"DB_NAME"`
	StripePublicKey string `mapstructure:"STRIPE_PUBLIC_KEY"`
	StripeSecretKey string `mapstructure:"STRIPE_SECRET_KEY"`
	Port            int    `mapstructure:"SERVER_PORT"`
}

var Module = fx.Options(
	fx.Provide(NewConfig),
)

func NewConfig() ServerConfig {
	config := ServerConfig{}
	viper.SetConfigFile("./.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("環境設定ファイルの読み込みに失敗しました")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	return config
}
