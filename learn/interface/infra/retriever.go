package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct {
	UserAgent string
	TimeOut int
}

func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
