package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	contentnegotiation "gitlab.com/jamietanna/content-negotiation-go"
)

func All(handler http.Handler) NegotiatorHandler {
	return NegotiatorHandler{
		Negotiator: contentnegotiation.NewNegotiator("*/*"),
		Handler:    handler,
	}
}

func ActivityStreams20(handler http.Handler) NegotiatorHandler {
	return NegotiatorHandler{
		Negotiator: contentnegotiation.NewNegotiator("application/ld+json"),
		Handler:    handler,
	}
}

func main() {
	router := chi.NewRouter()

	router.Get("/.well-known/webfinger", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Not yet implemented"))
	})

	router.Get("/", ToHandlerFunc(AcceptsHandler{
		ActivityStreams20(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte("Not yet implemnented"))
		})),
		All(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte("Welcome to my site!"))
		})),
	}))

	router.Post("/inbox", ToHandlerFunc(AcceptsHandler{}))
	router.Get("/outbox", ToHandlerFunc(AcceptsHandler{}))

	// router.Get("/followers", ToHandlerFunc(AcceptsHandler{}))
	// router.Get("/following", ToHandlerFunc(AcceptsHandler{}))
	// router.Get("/likes", ToHandlerFunc(AcceptsHandler{}))

	http.ListenAndServe(":8080", router)
}
