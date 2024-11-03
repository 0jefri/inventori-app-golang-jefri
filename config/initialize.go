package config

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var (
	Cfg *Config
)

func InitiliazeConfig() {

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(err)
	}

	appTokenExpire, _ := strconv.Atoi(Cfg.Token.Expire)

	// if err != nil {
	// 	panic(err)
	// }

	Cfg.TokenConfig = TokenConfig{
		ApplicationName:     Cfg.Token.Name,
		JWTSignatureKey:     []byte(Cfg.Token.Key),
		JWTSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: time.Duration(appTokenExpire) * time.Hour,
	}

}
