package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nguyenphucthienan/book-store-item-service/domain/item"
	"github.com/nguyenphucthienan/book-store-item-service/service"
	"github.com/nguyenphucthienan/book-store-item-service/utils/http_utils"
	"github.com/nguyenphucthienan/book-store-oauth-go/oauth"
	"github.com/nguyenphucthienan/book-store-utils-go/errors"
	"io/ioutil"
	"net/http"
)

var (
	ItemController itemControllerInterface = &itemController{}
)

type itemControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemController struct{}

func (c *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := errors.NewBadRequestError("Invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var newItem item.Item
	if err := json.Unmarshal(requestBody, &newItem); err != nil {
		responseErr := errors.NewBadRequestError("Invalid request body")
		http_utils.RespondError(w, responseErr)
		return
	}

	newItem.Seller = oauth.GetClientId(r)
	createdItem, createErr := service.ItemService.Create(newItem)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, createdItem)
}

func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {
	_, err := service.ItemService.Get(mux.Vars(r)["id"])
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
}
