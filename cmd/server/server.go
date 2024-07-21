package main

import (
	"GoNews_project/pkg/api"
	"GoNews_project/pkg/db"
	"GoNews_project/pkg/db/memdb"
	"GoNews_project/pkg/db/postgres"
	"GoNews_project/pkg/rss"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type jsonConfig struct {
	URLS   []string `json:"rss"`
	Period int      `json:"request_period"`
}

type server struct {
	db  db.Interface
	api *api.API
}

func main() {
	var srv server
	initStringDb := "postgres://postgres:password@192.168.1.191:5432/GoNews"
	db1, err := postgres.New(initStringDb)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Подключение к БД")
	db2, err := memdb.New()
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db1, db2
	srv.db = db1
	srv.api = api.New(srv.db)
	var conf jsonConfig
	log.Println("Чтение файла с сайтами")
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &conf)
	if err != nil {
		log.Fatal(err)
	}
	chPosts := make(chan []db.Post)
	chError := make(chan error)
	for _, url := range conf.URLS {
		go ParseUrls(url, &srv.db, chPosts, chError, conf.Period)
		log.Println("Запуск одной из горутин для парсинга")
	}
	go func() {
		for posts := range chPosts {
			log.Println("Добавление постов в базу")
			srv.db.AddPosts(posts)
		}
	}()
	go func() {
		for err := range chError {
			log.Println(err)
		}
	}()
	log.Println("Запуск сервера")
	err = http.ListenAndServe(":80", srv.api.Router())
	if err != nil {
		log.Fatal(err)
	}
}

func ParseUrls(url string, db *db.Interface, posts chan<- []db.Post, errors chan<- error, period int) {
	log.Println("Внутри функции парсинга начало")
	ticker := time.NewTicker(time.Minute * time.Duration(period))
	defer ticker.Stop()
	for range ticker.C {
		news, err := rss.ParseRss(url)
		if err != nil {
			errors <- err
			continue
		}
		posts <- news
	}
	log.Println("Внутри функции парсинга конец")
}
