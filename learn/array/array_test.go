package array_test

import (
	"fmt"
	"testing"
)

// 数组在go语言中用的比较少
func TestArrayInit(t *testing.T) {
	var a [5]int
	fmt.Println(a)
	b := [4]int{1,2,3,4}
	t.Log(b)
}

// 两种遍历方式
func TestArraySlice(t *testing.T) {
	a := [6]int{1,2,3,4,5,6}
	for i:=0; i<len(a); i++ {
		t.Log(a[i])
	}
	for _,v := range a {
		t.Log(v)
	}
}

func changeArrayValue(a *[5]int) {
	a[4] = 20
}

// 数组作函数参数
func TestArrayChange(t *testing.T){
	a := [5]int{1,2,3,4,5}
	changeArrayValue(&a)
	t.Log(a)
}

// 切片的初始化
func TestSliceInit(t *testing.T){
	a := [...]int{1,2,3,4,5}
	b := append(a[0:], 2)
	t.Log(b)

	var c []int
	t.Log(c)
	d := []int{1,2,3,4}
	d[0] = 2
	t.Log(d)
	d = append(d, 1)
	t.Log(len(d), cap(d))
	e := make([]int, 2, 4)
	t.Log(e)
}

func sliceChange(a []int) {
	a[2] = 10
}

// 切片作函数参数
func TestSliceChange(t *testing.T){
	a := []int{1,2,3,4}
	sliceChange(a)
	t.Log(a)
}

// map的初始化
func TestMapInit(t *testing.T){
	m := map[string]string{}
	m["name"] = "mengyuan"
	m["age"] = "12"
	m["gender"] = "男"
	if name,ok := m["name"]; ok{
		t.Log(name)
	} else {
		t.Log("error")
	}
	for _,v := range m {
		//t.Log(k)
		t.Log(v)
	}
}

type User struct {
	Uid int
	Name string
	Score int
}


func TestMapComplex(t *testing.T) {
	scores := make([]User, 0)
	scores = append(scores, User{
		Uid: 1,
		Name: "语文",
		Score: 2,
	})
	scores = append(scores, User{
		Uid: 1,
		Name: "数学",
		Score: 12,
	})
	// 数据组装
	ret := map[int][]int{}
	for _, items := range scores {
		ret[items.Uid] = append(ret[items.Uid], items.Score)
	}
	t.Log(ret)
	// format index
	userIndex := make(map[int]User)
	for _, v := range scores {
		if _,ok := userIndex[v.Uid]; !ok {
			userIndex[v.Uid] = User{
				Uid: v.Uid,
				Name: v.Name,
				Score: v.Score,
			}
		}
	}
	t.Log(userIndex)
}

