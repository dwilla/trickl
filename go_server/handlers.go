package main

import (
	"net/http"

	"github.com/dwilla/trickl/views"
)

func handlePost(w http.ResponseWriter, r *http.Request) {
	view := views.Home(Feed[CurrentPost].Post.Record.Text)

	CurrentPost++

	view.Render(r.Context(), w)
}
