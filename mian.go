package main

import (
	"fmt"
	"net/http"
)

func webHook(r http.ResponseWriter, w *http.Request)  {
	fmt.Println(r,w)
}

func main(){
	http.HandleFunc("/webHook",webHook)
	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		panic(fmt.Sprintf("服务启动异常 err[%s]",err))
	}
}
