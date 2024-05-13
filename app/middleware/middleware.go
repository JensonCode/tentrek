package middleware

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(h http.HandlerFunc, mws ...Middleware) http.HandlerFunc {

	for _, mw := range mws {
		h = mw(h)
	}

	return h
}
