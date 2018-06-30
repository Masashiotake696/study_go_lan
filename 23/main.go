package main

func main() {
	s := []string{"Apple", "Banana", "Cherry"}
	for _, v := range s {
		s = append(s, "Melon")
		println(v)
	}

	for i := 0; i < len(s); i++ {
		println(s[i])
		// s = append(s, "Melon")
	}
}
