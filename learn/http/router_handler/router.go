package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Run(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}


type appHandler func(writer http.ResponseWriter, request *http.Request) error


func ErrWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Printf("error is %s\n", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err) : // 文件找不到
				code = http.StatusNotFound
			case os.IsPermission(err): // 访问没有权限的文件
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError // 服务器内部错误 500
			}
			log.Printf("http code is %d \n", code)
			http.Error(writer, http.StatusText(code), code)
		}
	}
}