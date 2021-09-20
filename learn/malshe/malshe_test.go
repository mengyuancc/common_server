package malshe

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Order struct {
	Id string `json:"id"`
	// Name string `json:"name,omitempty"` //允许为空
	Items []OrderItem  `json:"items"`
	Num int            `json:"num" `
	TotalPrice float64 `json:"price"`
}

type OrderItem struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

func TestMarshal(t *testing.T) {
	order := Order{
		Id: "order_1",
		Items: []OrderItem{
			{Id: "item_1", Name: "learn go", Price: 15},
			{Id: "item_2", Name: "interview", Price: 25},
		},
		Num: 2,
		TotalPrice: 12.1,
	}
	s, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)

}

func TestUnmarshal(t *testing.T) {
	orderStr := `{"id":"order_1","items":[{"id":"item_1","name":"learn go","price":15},{"id":"item_2","name":"interview","price":25}],"num":2,"price":12.1}`
	var o Order
	err := json.Unmarshal([]byte(orderStr), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

func TestNLPMap(t *testing.T) {
	orderStr := `{"id":"order_1","items":[{"id":"item_1","name":"learn go","price":15},
					{"id":"item_2","name":"interview","price":25}],"num":2,"price":12.1}`
	// map的方式
	o := make(map[string]interface{})
	err := json.Unmarshal([]byte(orderStr), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
	// type assertion 的方式去获取数据
	fmt.Println(o["items"].([]interface{})[0].(map[string]interface{})["id"])
}

func TestNLPStruct(t *testing.T) {
	orderStr := `{"id":"order_1","items":[{"id":"item_1","name":"learn go","price":15},
					{"id":"item_2","name":"interview","price":25}],"num":2,"price":12.1}`
	// 以struct的方式
	o := struct {
		Items []struct{
			Name string `json:"name"`
		} `json:"items"`
	}{}
	err := json.Unmarshal([]byte(orderStr), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
	// type assertion
	fmt.Println(o.Items[0].Name)
}
