package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	home := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		tasks := map[string][]Task{
			"Tasks": {
				{Name: "Go to the drug store to buy medicines"},
			},
		}

		tmpl.Execute(w, tasks)
	}

	title := func(w http.ResponseWriter, r *http.Request) {
		RenderFragmentTemplate(w, "title.html")
	}

	form := func(w http.ResponseWriter, r *http.Request) {
		RenderFragmentTemplate(w, "form.html")
	}

	addItem := func(w http.ResponseWriter, r *http.Request) {
		taskName := r.PostFormValue("task_name")
		time.Sleep(1 * time.Second)
		tmpl := template.Must(template.ParseFiles("item.html"))
		task := Task{Name: taskName}
		tmpl.Execute(w, task)
	}

	checkDone := func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		done, _ := strconv.ParseBool(r.URL.Query().Get("done"))
		tmpl := template.Must(template.ParseFiles("item.html"))
		task := Task{Name: name, Done: !done}
		tmpl.Execute(w, task)
	}

	loading := func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		tmpl := template.Must(template.ParseFiles("loading.html"))
		data := Loading{Name: name}
		tmpl.Execute(w, data)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/title", title)
	http.HandleFunc("/form", form)
	http.HandleFunc("/add-item/", addItem)
	http.HandleFunc("/check/done", checkDone)
	http.HandleFunc("/loading", loading)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
