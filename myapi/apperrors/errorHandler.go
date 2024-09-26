package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/kenta-kenta/go-intermediate-myapi/api/middlewares"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError

	// errors.As関数で引数のerrをMyAppError型のappErrに変換する
	if !errors.As(err, &appErr) {
		// もし変換に失敗したらUnknownエラーを変数appErrに手動で格納
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := middlewares.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	// ユーザーに返すHTTPレスポンスコードを収めるための変数
	var statusCode int

	// StatusCodeに収めるレスポンスコードの内容を決める
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
