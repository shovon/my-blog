package main

import (
	"net/http"

	contentnegotiation "gitlab.com/jamietanna/content-negotiation-go"
)

type NegotiatorHandler struct {
	Negotiator contentnegotiation.Negotiator
	Handler    http.Handler
}

type AcceptsHandler []NegotiatorHandler

var _ http.Handler = AcceptsHandler{}

func (c AcceptsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, ch := range c {
		if IsAcceptable(ch.Negotiator, r.Header.Get("Accepts")) {
			ch.Handler.ServeHTTP(w, r)
			break
		}
	}

	w.WriteHeader(http.StatusNotAcceptable)
	w.Write([]byte("Failed to negotiate a valid content type"))
}
