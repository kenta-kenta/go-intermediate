package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	http.Handle("/", myMiddleware2(myMiddleware1(helloHandler)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 前処理
		io.WriteString(w, "Pre-process1\n")
		// もとハンドラを実行
		next.ServeHTTP(w, r)
		// 後処理
		io.WriteString(w, "Post-process1\n")
	})
}

func myMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 前処理
		io.WriteString(w, "Pre-process2\n")
		// もとハンドラを実行
		next.ServeHTTP(w, r)
		// 後処理
		io.WriteString(w, "Post-process2\n")
	})
}
