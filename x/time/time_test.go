package time

import (
	"github.com/openjw/genter/x/json"
	"testing"
	"time"
)

type Person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Birthday Time   `json:"birthday"`
}

func TestTimeJson(t *testing.T) {
	now := time.Now()
	t.Log(now)
	SetTimeJSONFormat("2006-01-02 15:04:05")
	src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51"}`
	p := new(Person)
	err := json.Unmarshal([]byte(src), p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
	t.Log(p.Birthday)
	js, _ := json.Marshal(p)
	t.Log(string(js))
}

func TestTimeJson_NewFormat(t *testing.T) {
	now := time.Now()
	t.Log(now)
	SetTimeJSONFormat("2006-01-02T15:04:05")
	src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30T16:09:51"}`
	p := new(Person)
	err := json.Unmarshal([]byte(src), p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
	t.Log(p.Birthday)
	js, _ := json.Marshal(p)
	t.Log(string(js))
}
