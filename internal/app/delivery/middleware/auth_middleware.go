package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/inventori-app-jeff/internal/model/dto"
	"github.com/inventori-app-jeff/utils/exception"
	"github.com/inventori-app-jeff/utils/security"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

// Middleware sederhana untuk otentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header authHeader
		// userOTP := c.PostForm("otp")

		if err := c.ShouldBindHeader(&header); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  exception.StatusUnauthorized,
				Message: exception.ErrTokenNotProvided.Error(),
			})
			return
		}

		if header.AuthorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  exception.StatusUnauthorized,
				Message: exception.ErrTokenNotProvided.Error(),
			})
			return
		}

		token := strings.Split(header.AuthorizationHeader, " ")[1]

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  exception.StatusUnauthorized,
				Message: exception.ErrTokenNotProvided.Error(),
			})
			return
		}

		claims, err := security.VerifyAccessToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  exception.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		// Simpan data pengguna dalam konteks
		c.Set("username", claims["username"])

		c.Next()

	}
}

func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		} else {
			c.Set("user", input)
			c.Next()
		}
	}
}
