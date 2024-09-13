package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
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
