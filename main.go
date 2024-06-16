package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sus/config"
	"sus/rsahelpers"
	"sus/services/rsaservice"

	"sus/prettyld"

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

func getURL(path string) string {
	protocol := "https"
	if !config.IsSecure() {
		protocol = "http"
	}

	return fmt.Sprintf("%s://%s%s", protocol, config.Host(), path)
}

func main() {
	router := chi.NewRouter()

	router.Get("/.well-known/webfinger", func(w http.ResponseWriter, r *http.Request) {

	})

	router.Get("/", ToHandlerFunc(ExactAcceptHandler{
		ActivityStreams20(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key, err := rsaservice.GetKey()
			if err != nil || key == nil {
				w.WriteHeader(500)
				w.Write([]byte("Error occurred"))
				return
			}

			str, err := rsahelpers.PublicKeyToPKIXString(&key.PublicKey)
			if err != nil || key == nil {
				w.WriteHeader(500)
				w.Write([]byte("Error occurred"))
				return
			}

			j := map[string]any{
				"@context": []string{
					"https://www.w3.org/ns/activitystreams",
					"https://w3id.org/security/v1",
				},
				"id":                getURL(""),
				"type":              "Person",
				"inbox":             getURL("/inbox"),
				"outbox":            getURL("/outbox"),
				"following":         getURL("/following"),
				"followers":         getURL("/followers"),
				"liked":             getURL("/liked"),
				"preferredUsername": config.PreferredUsername(),
				"publicKey": map[string]any{
					"id":           getURL("#main-key"),
					"owner":        getURL(""),
					"publicKeyPem": str,
				},
			}
			w.Header().Add("Content-Type", "application/activity+json")
			json.NewEncoder(w).Encode(j)
		})),
		All(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to my site!"))
		})),
	}))

	router.Post("/inbox", ToHandlerFunc(ExactAcceptHandler{
		ActivityStreams20(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Bad request"))
				return
			}

			type Activity struct {
				Type string `json:"@type"`
			}

			var activity Activity

			err = prettyld.Unmarshal(b, &activity, nil)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Bad request"))
				return
			}

			switch activity.Type {
			case "https://www.w3.org/ns/activitystreams#Follow":
				// TODO: handle a follow request
				type Follow struct {
					ID   string   `mapstructure:"@id"`
					Type []string `mapstructure:"@type"`

					// Actor is the one that is being followed
					Actor []prettyld.ID `mapstructure:"https://www.w3.org/ns/activitystreams#actor"`

					// Object is the follower that is requesting to follow.
					Object []prettyld.ID `mapstructure:"https://www.w3.org/ns/activitystreams#object"`
				}

				var followActivity Follow

				err := prettyld.Unmarshal(b, &followActivity, nil)
				if err != nil {
					w.WriteHeader(400)
					w.Write([]byte("Bad request"))
					return
				}

				w.WriteHeader(501)
				w.Write([]byte("Not yet implemented"))

			case "https://www.w3.org/ns/activitystreams#Undo":
				// TODO: handle an undo request, which includes follow requests
			}
		})),
	}))
	router.Post("/outbox", NotImplemented("Outbox"))

	router.Get("/followers", NotImplemented("Followers"))
	router.Get("/following", NotImplemented("Following"))
	router.Get("/likes", NotImplemented("Likes"))

	addr := ":8081"
	fmt.Printf("Server listening on %s\n", addr)
	http.ListenAndServe(addr, router)
}
