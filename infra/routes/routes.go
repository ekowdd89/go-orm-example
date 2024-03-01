package routes

import "github.com/gorilla/mux"

type RouterHandler func(*mux.Router) *mux.Router

func NewRouter() RouterHandler {
	return func(r *mux.Router) *mux.Router {
		return r
	}
}
