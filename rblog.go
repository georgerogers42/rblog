package rblog

import (
	"database/sql"
	"github.com/georgerogers42/rblog/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

var DB *sql.DB

var App = mux.NewRouter()

func init() {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	App.HandleFunc("/", Index)
	App.HandleFunc("/{slug}", Slug)
}

var BaseTpl = template.Must(template.ParseFiles("templates/base.tpl"))

var IndexTpl = template.Must(BaseTpl.ParseFiles("templates/index.tpl"))

func Index(w http.ResponseWriter, r *http.Request) {
	env := map[string]interface{}{}
	articles, err := models.AllArticles(DB)
	if err != nil {
		panic(err)
	}
	env["Articles"] = articles
	IndexTpl.Execute(w, env)

}
func Slug(w http.ResponseWriter, r *http.Request) {
}
