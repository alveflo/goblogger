package httphandlers

import (
	"net/http"

	"../entities"
	"../templates"
)

func ViewPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := entities.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}

	templates.RenderTemplate(w, "view", page)
}
