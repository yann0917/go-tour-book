package main

import (
	"net/http"
	"time"

	"github.com/yann0917/go-tour-book/blog-service/internal/routers"
)

func main() {
	r := routers.NewRouter()
	s := &http.Server{
		Addr:           "8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
