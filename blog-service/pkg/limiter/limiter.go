package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type IFace interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBucket(rules ...BucketRule) IFace
}

type Limiter struct {
	buckets map[string]*ratelimit.Bucket
}

type BucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
