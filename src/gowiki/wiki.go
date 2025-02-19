package main

import (
	"os"
	"net/http"
	"fmt"
	"log"
	"html/template"
)

type Page struct {
	Title 	string
	Body 	[]byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
}

return &Page{Title: title, Body: body}, nil
}

// Visualizar y cargar las páginas existentes
func viewHandler (w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body)
}

// Renderizar y cargar plantillas html
func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, _ := loadPage(title)

	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)

}

func main() {
	 
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
 