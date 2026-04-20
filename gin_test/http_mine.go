package gintest

import (
	"fmt"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hello world")
	res.Write([]byte("hello world"))
}

func MyHttp() {
	http.HandleFunc("/index", Index)
	fmt.Println(":8080 server running")
	http.ListenAndServe(":8080", nil)
}
