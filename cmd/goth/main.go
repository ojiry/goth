package main

import (
	"github.com/ojiry/goth/internal/server"
	"log"
)

func main() {
	s := server.NewServer()
	log.Fatal(s.ListenAndServe())
}
