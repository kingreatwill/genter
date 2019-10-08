package rabbitmq

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/streadway/amqp"
)

func ConsumeExtract(msg *amqp.Delivery) context.Context {

	// 提取头部信息.
	spCtx, err := MqExtract(msg.Headers)
	if err != nil {
		span := opentracing.StartSpan(
			"ConsumeMessage",
		)
		defer span.Finish()
		span.LogFields(log.String("mq.body", string(msg.Body)))
		//ext.Error.Set(span, true)
		span.LogFields(log.String("context", "没有上下文."), log.String("error", err.Error()))
		return opentracing.ContextWithSpan(context.Background(), span)
	} else {
		span := opentracing.StartSpan(
			"ConsumeMessage",
			opentracing.ChildOf(spCtx),
		)
		defer span.Finish()
		span.LogFields(log.String("mq.body", string(msg.Body)))
		return opentracing.ContextWithSpan(context.Background(), span)
	}
}

func ProducerInject(ctx context.Context, msg *amqp.Publishing) error {
	if !Enable {
		return nil
	}
	sparent := opentracing.SpanFromContext(ctx)
	if sparent == nil {
		span := opentracing.StartSpan(
			"PublishMessage",
		)
		defer span.Finish()
		//ext.Error.Set(span, true)
		span.LogFields(log.String("context", "没有上下文."))
		msg.Headers = make(map[string]interface{})
		// 注入头部信息;
		if err := MqInject(span, msg.Headers); err != nil {
			ext.Error.Set(span, true)
			span.LogFields(log.String("error", err.Error()))
			return err
		}
	} else {
		span := opentracing.StartSpan(
			"PublishMessage",
			opentracing.FollowsFrom(sparent.Context()),
		)
		defer span.Finish()
		msg.Headers = make(map[string]interface{})
		// 注入头部信息;
		if err := MqInject(span, msg.Headers); err != nil {
			ext.Error.Set(span, true)
			span.LogFields(log.String("error", err.Error()))
			return err
		}
	}
	return nil
}

type amqpHeadersCarrier map[string]interface{}

// ForeachKey conforms to the TextMapReader interface.
func (c amqpHeadersCarrier) ForeachKey(handler func(key, val string) error) error {
	for k, val := range c {
		v, ok := val.(string)
		if !ok {
			continue
		}
		if err := handler(k, v); err != nil {
			return err
		}
	}
	return nil
}

// Set implements Set() of opentracing.TextMapWriter.
func (c amqpHeadersCarrier) Set(key, val string) {
	c[key] = val
}

func MqInject(span opentracing.Span, hdrs amqp.Table) error {
	c := amqpHeadersCarrier(hdrs)
	return span.Tracer().Inject(span.Context(), opentracing.TextMap, c)
}

// Extract extracts the span context out of the AMQP header.
func MqExtract(hdrs amqp.Table) (opentracing.SpanContext, error) {
	c := amqpHeadersCarrier(hdrs)
	return opentracing.GlobalTracer().Extract(opentracing.TextMap, c)
}
