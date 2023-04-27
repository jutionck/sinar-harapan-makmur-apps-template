// Package middleware -> delivery/middleware
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

func LogRequestMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(cfg.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// set output file
	logger.SetOutput(file)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		latency := time.Now().Sub(startTime)
		requestLog := model.RequestLog{
			Latency:      latency,
			StatusCode:   c.Writer.Status(), // statusCode
			ClientIP:     c.ClientIP(),
			Method:       c.Request.Method,
			RelativePath: c.Request.URL.Path,
			UserAgent:    c.Request.UserAgent(),
		}
		switch {
		case c.Writer.Status() >= 500:
			logger.Error(requestLog)
		case c.Writer.Status() >= 400:
			logger.Warn(requestLog)
		default:
			logger.Info(requestLog) // >= 100 .. 200
		}
	}
}
