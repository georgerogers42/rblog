package rblog

import (
	"github.com/gorilla/mux"
	"net/http"
)

var App = mux.NewRouter()

func init() {
	App.HandleFunc("/", Index)
	App.HandleFunc("/{slug}", Slug)
}

func Index(w http.ResponseWriter, r *http.Request) {}
func Slug(w http.ResponseWriter, r *http.Request)  {}

