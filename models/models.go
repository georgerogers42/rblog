package models
import (
	"time"
	"database/sql"
	_ "github.com/lib/pq"
)

type Author struct {
	Id int
	Pseudonym, Name, EncPass string
}

type Article struct {
	Id int
	Author
	Slug, Title, Contents string
	Posted, Updated time.Time
}

func AllArticles(db *sql.DB) ([]Article, error) {
	rows, err := db.Query("SELECT author.id, pseudonym, name, encpass, article.id, slug, title, contents, posted, updated FROM author, article WHERE author.id = article.fk_author ORDER BY posted DESC")
	if err != err {
		return nil, err
	}
	ret := []Article{}
	for rows.Next() {
		v := Article{}
		rows.Scan(&v.Author.Id, &v.Author.Pseudonym, &v.Author.Name, &v.Author.EncPass)
		rows.Scan(&v.Slug, &v.Title, &v.Posted, &v.Updated)
		ret = append(ret, v)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
