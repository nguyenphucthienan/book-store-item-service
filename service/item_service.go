package service

import (
	"github.com/nguyenphucthienan/book-store-item-service/domain/item"
	"github.com/nguyenphucthienan/book-store-utils-go/errors"
	"net/http"
)

var (
	ItemService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(item.Item) (*item.Item, errors.RestError)
	Get(string) (*item.Item, errors.RestError)
}

type itemService struct{}

func (s *itemService) Create(item.Item) (*item.Item, errors.RestError) {
	return nil, errors.NewRestError(
		"Not implemented",
		http.StatusNotImplemented,
		"not_implemented",
		nil,
	)
}

func (s *itemService) Get(string) (*item.Item, errors.RestError) {
	return nil, errors.NewRestError(
		"Not implemented",
		http.StatusNotImplemented,
		"not_implemented",
		nil,
	)
}
