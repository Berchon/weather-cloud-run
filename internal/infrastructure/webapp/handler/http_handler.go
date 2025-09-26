package handler

import "net/http"

type HttpHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}
