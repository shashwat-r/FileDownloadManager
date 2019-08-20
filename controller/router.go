package controller

import (
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	RequestType := r.Method
	RequestPath := r.URL.Path

	switch RequestType {
	case "GET":
		sections := strings.Split(RequestPath[1:], "/")
		if sections[0] == "health" {
			Health(w, r)
		} else if sections[0] == "downloads" {
			id := sections[1]
			WriteJson(w, id)
		} else if sections[0] == "files" {
			ListFiles(w)
		}
	case "POST":
		sections := strings.Split(RequestPath[1:], "/")
		if sections[0] == "download" {
			Download(w, r)
		}
	}
}

