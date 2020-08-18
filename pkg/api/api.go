package api

import (
	"binarytree/pkg/tree/binary"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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
}

type Finder interface {
	Find(int) *binary.TreeNode
}

type Inserter interface {
	Insert(int)
}

type Remover interface {
	Remove(int)
}

func (api *API) ListenAndServe() error {

	api.initHandlers()

	return nil
}

func (api *API) initHandlers() {

}

func (api *API) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	return api.httpServer.Shutdown(ctx)
}

func (api *API) searchHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (api *API) deleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	val, err := getIntFromRequest(r, "val")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.bstRemover.Remove(val)

	w.WriteHeader(http.StatusOK)
}

func (api *API) insertHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
