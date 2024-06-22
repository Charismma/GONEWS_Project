package db

//Структура публикации
type Post struct {
	ID      int
	Title   string
	Content string
	PubTime int64
	Link    string
}

//Задаем интерфес для работы с базой в памяти, для облегчения разработки
type Interface interface {
	Posts(int) ([]Post, error)
	AddPosts([]Post) error
}
