package http_test

import (
	"fmt"
	"io/ioutil"
	router "iris_server/learn/http/router_handler"
	"log"
	"net/http"
	"os"
	"testing"
)


func TestRequest(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	resp, err := http.DefaultClient.Do(request)
	// 请求失败直接抛出错误
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

func TestWebServer1(t *testing.T) {
	// 文件服务器
	http.HandleFunc("/list/",
		func(writer http.ResponseWriter, request *http.Request) {
			path := request.URL.Path[len("/list/"):]
			log.Print(path)
			file, err := os.Open(path)
			// 1.直接抛出err
			/* if err != nil {
				panic(err)
			} */
			// 2.使用http.err
			if err != nil {
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				// http.Error(writer, err.Error(), http.StatusInternalServerError)
				// panic(err)
				return
			}
			defer file.Close()
			all, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			writer.Write(all)
		})

	err := http.ListenAndServe(":8889", nil)
	if err != nil {
		panic(err)
	}
}

func TestWebServer2(t *testing.T) {
	http.HandleFunc("/list/", router.ErrWrapper(router.Run))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
