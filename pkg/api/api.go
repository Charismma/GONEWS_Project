package api

import (
	"GoNews_project/pkg/db"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
	db     db.Interface
	router *mux.Router
}

func New(db db.Interface) *API {
	api := API{
		db: db,
	}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

func (api *API) endpoints() {
	api.router.HandleFunc("/news/{n}", api.postsHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp"))))
}

func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) postsHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["n"]
	n, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post, err := api.db.Posts(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(post)

}
