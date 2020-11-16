package tools

import (
	"net/http"
	"io"
	"log"
	"os"
)

/**
* 请求地址并且把结果打印到控制台
*/
func Curl() {
	//执行请求
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)		
	}

	//把结果保存到文件内
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	//同时输出到控制台与文件
	dest := io.MultiWriter(os.Stdout, file)

	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}