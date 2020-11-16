package demos

import (
	"fmt"
	"bytes"
	"mime/multipart"
	"os"
	"path/filepath"
	"io"
	"net/http"
	"io/ioutil"
)

/**
* 使用原生代码发送post请求示例
*/
func SendPostReq() {
	fmt.Println("send Post Req demo")
	//地址
	url := "http://123.com/show"
	method := "POST"
	//实例化multipart
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	//普通参数写入
	_ = writer.WriteField("name", "normal param name")
	_ = writer.WriteField("age", "normal param age")
	//文件参数写入
	file, errFile1 := os.Open("/C:/test.xlsx")
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base("/C:/test.xlsx"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	//
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建http客户端
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	//添加Cookie 
	req.Header.Add("Cookie", "token=1233456")
	//设置请求方式，
	req.Header.Set("Content-Type", writer.FormDataContentType())
	//发起请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	//获取返回结果
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

/**
* 使用原生代码发送get请求示例
*/
func SendGetReq() {
	fmt.Println("send Post Req demo")
	//地址
	url := "http://123.com/show?id=5"
	method := "GET"
	//实例化multipart
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建客户端
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	//设置请求头
	req.Header.Add("Cookie","token=123")
	//设置请求方式
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	//获取结果
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}