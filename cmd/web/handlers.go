package main

import (
	"fmt"
	"net/http"

	"github.com/adelplace/snippetbox/pkg/models"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	result, err := app.snippets.Latest()
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	}
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Snippets: result})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	if id == "" {
		app.notFound(w)
		return
	}

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

	app.render(w, r, "show.page.tmpl", &templateData{Snippet: result})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	id := primitive.NewObjectID()
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	_, err = app.snippets.Insert(&id, title, content)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%s", id.Hex()), http.StatusSeeOther)
}
