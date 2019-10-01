package main

import (
	"net/http"
)

func server() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, req *http.Request) {
		writer.Write([]byte("Hi JohnMark!"))
	})

	http.ListenAndServe(":8088", nil)
}
