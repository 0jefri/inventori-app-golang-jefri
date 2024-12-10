package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inventori-app-jeff/config"
	"github.com/sirupsen/logrus"
)

type RequestLog struct {
	StartTime  time.Time
	EndTime    time.Duration
	StatusCode int
	ClientIP   string
	Method     string
	Path       string
	UserAgent  string
}

func LogRequestMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	file, err := os.OpenFile(config.Cfg.File.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logger.SetOutput(file)

	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		requestLog := RequestLog{
			StartTime:  time.Now(),
			EndTime:    time.Since(time.Now()),
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			UserAgent:  c.Request.UserAgent(),
		}

		switch {
		case c.Writer.Status() >= 500:
			logger.Error(requestLog)
		case c.Writer.Status() >= 400:
			logger.Warn(requestLog)
		default:
			logger.Info(requestLog)
		}
	}
}
