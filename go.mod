module github.com/openjw/genter

go 1.12

require (
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/jinzhu/gorm v1.9.11 // indirect
	github.com/json-iterator/go v1.1.7
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271 // indirect
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.19.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/atomic v1.4.0 // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/yaml.v2 v2.2.2
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.24.0
