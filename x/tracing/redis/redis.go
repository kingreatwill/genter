package redis

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

func RedigoDo(ctx context.Context, stats interface{}, fn func(commandName string, args ...interface{}) (interface{}, error), commandName string, args ...interface{}) (ctx2 context.Context, reply interface{}, err error) {
	span, ctx1 := opentracing.StartSpanFromContext(ctx, "redis-"+commandName)
	defer span.Finish()
	ext.PeerService.Set(span, "redis")
	ext.DBType.Set(span, "redis")
	span.SetTag("db.method", commandName)
	if commandName == "set" {
		span.LogFields(log.Object("redis.args", args[0]))
	} else if commandName == "hset" {
		span.LogFields(log.Object("redis.args", args[0:1]))
	} else {
		span.LogFields(log.Object("redis.args", args))
	}
	span.LogFields(log.Object("redis.pool.stats", stats))
	reply, err = fn(commandName, args...)
	if err != nil {
		ext.Error.Set(span, true)
		span.LogFields(log.String("error", err.Error()))
	}
	return ctx1, reply, err
}
