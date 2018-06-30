package main

type User struct {
	Id   int
	Name string
}

func main() {
	m1 := map[User]string{
		// Userはつけてもつけなくても良い
		User{Id: 1, Name: "Taro"}: "Tokyo",
		{Id: 2, Name: "Jaro"}:     "Osaka",
	}

	m2 := map[int]User{
		1: User{Id: 11, Name: "Bob"},
	}

	m3 := map[int]map[int]string{
		1: {4: "Apple", 2: "Banana"},
	}
}
