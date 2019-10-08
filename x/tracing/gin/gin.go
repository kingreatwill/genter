package gin

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

// opentracing中间件;
func GinOpenTracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" || strings.HasSuffix(c.Request.URL.Path, "/healthcheck") {
			c.Next()
			return
		}
		sp := opentracing.SpanFromContext(c.Request.Context())
		var span opentracing.Span
		if sp != nil {
			span = StartSpanWithParent(sp.Context(), c.Request.Method, c.Request.URL.Path)
		} else {
			span = StartSpanWithHeader(&c.Request.Header, c.Request.Method, c.Request.URL.Path)
		}
		span.SetTag("current-goroutines", runtime.NumGoroutine())
		defer span.Finish()
		span.SetTag(string(ext.HTTPStatusCode), c.Writer.Status())
		data, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		span.LogFields(log.Object("http.header", c.Request.Header), log.Object("http.query", c.Request.URL.RawQuery), log.Object("http.raw", string(data)))
		c.Request = c.Request.WithContext(opentracing.ContextWithSpan(c.Request.Context(), span))
		// 记录span, 方便追查;
		c.Header("Span", fmt.Sprintf("%+v", span))
		c.Next()
	}
}

// StartSpanWithParent will start a new span with a parent span;
func StartSpanWithParent(parent opentracing.SpanContext, method, path string) opentracing.Span {
	options := []opentracing.StartSpanOption{
		opentracing.Tag{Key: ext.SpanKindRPCServer.Key, Value: ext.SpanKindRPCServer.Value},
		opentracing.Tag{Key: string(ext.HTTPMethod), Value: method},
		opentracing.Tag{Key: string(ext.HTTPUrl), Value: path},
	}

	if parent != nil {
		options = append(options, opentracing.ChildOf(parent))
	}
	operationName := fmt.Sprintf("%s %s", method, path)
	return opentracing.StartSpan(operationName, options...)
}

// rest client;
func StartSpanWithHeader(header *http.Header, method, path string) opentracing.Span {
	var wireContext opentracing.SpanContext
	if header != nil {
		wireContext, _ = opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(*header))
	}
	span := StartSpanWithParent(wireContext, method, path)
	return span
}
