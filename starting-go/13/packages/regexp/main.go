package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.MatchString("A", "ABC"))
	fmt.Println(regexp.MatchString("A", "XYZ"))

	regexp.Compile("A")
	regexp.MustCompile("ABC")
	regexp.MustCompile("\\d")
	regexp.MustCompile(`\d`)

	re1 := regexp.MustCompile(`(?i)abc`)
	fmt.Println(re1.MatchString("ABC"))

	re2 := regexp.MustCompile(`^XYZ$`)
	fmt.Println(re2.MatchString("XYZ"))
	fmt.Println(re2.MatchString(" XYZ "))

	re3 := regexp.MustCompile(`bc`)
	fmt.Println(re3.MatchString("abcdefg"))

	re4 := regexp.MustCompile(`.`)
	fmt.Println(re4.MatchString("ABC"))
	fmt.Println(re4.MatchString("日本語"))
	fmt.Println(re4.MatchString("\n"))

	re5 := regexp.MustCompile(`abc|xyz`)
	fmt.Println(re5.MatchString("abc"))
	fmt.Println(re5.MatchString("xyz"))

	re6 := regexp.MustCompile(`a+b*`)
	fmt.Println(re6.MatchString("ab"))
	fmt.Println(re6.MatchString("a"))
	fmt.Println(re6.MatchString("aaaaaaaaaaaaaaab"))
	fmt.Println(re6.MatchString("b"))

	re7 := regexp.MustCompile(`A+?A+?X`)
	fmt.Println(re7.MatchString("AX"))
	fmt.Println(re7.MatchString("AAX"))
	fmt.Println(re7.MatchString("AAAX"))
	fmt.Println(re7.MatchString("AAAAAX"))

	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("X"))
	fmt.Println(re8.MatchString("Y"))
	fmt.Println(re8.MatchString("Z"))
	fmt.Println(re8.MatchString("A"))

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re9.MatchString("ABC"))
	fmt.Println(re9.MatchString("x01"))
	fmt.Println(re9.MatchString("abcdefg"))
	fmt.Println(re9.MatchString("日本語"))

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re10.MatchString("ABC"))
	fmt.Println(re10.MatchString("x01"))
	fmt.Println(re10.MatchString("abcdefg"))
	fmt.Println(re10.MatchString("日本語"))

	re11 := regexp.MustCompile(`\d+`)
	fmt.Println(re11.MatchString("12345"))
	fmt.Println(re11.MatchString("X=1"))
	fmt.Println(re11.MatchString("abcde"))

	re12 := regexp.MustCompile(`^\p{Katakana}+$`)
	fmt.Println(re12.MatchString("アイウエオ"))
	fmt.Println(re12.MatchString("ｱｲｳｴｵ"))
	fmt.Println(re12.MatchString("あいうえお"))

	re13 := regexp.MustCompile(`(佐藤|鈴木)(太郎|花子)`)
	fmt.Println(re13.MatchString("佐藤太郎"))
	fmt.Println(re13.MatchString("佐藤花子"))
	fmt.Println(re13.MatchString("鈴木花子"))

	re14 := regexp.MustCompile(`\w+`)
	fmt.Println(re14.FindString("abc xyz 999"))
	fmt.Println(re14.FindAllString("abc xyz 999", 2))
	fmt.Println(re14.FindAllString("abc xyz 999", -1))

	re15 := regexp.MustCompile(`\s+`)
	fmt.Println(re15.Split("A B C D\tE", 3))
	fmt.Println(re15.Split("A B C D\tE", -1))

	re16 := regexp.MustCompile(`佐藤`)
	fmt.Println(re16.ReplaceAllString("佐藤さん鈴木さん", "田中"))
	fmt.Println(re16.ReplaceAllString("XYZ", "田中"))

	re17 := regexp.MustCompile(`、|。`)
	fmt.Println(re17.ReplaceAllString("私は、Go言語を使用する、プログラマーです。", ""))

	re18 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
00-1111-2222
3333-44-55
666-777-888
9-9-9`
	ms := re18.FindAllStringSubmatch(s, -1)
	for _, v := range ms {
		fmt.Println(v)
	}

	re19 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	fmt.Println(re19.ReplaceAllString("Tel: 000-111-222", "$3-$2-$1"))
}
