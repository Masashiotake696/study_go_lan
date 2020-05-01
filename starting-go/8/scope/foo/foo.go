package foo

// 定数
const (
	A   = 1 // 外部公開
	abc = "abc"
)

// パッケージ変数
var (
	m = 256
	N = 512 // 外部公開
)

// DoSomething は何かしらをする関数です
// 外部公開
func DoSomething() {
}

func doSomething() {
}
