package main

import (
	"fmt"
	"net/http"

	contentnegotiation "gitlab.com/jamietanna/content-negotiation-go"
)

type NegotiatorHandler struct {
	Types   []string
	Handler http.Handler
}

// ExactAcceptHandler is an opinionated handler that only handles HTTP reqeuests
// where the
type ExactAcceptHandler []NegotiatorHandler

var _ http.Handler = ExactAcceptHandler{}

func deriveMediaTypeString(mt contentnegotiation.MediaType) string {
	mts := mt.GetType() + "/" + mt.GetSubType()
	return mts
}

func (c ExactAcceptHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, ch := range c {
		negotiator := contentnegotiation.NewNegotiator(ch.Types...)
		_, client, _ := negotiator.Negotiate(r.Header.Get("Accept"))
		for _, el := range ch.Types {
			fmt.Println(deriveMediaTypeString(client), el)
			if deriveMediaTypeString(client) == el {
				ch.Handler.ServeHTTP(w, r)
				return
			}
		}
	}

	w.WriteHeader(http.StatusNotAcceptable)
	w.Write([]byte("Failed to negotiate a valid content type"))
}
