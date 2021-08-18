package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeHandle(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
}

func aboutHandle(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func noFountHandle(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}
func articlesShowHandle(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章 ID："+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "创建新的文章")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandle).Name("home")
	router.HandleFunc("/about", aboutHandle).Methods("GET").Name("about")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandle).Methods("GET").Name("articles.show")

	router.NotFoundHandler = http.HandlerFunc(noFountHandle)

	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "23")
	fmt.Println("articleURL: ", articleURL)

	i := http.ListenAndServe(":80", router)
	log.Println(i)
}

