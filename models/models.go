package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

type Author struct {
	Id                       int
	Pseudonym, Name, EncPass string
}

type Article struct {
	Author
	Id                    int
	Slug, Title, Contents string
	Posted, Updated       time.Time
}

func AllArticles(db *sql.DB) ([]*Article, error) {
	rows, err := db.Query("SELECT author.id, pseudonym, name, encpass, article.id, slug, title, contents, posted, updated FROM author, article WHERE author.id = article.fk_author ORDER BY posted DESC")
	if err != nil {
		return nil, err
	}
	ret := []*Article{}
	for rows.Next() {
		v := Article{}
		rows.Scan(
			&v.Author.Id, &v.Author.Pseudonym, &v.Author.Name, &v.Author.EncPass,
			&v.Id, &v.Slug, &v.Title, &v.Contents, &v.Posted, &v.Updated,
		)
		ret = append(ret, &v)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
