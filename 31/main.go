package main

func main() {
	p := &[3]int{1, 2, 3}
	println((*p)[1]) // ポインタ配列の値を参照する
	println(p[1])    // こう書いても良い
	println(cap(p))  // capはcapacity
	println(len(p))  // lenはlength

	for _, v := range p { // keyを握りつぶしてループを回す
		println(v)
	}
}
