package main

import (
	"log"
	"net/http"

	"github.com/LeoCBS/garden/parameter"
	"github.com/LeoCBS/garden/server"
	"github.com/LeoCBS/garden/storage"
)

func main() {
	// TODO get param from cli ou service discovery
	s, err := storage.NewStorage("mongodb://localhost:27017", "garden", "parameters")
	if err != nil {
		panic(err)
	}
	p := parameter.NewParameter(s)
	mux := server.NewServer(p)
	sr := &http.Server{
		Addr:    ":8080",
		Handler: mux.ServeMux,
	}
	log.Fatal(sr.ListenAndServe())
}
