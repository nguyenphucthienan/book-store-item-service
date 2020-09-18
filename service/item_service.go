package service

import (
	"github.com/nguyenphucthienan/book-store-item-service/domain/item"
	"github.com/nguyenphucthienan/book-store-item-service/domain/query"
	"github.com/nguyenphucthienan/book-store-utils-go/errors"
)

var (
	ItemService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(item.Item) (*item.Item, errors.RestError)
	Get(string) (*item.Item, errors.RestError)
	Search(query.EsQuery) ([]item.Item, errors.RestError)
}

type itemService struct{}

func (s *itemService) Create(item item.Item) (*item.Item, errors.RestError) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(id string) (*item.Item, errors.RestError) {
	returnedItem := item.Item{Id: id}
	if err := returnedItem.Get(); err != nil {
		return nil, err
	}
	return &returnedItem, nil
}

func (s *itemService) Search(query query.EsQuery) ([]item.Item, errors.RestError) {
	dao := item.Item{}
	return dao.Search(query)
}
