package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static/*
var fs embed.FS

func main() {
	log.Println("Serving files from embed.FS")
	http.ListenAndServe(
		":8080",
		http.FileServer(http.FS(fs)),
	)
}
