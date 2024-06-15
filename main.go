package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func All(handler http.Handler) NegotiatorHandler {
	return NegotiatorHandler{
		Types:   []string{"text/html", "*/*"},
		Handler: handler,
	}
}

func ActivityStreams20(handler http.Handler) NegotiatorHandler {
	return NegotiatorHandler{
		Types:   []string{"application/ld+json"},
		Handler: handler,
	}
}

func NotImplemented(subject string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprintf("Not yet implemented: %s", subject)))
	}
}

func main() {
	router := chi.NewRouter()

	router.Get("/.well-known/webfinger", NotImplemented("Webfinger endpoint"))

	router.Get("/", ToHandlerFunc(ExactAcceptHandler{
		ActivityStreams20(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			j := map[string]any{
				"@context": []string{
					"https://www.w3.org/ns/activitystreams",
					"https://w3id.org/security/v1",
				},
				"id":                "https://source.example.com",
				"type":              "Person",
				"inbox":             "https://sources.example.com/inbox",
				"outbox":            "https://sources.example.com/outbox",
				"following":         "https://sources.example.com/following",
				"followers":         "https://sources.example.com/followers",
				"liked":             "https://sources.example.com/liked",
				"preferredUsername": "actor",
				// "publicKey": map[string]any{
				// 	"id":           "https://source.example.com#main-key",
				// 	"owner":        "https://source.example.com",
				// 	"publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArLEIhmSM4UXoUbh/UNri\nOmsruokiG4GU0jz7R/rZ3lC0kGEMEJpk7x8hLEtg0DhV9DW3jPOsPi1KvLRkTgiE\nCSEEG+ULqZ3/WTZR3VX+/Tb1huemD2rBZkv9vpL+3qSRuFTvcMumonVuJ6rtT3pG\nTbsXlYmp2n7VkbPQPz6Wy3R7YeGmdNxtRiccwrpeovc+kCCoY/t467cK1ON+FDrq\nT/xgNhG2jPfotMF3ixk5/EQuakKEz2YQP4duD6D86QciZQWjw5YMv96NxV6D24CV\nn8HxEcxM5AfWvqbNLpEvi6UBUVCnM4IzJTlboPBO4tUPSu01YDqb8jbTC0f6rOCZ\nOQIDAQAB\n-----END PUBLIC KEY-----\n",
				// },
			}
			w.Header().Add("Content-Type", "application/activity+json")
			json.NewEncoder(w).Encode(j)
		})),
		All(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to my site!"))
		})),
	}))

	router.Post("/inbox", NotImplemented("Inbox"))
	router.Post("/outbox", NotImplemented("Outbox"))

	router.Get("/followers", NotImplemented("Followers"))
	router.Get("/following", NotImplemented("Following"))
	router.Get("/likes", NotImplemented("Likes"))

	addr := ":8081"
	fmt.Printf("Server listening on %s\n", addr)
	http.ListenAndServe(addr, router)
}
