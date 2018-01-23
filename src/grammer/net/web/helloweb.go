package main

import (
	"net/http"
	"io"
	"log"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello web!!!")
}

func main(){

	//URL对于的handler
	http.HandleFunc("/hello",helloHandler)
	//监听端口
	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err.Error())
	}

}
