package main

import (
	"fmt"
	"net/http"
)

func handelFunc(w http.ResponseWriter, r *http.Request)  {

	if r.URL.Path == "/" {
		_, _ = fmt.Fprint(w, "<h1>根路径</h1>"+r.URL.Path)
	}else if r.URL.Path == "/about" {
		_, _ = fmt.Fprint(w, "<h2>about 路径</h2>"+r.URL.Path)
	}else {
		_, _ = fmt.Fprint(w, "其他路径")
	}
}
func main() {
	http.HandleFunc("/", handelFunc)
	_ = http.ListenAndServe(":80", nil)
}

