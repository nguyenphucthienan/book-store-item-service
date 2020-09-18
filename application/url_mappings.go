package application

import (
	"github.com/nguyenphucthienan/book-store-item-service/controller"
	"net/http"
)

const (
	apiPrefix = "/api"
)

func mapUrls() {
	router.HandleFunc(apiPrefix+"/items", controller.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc(apiPrefix+"/items/{id}", controller.ItemController.Get).Methods(http.MethodGet)
}
