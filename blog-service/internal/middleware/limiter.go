package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yann0917/go-tour-book/blog-service/pkg/app"
	"github.com/yann0917/go-tour-book/blog-service/pkg/errcode"
	"github.com/yann0917/go-tour-book/blog-service/pkg/limiter"
)

func RateLimiter(l limiter.IFace) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)

			if count == 0 {
				resp := app.NewResp(c)
				resp.ToErrorResp(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
