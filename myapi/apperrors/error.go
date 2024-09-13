package apperrors

type MyAppError struct {
	// ErrCode型のErrCodeフィールド
	// (フィールド名を省略した場合、型名がそのままフィールド名になる)
	ErrCode

	// string型のMessageフィールド
	Message string

	Err error
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
