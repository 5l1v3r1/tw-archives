package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func init() {
	router := httprouter.New()
	router.NotFound = HandleNotFound
	router.GET("/", IndexHandler)
	http.Handle("/", router)
}
