package main

import (
	"fmt"
	"github.com/guotie/gogb2312"
	"io/ioutil"
	"net/http"
)

const (
	URI string = "http://www.xiexingcun.com/wenxueyuedu/index.html"
)

func main() {
	res, err := http.Get(URI)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	output, err, _, _ := gogb2312.ConvertGB2312(body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
