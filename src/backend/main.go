package main

import ("net/http"
"github.com/gorilla/mux")

func main(){
	// http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("hello world!"))
	// })
	http.ListenAndServe(":8080",nil)
}