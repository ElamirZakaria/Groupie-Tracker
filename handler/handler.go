package handler

import (
	"bytes"
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type App struct {
	Artists []models.Artist
}
type errorData struct {
	Msg    string
	Status int
}

var tmpl *template.Template

func render(w http.ResponseWriter, status int, templateName string, data any) {

	var err error
	if tmpl == nil {
		tmpl, err = template.ParseFiles("templates/artist.html", "templates/index.html", "templates/error.html")
		if err != nil {
			renderError(w, errorData{Msg: "500: Internal Server Error. Failed to load Templates", Status: http.StatusInternalServerError})
			return

		}
	}
	var buff bytes.Buffer
	err = tmpl.ExecuteTemplate(&buff, templateName, data)
	if err != nil {
		renderError(w, errorData{Msg: "500 Internal Server Error", Status: http.StatusInternalServerError})
		return
	}
	w.WriteHeader(status)
	buff.WriteTo(w)
}
func renderError(w http.ResponseWriter, data errorData) {

	var err error
	if tmpl == nil {
		tmpl, err = template.ParseFiles("templates/artist.html", "templates/index.html", "templates/error.html")

		if err != nil {
			http.Error(w, data.Msg, data.Status)
			return
		}
	}
	var buff bytes.Buffer
	err = tmpl.ExecuteTemplate(&buff, "error.html", data)
	if err != nil {
		http.Error(w, data.Msg, data.Status)
		return
	}
	w.WriteHeader(data.Status)
	buff.WriteTo(w)

}
func (app *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, errorData{Msg: "404 Page Not Found", Status: http.StatusNotFound})
		return
	}
	if r.Method != http.MethodGet {
		renderError(w, errorData{Msg: "405 Method Not Allowed", Status: http.StatusMethodNotAllowed})
		return
	}
	render(w, http.StatusOK, "index.html", app.Artists)

}
func (app *App) ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, errorData{Msg: "405 Method Not Allowed", Status: http.StatusMethodNotAllowed})
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		renderError(w, errorData{Msg: "Error 400 Bad Request", Status: http.StatusBadRequest})
		return
	}
	if id <= 0 || id > len(app.Artists) {
		renderError(w, errorData{Msg: "Error 404 Artist Not Found", Status: http.StatusNotFound})
		return
	}
	render(w, http.StatusOK, "artist.html", app.Artists[id-1])

}
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, errorData{Msg: "405 Method Not Allowed", Status: http.StatusMethodNotAllowed})
		return
	}
	if strings.HasSuffix(r.URL.Path, "/") {
		renderError(w, errorData{Msg: "404 page not found", Status: http.StatusNotFound})
		return
	}
	filePath := strings.TrimPrefix(r.URL.Path, "/static/")
	if info, err := os.Stat("static/" + filePath); err != nil || info.IsDir() {
		renderError(w, errorData{Msg: "404 page not found", Status: http.StatusNotFound})
		return
	}
	fs := http.FileServer(http.Dir("static"))
	http.StripPrefix("/static/", fs).ServeHTTP(w, r)
}
