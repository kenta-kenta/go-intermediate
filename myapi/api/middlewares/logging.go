package middlewares

import (
	"log"
	"net/http"
)

// 自作ResponseWriterを作る
type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

// コンストラクタを作る
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// WriteHeaderメソッドを作る
func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

// ミドルウェアの中身
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// リクエスト情報をロギング
		log.Println(req.RequestURI, req.Method)

		// 自作のResponseWriterを作って
		rlw := NewResLoggingWriter(w)

		// それをハンドラに渡す
		next.ServeHTTP(rlw, req)

		// 自作ResponseWriterからロギングしたいデータを出す
		log.Println("res: ", rlw.code)
	})
}
