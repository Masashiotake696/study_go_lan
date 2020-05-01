package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

var tpl *template.Template

func handler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("p") == "Gopher" {
		if err := tpl.Execute(w, nil); err != nil {
			log.Println(err)
		}
	} else {
		switch s := rand.Intn(6) + 1; s {
		case 6:
			fmt.Fprint(w, "大吉")
		case 5, 4:
			fmt.Fprint(w, "中吉")
		case 3, 2:
			fmt.Fprint(w, "吉")
		default:
			fmt.Fprint(w, "凶")
		}
	}
}

func main() {
	html := "<html><body>Gopherさんの運勢は「<b>大吉</b>」です</body></html>"
	tpl = template.Must(template.New("").Parse(html))

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
