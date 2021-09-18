package main

import "github.com/gorilla/mux"

func (p *Plugin) InitAPI() *mux.Router {
	r := mux.NewRouter()
	return r
}
