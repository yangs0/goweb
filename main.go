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

	router := http.NewServeMux()
	router.HandleFunc("/", defaultHandle)
	router.HandleFunc("/about", aboutHandle)

	router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, "访问文章列表")
		case "POST":
			fmt.Fprint(w, "创建新的文章")
		}
	})

	i := http.ListenAndServe(":80", router)
	log.Println(i)
}

