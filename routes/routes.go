package routes

import (
	u "favart-api/utility"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const basePath = "./media/"

// AppRouter defines all of the routes for the application.
func AppRouter() *Router {
	r := NewRouter()

	r.Get("/", index)
	r.Get("/media", getMedia)
	r.Post("/media", addMedia)
	r.Get("/file", file)

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	m := u.PlainTextMessage{Message: "Hello world!"}
	u.Respond(w, http.StatusOK, m)
}

func getMedia(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		e := u.ErrorMessage{Error: err.Error()}
		u.Respond(w, http.StatusInternalServerError, e)
		return
	}

	var names []string
	for _, file := range files {
		name := file.Name()
		isValidImageFile := strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".png")

		if !isValidImageFile && !file.IsDir() {
			continue
		}
		names = append(names, name)
	}

	u.Respond(w, http.StatusOK, names)
}

func addMedia(w http.ResponseWriter, r *http.Request) {
	path := r.PostFormValue("path")
	if path == "" {
		e := u.ErrorMessage{Error: "missing required parameter 'path'"}
		u.Respond(w, http.StatusBadRequest, e)
		return
	}

	err := os.MkdirAll(basePath+path, os.ModePerm)
	if err != nil {
		e := u.ErrorMessage{Error: err.Error()}
		u.Respond(w, http.StatusInternalServerError, e)
		return
	}

	m := u.PlainTextMessage{Message: "created"}
	u.Respond(w, http.StatusCreated, m)
}

func file(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	if path == "" {
		path = "./media"
	}

	id := r.FormValue("id")
	if id == "" {
		e := u.ErrorMessage{Error: "missing required parameter 'id'"}
		u.Respond(w, http.StatusBadRequest, e)
		return
	}

	f := path + "/" + id
	http.ServeFile(w, r, f)
}
