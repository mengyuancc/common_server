package interface_t_test

import (
	"fmt"
	"iris_server/common"
	"iris_server/learn/interface/infra"
	"iris_server/learn/interface/testing"
)

type retriever interface {
	Get(string) string
}

// 判断实现接口的结构的类型
func getType(r retriever) {
	switch v := r.(type) {
	case infra.Retriever:
		fmt.Println("UserAgent is", v.UserAgent)
	case testing.Retriever:
		fmt.Println(v.Content)
	}
}

// 引入接口 示例
func interfaceTest() {
	var curl retriever
	// 实现接口的结构体的值
	curl = infra.Retriever{UserAgent: "chrome", TimeOut: 10}
	// fmt.Printf("%T  %v", curl, curl)
	curl = testing.Retriever{Content: "this is fake"}
	// fmt.Printf("%T  %v", curl, curl)
	// 判断实现接口的结构的类型
	getType(curl)
	// type assertion 编辑器会自动提醒
	realRetriever := curl.(testing.Retriever)
	fmt.Println(realRetriever.Content)
}

// 自定义队列
func queueTest () {
	q := common.Queue{"xiaoming"}
	q.Push("xiaohong")
	fmt.Println(q.IsEmpty())
	fmt.Println(q)
}
