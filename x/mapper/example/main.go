package main

import (
	"fmt"
	"github.com/openjw/genter/x/mapper"
	time2 "github.com/openjw/genter/x/time"
	"time"
)

type (
	User struct {
		Name     string
		Age      int
		Id       string `mapper:"_id"`
		AA       string `json:"Score,omitempty"`
		Data     []byte
		Students []Student
		Time     time.Time
	}

	Student struct {
		Name  string
		Age   int
		Id    string `mapper:"_id"`
		Score string
	}

	Teacher struct {
		Name  string
		Age   int
		Id    string `mapper:"_id"`
		Level string
	}

	JsonUser struct {
		Name string
		Age  int
		Time time2.Time
	}
)

func init() {
	mapper.Register(&User{})
	mapper.Register(&Student{})
}

func main() {
	user := &User{}
	userMap := &User{}
	teacher := &Teacher{}
	student := &Student{Name: "test", Age: 10, Id: "testId", Score: "100"}
	valMap := make(map[string]interface{})
	valMap["Name"] = "map"
	valMap["Age"] = 10
	valMap["_id"] = "x1asd"
	valMap["Score"] = 100
	valMap["Data"] = []byte{1, 2, 3, 4}
	valMap["Students"] = []byte{1, 2, 3, 4} //[]Student{*student}
	valMap["Time"] = time.Now()

	mapper.SetEnabledTypeChecking(true)

	mapper.Mapper(student, user)
	mapper.AutoMapper(student, teacher)
	mapper.MapperMap(valMap, userMap)

	fmt.Println("student:", student)
	fmt.Println("user:", user)
	fmt.Println("teacher", teacher)
	fmt.Println("userMap:", userMap)

	jsonUser := &JsonUser{
		Name: "json",
		Age:  1,
		Time: time2.Now(),
	}

	fmt.Println(jsonUser)
}
