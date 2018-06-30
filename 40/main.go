package main

type IntPair [2]int

// ペアを返すメソッド
func (ip IntPair) First() int {
	return ip[0]
}

// ペアの末尾を返す
func (ip IntPair) Last() (i int) {
	i = ip[1]

	// (i int)で定義しているため、iが返る
	return
}

type Strings []string // スライス型

func (s Strings) Join(d string) (sum string) {
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}

	// (sum string)で定義しているため、sumが返る
	return
}

func main() {
	ip := IntPair{1, 2}
	println(ip.First)
	println(ip.Last)
	println(Strings{"A", "B", "C"}.Join(","))
}
