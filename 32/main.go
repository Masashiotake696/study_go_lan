package main

type Callback func(i int) int

// 合計値をコールバック関数に渡す関数
func Sum(ints []int, c Callback) int {
	var sum int
	for _, v := range ints {
		sum += v
	}
	return c(sum)
}

// goにはクラスはないが構造体(継承がない構造体)はある
// 構造体はオブジェクト指向ではなく、何かの塊というイメージ
func main() {
	type MyInt int // MyIntはintのエイリアス

	var n1 MyInt = 5
	n2 := MyInt(7)
	println(n1)
	println(n2)

	// 変数名の先頭を大文字にすると外部パッケージで呼び出せる
	type (
		IntPair     [2]int                // 中身がint型で長さが2の配列
		Strings     []string              // 中身がstring型のスライス
		areaMap     map[string][2]float64 // keyがstringで中身が長さ2でfloat型の配列をもつマップ
		IntsChannel chan []int            // 中身がint型のチャネル
	)

	_ = IntPair{1, 2}
	_ = Strings{"Apple", "Banana"}
	_ = areaMap{"Tokyo": {35.111, 12.111}}
	_ = make(IntsChannel)

	Sum([]int{1, 2, 3}, func(a int) int {
		return a + 1
	})
}
