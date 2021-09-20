package common

import (
	"fmt"
	"io/ioutil"
)

/**
 * 读取文件
 */
func ReadFile(filename string) (string, error){
	if content, err := ioutil.ReadFile(filename); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%s", content), nil
	}
}
