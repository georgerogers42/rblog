package models
import (
	"time"
	"database/sql"
	_ "github.com/lib/pq"
)

type Author struct {
	Pseudonym, Name, EncPass string
}

type Article struct {
	Author
	Slug, Title, Contents string
	Posted, Updated *time.Time
}

func AllArticles() []Article {
	
}
