package main

import (
	"net/http"

	"github.com/LeoCBS/garden/parameter"
	"github.com/LeoCBS/garden/storage"
)

func main() {
	// TODO get param from cli ou service discovery
	s := storage.NewStorer("mongodb://localhost:27017", "garden", "parameters")
	p := parameter.NewParameter(s)
	mux := server.NewSever(p)
	http.ListenAndServe(":3000", mux)
}
