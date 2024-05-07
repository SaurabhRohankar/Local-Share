package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type upload struct {
	l *log.Logger
}

func NewUpload(l *log.Logger) *upload {
	return &upload{l}
}

func (u upload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		u.l.Println("Error while reading file: ", err)
	}

	dst, createFileErr := os.Create(filepath.Join("data/", header.Filename))
	if createFileErr != nil {
		u.l.Println("Error while creating file: ", createFileErr)
	}

	_, writeErr := io.Copy(dst, file)
	if writeErr != nil {
		u.l.Println("Error while writing file", writeErr)
	}

	u.l.Printf("File written to %v", dst)
}
