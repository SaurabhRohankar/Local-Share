package handlers

import (
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

type home struct {
	l *log.Logger
}

func init() {
	tmpl = template.Must(template.ParseFiles("frontend/upload.html"))
}

func NewHome(l *log.Logger) *home {
	return &home{l}
}

func (h home) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	tmpl.Execute(rw, "upload.html")
}
