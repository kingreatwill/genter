package json

import (
	"github.com/openjw/genter/x/json/jsoniter"
	"github.com/openjw/genter/x/json/original"
	"testing"
	"time"
)

// 原生性能测试;
func Benchmark_Marshal_original(b *testing.B) {
	j := original.New()
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		_, _ = j.Marshal(testStruct{})
	}
}

// jsoniter性能测试;
func Benchmark_Marshal_jsoniter(b *testing.B) {
	j := jsoniter.New()
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		_, _ = j.Marshal(testStruct{})
	}
}

type testStruct struct {
	Str  string
	Int  int
	Time time.Time
}

// 原生性能测试;
func Benchmark_Unmarshal_original(b *testing.B) {
	j := original.New()
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		_ = j.Unmarshal([]byte(`
{
"Str": "123",
"Int": 1,
"Time": "2019-08-30 11:30:39"
}`), &testStruct{})
	}
}

// jsoniter性能测试;
func Benchmark_Unmarshal_jsoniter(b *testing.B) {
	j := jsoniter.New()
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		_ = j.Unmarshal([]byte(`
{
"Str": "123",
"Int": 1,
"Time": "2019-08-30 11:30:39"
}`), &testStruct{})
	}
}
