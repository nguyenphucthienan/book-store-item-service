package item

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nguyenphucthienan/book-store-item-service/client/elasticsearch"
	"github.com/nguyenphucthienan/book-store-item-service/domain/query"
	restErrors "github.com/nguyenphucthienan/book-store-utils-go/errors"
	"strings"
)

const (
	itemsIndex = "items"
	itemType   = "_doc"
)

func (i *Item) Save() restErrors.RestError {
	result, err := elasticsearch.Client.Index(itemsIndex, itemType, i)
	if err != nil {
		return restErrors.NewInternalServerError("Error when trying to save item",
			errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() restErrors.RestError {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(itemsIndex, itemType, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return restErrors.NewNotFoundError(fmt.Sprintf("No item found with id %s", i.Id))
		}
		return restErrors.NewInternalServerError(fmt.Sprintf("Error when trying to get id %s", i.Id),
			errors.New("database error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return restErrors.NewInternalServerError("Error when trying to parse database response",
			errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return restErrors.NewInternalServerError("Error when trying to parse database response",
			errors.New("database error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query query.EsQuery) ([]Item, restErrors.RestError) {
	result, err := elasticsearch.Client.Search(itemsIndex, query.Build())
	if err != nil {
		return nil, restErrors.NewInternalServerError("Error when trying to search documents",
			errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, restErrors.NewInternalServerError("Error when trying to parse response",
				errors.New("database error"))
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, restErrors.NewNotFoundError("No items found matching given criteria")
	}
	return items, nil
}
