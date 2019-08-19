package controller

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	RequestType := r.Method
	RequestPath := r.URL.Path

	switch RequestType {
	case "GET":
		switch RequestPath {
		case "/health":
			Health(w, r)
		}
	case "POST":
		switch RequestPath {
		case "/download":
			Download(w, r)
		}
	}
}

