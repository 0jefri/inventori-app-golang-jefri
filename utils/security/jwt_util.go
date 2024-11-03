package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/inventori-app-jeff/config"
	"github.com/inventori-app-jeff/internal/model"
	"github.com/inventori-app-jeff/utils/exception"
)

func CreateAccessToken(user *model.User) (string, error) {
	now := time.Now().UTC()
	end := now.Add(config.Cfg.TokenConfig.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Cfg.TokenConfig.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(config.Cfg.TokenConfig.JWTSigningMethod, claims)
	ss, err := token.SignedString(config.Cfg.TokenConfig.JWTSignatureKey)

	if err != nil {
		return "", exception.ErrFailedCreateToken
	}

	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != config.Cfg.TokenConfig.JWTSigningMethod {
			return nil, exception.ErrInvalidTokenStringMethod
		}
		return config.Cfg.TokenConfig.JWTSignatureKey, nil
	})

	if err != nil {
		return nil, exception.ErrInvalidParseToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid || claims["iss"] != config.Cfg.TokenConfig.ApplicationName {
		return nil, exception.ErrInvalidTokenMapclaims
	}

	return claims, nil
}
