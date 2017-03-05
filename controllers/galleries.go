package controllers

import (
	"fmt"
	"net/http"

	"usegolang.com/views"
)

type Galleries struct {
	NewView *views.View
}

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

// New is used to render the form where a user can create a new user account
//
// GET /galleries
func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
