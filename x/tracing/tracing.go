package tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tracinglog "github.com/opentracing/opentracing-go/log"
)

func ErrorLog(ctx context.Context, err error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		ext.Error.Set(span, true)
		span.LogFields(tracinglog.String("errorlog", err.Error()))
	} else {
		span, _ := opentracing.StartSpanFromContext(ctx, "ErrorLog")
		ext.Error.Set(span, true)
		span.LogFields(tracinglog.String("errorlog", err.Error()))
		defer span.Finish()
	}
}

func NewContext(ctx context.Context) context.Context {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		return opentracing.ContextWithSpan(context.Background(), span)
	}
	return context.Background()
}

func StartSpan(ctx context.Context, name string, args ...interface{}) opentracing.Span {
	span, _ := opentracing.StartSpanFromContext(ctx, name)
	span.LogFields(tracinglog.Object("args", args))
	return span
}

func TracingFunc(ctx context.Context, name string, fn func()) {
	sp := StartSpan(ctx, "fn", name)
	defer sp.Finish()
	fn()
}

func Log(ctx context.Context, key string, value interface{}) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.LogFields(tracinglog.Object(key, value))
	}
	context.Background()
}

func LogDebug(ctx context.Context, key string, value interface{}) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.LogFields(tracinglog.Object(key, value))
	}
	context.Background()
}

func SetTag(ctx context.Context, key string, value interface{}) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag(key, value)
	}
}
