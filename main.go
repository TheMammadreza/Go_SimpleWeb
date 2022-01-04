package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Page struct {
	Link  string
	Port  string
	Title string
	Time  string
}

func main() {
	page := Page{
		Link:  "http://localhost",
		Port:  ":8000",
		Title: "Go is here now!",
		Time:  time.Now().Format(time.Stamp),
	}

	template := template.Must(template.ParseFiles("template/template.html"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if title := r.FormValue("Title"); title != "" {
			page.Title = title
		}
		if err := template.ExecuteTemplate(w, "template.html", page); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Click to open the page: " + page.Link + page.Port + "/")
	fmt.Println(http.ListenAndServe(page.Port, nil))
}
