package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Server Config
type Server struct {
	Port int `mapstructure:"port"`
}

// Database Config
type Database struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
	Driver   string `mapstructure:"driver"`
}

// File Config
type File struct {
	Path            string `mapstructure:"path"`
	UserPicturePath string `mapstructure:"picture_path"`
}

// Token Config
type Token struct {
	Name   string `mapstructure:"name"`
	Key    string `mapstructure:"key"`
	Expire string `mapstructure:"expire"`
}

type TokenConfig struct {
	ApplicationName     string
	JWTSignatureKey     []byte
	JWTSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

// Configuration
type Config struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
	File     `mapstructure:"file"`
	Token    `mapstructure:"token"`
	TokenConfig
	DefaultRowsPerPage string `mapstructure:"DEFAULT_ROWS_PER_PAGE"`
}
