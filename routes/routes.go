package routes

import (
	u "favart-api/utility"
	"fmt"
	"io"
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
	r.Get("/file", getFile)
	r.Post("/file", addFile)

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	m := u.PlainTextMessage{Message: "Hello world!"}
	u.Respond(w, http.StatusOK, m)
}

func getMedia(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	files, err := ioutil.ReadDir(basePath + path)
	if err != nil {
		e := u.ErrorMessage{Error: err.Error()}
		u.Respond(w, http.StatusInternalServerError, e)
		return
	}

	var fileInfos []u.FileInfoMessage
	for _, file := range files {
		var info u.FileInfoMessage

		name := file.Name()
		isValidImageFile := strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".png")

		if !isValidImageFile && !file.IsDir() {
			continue
		}

		info.Name = name
		info.Size = file.Size()

		fileInfos = append(fileInfos, info)
	}

	u.Respond(w, http.StatusOK, fileInfos)
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

func getFile(w http.ResponseWriter, r *http.Request) {
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

func addFile(w http.ResponseWriter, r *http.Request) {
	path := r.PostFormValue("path")
	upload, header, err := r.FormFile("upload")
	defer upload.Close()

	if err != nil {
		e := u.ErrorMessage{Error: err.Error()}
		u.Respond(w, http.StatusInternalServerError, e)
		return
	}

	fp := basePath + header.Filename
	if path != "" {
		fp = fmt.Sprintf("%s%s/%s", basePath, path, header.Filename)
	}

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()

	if err != nil {
		e := u.ErrorMessage{Error: err.Error()}
		u.Respond(w, http.StatusInternalServerError, e)
		return
	}

	io.Copy(file, upload)

	m := u.PlainTextMessage{Message: "created"}
	u.Respond(w, http.StatusCreated, m)
}
