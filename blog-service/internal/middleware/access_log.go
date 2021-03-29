package middleware

import (
	"bytes"
	"time"

	"github.com/yann0917/go-tour-book/blog-service/global"

	"github.com/yann0917/go-tour-book/blog-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Writer(p []byte) (int, error) {

	if n, err := w.Write(p); err != nil {
		return n, err
	}

	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		s := "access log: method: %s, status_code: %d" +
			"begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).Infof(c, s,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime)
	}
}
