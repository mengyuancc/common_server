package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) {
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("content is \n %s", content)
	}
}

func main() {
	readFile("abc.txt")
}
