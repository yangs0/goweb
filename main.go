package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHandle(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>this first page.</h1>")
	}else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "找不到链接")
	}
}

func aboutHandle(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	fmt.Fprint(w, "<h2>This is a about link. </h2>")
}
func main() {

	http.HandleFunc("/", defaultHandle)
	http.HandleFunc("/about", aboutHandle)

	i := http.ListenAndServe(":80", nil)
	log.Println(i)
}

