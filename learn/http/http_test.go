package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)


func TestBody(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}