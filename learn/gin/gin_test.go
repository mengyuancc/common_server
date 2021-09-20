package gin_test

import (
	"testing"
	"time"
)

func TestGin(t *testing.T) {
	s := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	// 定义middleware 所有的请求都会走middleware
	// gin 的context 实现 go的context接口
	s.Use(func(context *gin.Context) {
		// logger.Info("incoming request", zap.String("path", context.Request.URL.Path))
		s := time.Now()
		context.Next()
		logger.Info("incoming request", zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))
	}, func(context *gin.Context) {
		// 生成请求的唯一id
		context.Set("requestId", rand.Int())
	})

	s.GET("/ping", func(context *gin.Context) {
		if rid, ext := context.Get("requestId"); !ext {
			context.JSON(200, gin.H{
				"message":"pong",
			})
		} else {
			context.JSON(200, gin.H{
				"message":"pong",
				"request_id": rid,
			})
		}
	})
	s.Run()
}