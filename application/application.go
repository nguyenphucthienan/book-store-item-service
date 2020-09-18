package application

import (
	"github.com/gorilla/mux"
	"github.com/nguyenphucthienan/book-store-item-service/client/elasticsearch"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func Start() {
	elasticsearch.Init()
	mapUrls()

	server := &http.Server{
		Addr: "localhost:8081",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
