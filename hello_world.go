package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/home.tmpl")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, "Текст, выводимый на странице")
}
