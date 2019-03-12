package main

import (
	"fmt"
	"html/template"
	"net/http"
	"snippetbox/pkg/models"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		app.notFound(w)
		return
	}
	result, err := app.snippets.Get(objectID)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	}
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "%v", result)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id := primitive.NewObjectID()
	title := "toto"
	content := "content"

	_, err := app.snippets.Insert(&id, title, content)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Print("insert")
	fmt.Fprintf(w, "Display a specific snippet with ID %s...", id.Hex())
	//http.Redirect(w, r, fmt.Sprintf("/snippet?id=%s", id.Hex()), http.StatusSeeOther)
}
