package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

var (
	ErrDoesntContainKey = errors.New("query doesn't contain the key")
)

type API struct {
	httpServer *http.Server

	bstInserter Inserter
	bstFinder   Finder
	bstRemover  Remover

	middlewares []Middleware
}

func New(bst BinarySearchTree, addr string) *API {
	return &API{
		httpServer: &http.Server{
			Addr: addr,
		},
		bstFinder:   bst,
		bstInserter: bst,
		bstRemover:  bst,
		middlewares: make([]Middleware, 0),
	}
}

func (api *API) ListenAndServe() error {

	api.initHandlers()

	return nil
}

func (api *API) initHandlers() {
	router := httprouter.New()

	router.GET("/search", handledWith(api.searchHandler, api.middlewares...))
	router.DELETE("/delete", handledWith(api.deleteHandler, api.middlewares...))
	router.POST("/insert", handledWith(api.insertHandler, api.middlewares...))

	// todo check api.httpServer is nil
	api.httpServer.Handler = router
}

func (api *API) Shutdown(ctx context.Context) error {

	return api.httpServer.Shutdown(ctx)
}

func (api *API) searchHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	val, err := getIntFromRequest(r, "val")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	node := api.bstFinder.Find(val)

	if node == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(node.Value())
}

func getIntFromRequest(r *http.Request, key string) (int, error) {
	val := r.URL.Query().Get(key)

	if val == "" {
		return 0, errors.Wrap(ErrDoesntContainKey, key)
	}

	return strconv.Atoi(val)
}

func (api *API) deleteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	val, err := getIntFromRequest(r, "val")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.bstRemover.Remove(val)

	w.WriteHeader(http.StatusOK)
}

func (api *API) insertHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	value, err := getValueFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.bstInserter.Insert(value)

	w.WriteHeader(http.StatusOK)
}

func getValueFromRequest(r *http.Request) (int, error) {
	var valDto insertValueDto
	if err := json.NewDecoder(r.Body).Decode(&valDto); err != nil {
		return 0, err
	}

	return valDto.IntValue, nil
}

type insertValueDto struct {
	IntValue int `json:"value"`
}
