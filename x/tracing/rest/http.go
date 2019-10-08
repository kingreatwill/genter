package rest

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

// InjectTraceID to Header;
func Inject(ctx opentracing.SpanContext, header http.Header) {
	opentracing.GlobalTracer().Inject(
		ctx,
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header))
}

// InjectTraceID to Header;
func InjectTraceID(ctx context.Context, header http.Header) {
	span := opentracing.SpanFromContext(ctx)
	spanCtx := span.Context()
	Inject(spanCtx, header)
}
