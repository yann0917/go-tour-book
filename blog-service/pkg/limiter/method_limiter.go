package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() IFace {
	l := &Limiter{buckets: map[string]*ratelimit.Bucket{}}
	return MethodLimiter{
		Limiter: l,
	}
}

func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

//GetBucket 获取令牌桶
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.buckets[key]
	return bucket, ok
}

//AddBucket 新增多个令牌桶
func (l MethodLimiter) AddBucket(rules ...BucketRule) IFace {
	for _, rule := range rules {
		if _, ok := l.buckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval,
				rule.Capacity,
				rule.Quantum,
			)
			l.buckets[rule.Key] = bucket
		}
	}
	return l
}
