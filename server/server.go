package server

import (
	"fmt"
	"github.com/shakesearch/searcher"
	"log"
	"net/http"
	"os"
)

type Server struct {
	searcher *searcher.Searcher
}


func (s * Server)InitRoutes(){
	http.HandleFunc("/search", s.handleSearch())
}

func (s * Server)Init(){

	s.searcher = new(searcher.Searcher)
	s.searcher.Init()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	s.InitRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

