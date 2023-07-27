

## http client
如果要追踪各个时间段的请求, 可以使用httptrace
[httpstat](https://github.com/davecheney/httpstat)
`httpstat https://example.com/ `
可以看到 `DNS Lookup`, `TCP Connection`, `TLS Handshake`(http没有, 其它的顺序一样), `Server Processing`, `Content Transfer`

原理就是利用httptrace.ClientTrace来做各个阶段的监控

> [go-httpstat](https://github.com/tcnksm/go-httpstat) 类似


``` golang
import (
   "net/http"
   "net/http/httptrace"

   "github.com/opentracing/opentracing-go"
   "github.com/opentracing/opentracing-go/log"
   "golang.org/x/net/context"
)

// We will talk about this later
var tracer opentracing.Tracer

func AskGoogle(ctx context.Context) error {
   // retrieve current Span from Context
   var parentCtx opentracing.SpanContext
   parentSpan := opentracing.SpanFromContext(ctx); 
   if parentSpan != nil {
      parentCtx = parentSpan.Context()
   }

   // start a new Span to wrap HTTP request
   span := tracer.StartSpan(
      "ask google",
      opentracing.ChildOf(parentCtx),
   )

   // make sure the Span is finished once we're done
   defer span.Finish()

   // make the Span current in the context
   ctx = opentracing.ContextWithSpan(ctx, span)

   // now prepare the request
   req, err := http.NewRequest("GET", "http://google.com", nil)
   if err != nil {
      return err
   }

   // attach ClientTrace to the Context, and Context to request 
   trace := NewClientTrace(span)
   ctx = httptrace.WithClientTrace(ctx, trace)
   req = req.WithContext(ctx)

   // execute the request
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   
   // Google home page is not too exciting, so ignore the result
   res.Body.Close()
   return nil
}




func NewClientTrace(span opentracing.Span) *httptrace.ClientTrace {
   trace := &clientTrace{span: span}
   return &httptrace.ClientTrace {
      DNSStart: trace.dnsStart,
      DNSDone:  trace.dnsDone,
   }
}

// clientTrace holds a reference to the Span and
// provides methods used as ClientTrace callbacks
type clientTrace struct {
   span opentracing.Span
}

func (h *clientTrace) dnsStart(info httptrace.DNSStartInfo) {
   h.span.LogKV(
      log.String("event", "DNS start"),
      log.Object("host", info.Host),
   )
}

func (h *clientTrace) dnsDone(httptrace.DNSDoneInfo) {
   h.span.LogKV(log.String("event", "DNS done"))
}
```


使用 https://github.com/opentracing-contrib/go-stdlib
```golang
package main

import (
   "fmt"
   "io/ioutil"
   "log"
   "net/http"

   "github.com/opentracing-contrib/go-stdlib/nethttp"
   "github.com/opentracing/opentracing-go"
   "github.com/opentracing/opentracing-go/ext"
   otlog "github.com/opentracing/opentracing-go/log"
   "golang.org/x/net/context"
)

func runClient(tracer opentracing.Tracer) {
   // nethttp.Transport from go-stdlib will do the tracing
   c := &http.Client{Transport: &nethttp.Transport{}}

   // create a top-level span to represent full work of the client
   span := tracer.StartSpan(client)
   span.SetTag(string(ext.Component), client)
   defer span.Finish()
   ctx := opentracing.ContextWithSpan(context.Background(), span)

   req, err := http.NewRequest(
      "GET",
      fmt.Sprintf("http://localhost:%s/", *serverPort),
      nil,
   )
   if err != nil {
      onError(span, err)
      return
   }

   req = req.WithContext(ctx)
   // wrap the request in nethttp.TraceRequest
   req, ht := nethttp.TraceRequest(tracer, req)
   defer ht.Finish()

   res, err := c.Do(req)
   if err != nil {
      onError(span, err)
      return
   }
   defer res.Body.Close()
   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      onError(span, err)
      return
   }
   fmt.Printf("Received result: %s\n", string(body))
}

func onError(span opentracing.Span, err error) {
   // handle errors by recording them in the span
   span.SetTag(string(ext.Error), true)
   span.LogKV(otlog.Error(err))
   log.Print(err)
}
```