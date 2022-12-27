package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"logo":  "Company",
		"title": " Golang ",
		"text":  "Golang Adalah Salah Satu Bahasa Pemograman",
		"path":  "1.png",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func handlerDetail(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "detail.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"logo":  "Company",
		"title": " Golang ",
		"text":  "Golang Adalah Salah Satu Bahasa Pemograman",
		"path":  "1.png",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/detail", handlerDetail)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
