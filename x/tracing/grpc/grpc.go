package tracing

import (
	"context"
	"log"
	"strings"

	tracinglog "github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// mdWrite implement TextMapWriter
type mdWriter struct {
	metadata.MD
}

func (m mdWriter) Set(key, val string) {
	lKey := strings.ToLower(key)
	if v, ok := m.MD[lKey]; ok {
		m.MD[lKey] = append(v, val)
	} else {
		m.MD[lKey] = []string{val}
	}
}

// mdReader implement TextMapReader
type mdReader struct {
	metadata.MD
}

func (m mdReader) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range m.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

//OpenTracingClientInterceptor  rewrite client's interceptor with open tracing
func OpenTracingClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, resp interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		var parentCtx opentracing.SpanContext
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			parentCtx = parent.Context()
		}
		cliSpan := tracer.StartSpan(
			"rpc "+method,
			opentracing.ChildOf(parentCtx),
			opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
			ext.SpanKindRPCClient,
		)
		defer cliSpan.Finish()
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}
		err := tracer.Inject(cliSpan.Context(), opentracing.TextMap, mdWriter{md})
		if err != nil {
			log.Panicf("inject to metadata err %v", err)
		}
		//将metadata数据装入context中
		ctx = metadata.NewOutgoingContext(ctx, md)
		err = invoker(ctx, method, req, resp, cc, opts...)
		if err != nil {
			cliSpan.LogFields(tracinglog.Error(err))
		}
		return err
	}
}

//OpentracingServerInterceptor rewrite server's interceptor with open tracing
func OpentracingServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		// 拦截内容
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		spanContext, err := tracer.Extract(opentracing.TextMap, mdReader{md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			log.Panicf("extract from metadata err %v", err)
		}
		srvSpan := tracer.StartSpan(
			"rpc "+info.FullMethod,
			ext.RPCServerOption(spanContext),
			opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
			ext.SpanKindRPCServer,
		)
		defer srvSpan.Finish()
		ctx = opentracing.ContextWithSpan(ctx, srvSpan)
		// 继续处理
		return handler(ctx, req)
	}
}
