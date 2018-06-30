package main

const X = 1 // 大文字にすると外部から呼び出せる

func main() {
	// const X = 1
	const (
		X = 1
		Y = 1
		Z = 1
	)

	const (
		XX = 1 // YとZにも1が入る
		YY
		ZZ
		XX1, YY1 = 33, 33
	)

	const (
		XXX1 = 1
		YYY1
		ZZZ1
		XXX2 = "あ"
		YYY2
	)

	const (
		XXXX = 2
		YYYY = 3
		ZZZZ = XXXX + YYYY
	)

	const (
		XXXXX = `afdjalfjdl
		fdajfla
		fjldj

		fajlfjdal`
	)

	const (
		XYZ1 = iota
		XYZ2
		XYZ3
	)
	println(XYZ1)
	println(XYZ2)
	println(XYZ3)

	const (
		A = iota // 1
		B        // 2
		C        // 3
		D = 17   // 17
		E = iota // 5
		F        // 6
	)
	println(A)
	println(B)
	println(C)
	println(D)
	println(E)
	println(F)

	const (
		AA = 999  // 999
		BB = iota // 1
		CC        // 2
	)

	println(AA)
	println(BB)
	println(CC)
}
