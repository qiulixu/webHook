package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func webHook(w http.ResponseWriter, r *http.Request)  {
	body,err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(fmt.Sprintf("读取文件异常 err[%s]",err))
	}
	fmt.Println(string(body))
}

func main(){
	http.HandleFunc("/webHook",webHook)
	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		panic(fmt.Sprintf("服务启动异常 err[%s]",err))
	}
}
