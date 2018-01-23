package main

import (
	"net/http"
	"fmt"
)

const SERVER_PORT = 8080
const SERVER_DOMAIN  = "localhost"
const RESPONSE_TEMPLATE = "hello"
const WEB_ROOT ="src/grammer/https"

func rootHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","text/html")
	w.Header().Set("Content-Length",fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func main(){
	http.HandleFunc(fmt.Sprintf("%s:%d/",SERVER_DOMAIN,SERVER_PORT),rootHandler)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d",SERVER_PORT),WEB_ROOT+"/rui.crt",WEB_ROOT+"/rui.key",nil)
	if err != nil{
		fmt.Println(err.Error())
	}
}
