package memdb

import "GoNews_project/pkg/db"

type Store struct {
}

//Функция конструткор
func New() *Store {
	return new(Store)
}

//Вывод всех постов
func Posts(n int) ([]db.Post, error) {
	return posts, nil
}

//Добавление постов
func AddPosts([]db.Post) error {
	return nil
}

var posts = []db.Post{
	{
		ID:      1,
		Title:   "Заголовк 1",
		Content: "Содержание 1",
		PubTime: 0,
		Link:    "Ссылка 1",
	},
	{
		ID:      2,
		Title:   "Заголовк 2",
		Content: "Содержание 2",
		PubTime: 0,
		Link:    "Ссылка 2",
	},
}
