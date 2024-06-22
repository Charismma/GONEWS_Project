package postgres

import (
	"GoNews_project/pkg/db"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

// Функция конструктор
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

func (s *Storage) Posts(n int) ([]db.Post, error) {
	rows, err := s.db.Query(context.Background(), `SELECT id,title,content,pubtime,link FROM posts ORDER BY pubtime
	LIMIT $1`, n)
	if err != nil {
		return nil, err
	}
	var posts []db.Post
	for rows.Next() {
		var post db.Post
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PubTime,
			&post.Link,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, rows.Err()

}

func (s *Storage) AddPosts(posts []db.Post) error {
	for _, post := range posts {
		_, err := s.db.Exec(context.Background(), `INSERT INTO posts(id,title,content,pubtime,link) VALUES($1,$2,$3,$4,$5)`,
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PubTime,
			&post.Link,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
