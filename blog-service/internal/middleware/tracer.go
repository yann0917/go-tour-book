package middleware

import (
	"context"

	"github.com/uber/jaeger-client-go"

	"github.com/opentracing/opentracing-go/ext"

	"github.com/yann0917/go-tour-book/blog-service/global"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {

		var newCtx context.Context
		var span opentracing.Span

		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier{},
		)

		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{
					Key:   string(ext.Component),
					Value: "HTTP",
				},
			)
		}
		defer span.Finish()

		var traceID string
		var SpanID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerCtx := spanContext.(jaeger.SpanContext)
			traceID = jaegerCtx.TraceID().String()
			SpanID = jaegerCtx.SpanID().String()
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", SpanID)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}

}
