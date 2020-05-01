package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	// text := "{{ . }}\n"

	// tpl, err := template.New("").Parse(text)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := tpl.Execute(os.Stdout, "Hello, World"); err != nil {
	// 	log.Fatal(err)
	// }

	// 	text := `Name: {{ .name }}
	// Age: {{ .age }}`

	// 	tpl, err := template.New("").Parse(text)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	m := map[string]interface{}{
	// 		"name": "Otake",
	// 		"age":  24,
	// 	}

	// 	if err := tpl.Execute(os.Stdout, m); err != nil {
	// 		log.Fatal(err)
	// 	}

	text := `Name: {{ .Name }}
Age: {{ .Age }}`

	tpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	v := struct {
		Name string
		Age  int
	}{
		Name: "Otake",
		Age:  24,
	}

	if err := tpl.Execute(os.Stdout, v); err != nil {
		log.Fatal(err)
	}
}
