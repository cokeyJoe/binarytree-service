package api

import "github.com/julienschmidt/httprouter"

type Middleware func(httprouter.Handle) httprouter.Handle

func handledWith(handler httprouter.Handle, middlewares ...Middleware) httprouter.Handle {

	for _, m := range middlewares {
		handler = m(handler)
	}

	return handler
}

func (api *API) Use(middlewares ...Middleware) {
	api.middlewares = append(api.middlewares, middlewares...)
}
