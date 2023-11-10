package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const dataPath = "data/"
const templatePath = "templates/"
const isDev = true

var templates = template.Must(template.New("templates").Funcs(template.FuncMap{
	"unescapeHTML": unescapeHTML,
}).ParseFiles(
	templatePath+"edit.html",
	templatePath+"view.html",
	templatePath+"index.html"),
)
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  string
}

func (p *Page) save() error {
	filename := dataPath + p.Title + ".txt"
	return os.WriteFile(filename, []byte(p.Body), 0600)
}

func loadPage(title string) (*Page, error) {
	filename := dataPath + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(body)}, nil
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		log.Println("not found")
		return "", errors.New("Invalid Page Title")
	}
	log.Println("rendering", m)
	return m[2], nil // The title is the second subexpression
}

func renderTempalte(w http.ResponseWriter, templateName string, p *Page) {
	// In development environment, caching is unnecessary
	if isDev {
		filename := "templates/" + templateName + ".html"
		content, _ := os.ReadFile(filename)
		t, _ := template.
			New(templateName).
			Funcs(template.FuncMap{
				"unescapeHTML": unescapeHTML,
			}).
			Parse(string(content))
		t.Execute(w, p)

		// t, err := template.ParseFiles(filename)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
		// t.Execute(w, p)
	} else {
		err := templates.ExecuteTemplate(w, templateName+".html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func unescapeHTML(s string) template.HTML {
	return template.HTML(s)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTempalte(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTempalte(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: body}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/welcome", "/index", "/index.htm":
		renderTempalte(w, "index", nil)
	default:
		http.FileServer(http.Dir("public")).ServeHTTP(w, r)
	}
}

func listPages(w http.ResponseWriter, r *http.Request) {
	files, _ := os.ReadDir("data/")
	names := make([]string, len(files))
	for i, de := range files {
		name := de.Name()[:strings.LastIndex(de.Name(), ".txt")]
		names[i] = name
	}
	json, err := json.Marshal(names)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(json)
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/list", listPages)
	// address := ":8000"
	address := "localhost:8000"
	fmt.Printf("👌 Wiki is running on: http://%s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
