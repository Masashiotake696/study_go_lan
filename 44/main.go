package main

/**
GoLangにはerror型がある
type error interface {
	Error() string // 返り値がString型のメソッド
}
**/

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Messages
}

func Raiseerror() error {
	return MyError{Message: "エラーだよ", ErrCode: 1234}
}

func main() {

}

func tmp() (err error) {
	return
}
