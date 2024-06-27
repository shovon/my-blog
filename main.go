package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sus/adapters/followersadapter"
	"sus/config"
	"sus/logger"
	"sus/rsahelpers"
	"sus/services/rsaservice"
	"sus/services/sqliteservice"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
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
				// TODO: add a shared inbox, for certain Mastodon implementations that
				//   expect it.
			}
			w.Header().Add("Content-Type", "application/activity+json")
			json.NewEncoder(w).Encode(j)
		})),
		All(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to my site!"))
		})),
	}))

	sqliteservice.GetWriteableDB()

	router.Post("/inbox", ToHandlerFunc(Inbox(
		followersadapter.NewFollowersManagerSQLite(
			sqliteservice.GetWriteableDB(),
			sqliteservice.GetReadableDB(),
		),
	)))
	router.Post("/outbox", NotImplemented("Outbox"))

	router.Get("/followers", NotImplemented("Followers"))
	router.Get("/following", NotImplemented("Following"))
	router.Get("/likes", NotImplemented("Likes"))

	router.Route("/actor", func(r chi.Router) {
		r.Get("/", ToHandlerFunc(ExactAcceptHandler{
			ActivityStreams20(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// TODO: use a separate rsaservice for this.
				key, err := rsaservice.GetKey()

				if err != nil || key == nil {
					w.WriteHeader(500)
					w.Write([]byte("Error occurred"))
					return
				}

				str, err := rsahelpers.PublicKeyToPKIXString(&key.PublicKey)
				if err != nil {
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
					"preferredUsername": "siteactor",
					"publicKey": map[string]any{
						"id":           getURL("#main-key"),
						"owner":        getURL(""),
						"publicKeyPem": str,
					},
				}
				w.Header().Add("Content-Type", "application/activity+json")
				json.NewEncoder(w).Encode(j)
			})),
		}))
	})

	addr := ":8081"

	logger.Logger().Info("Server listening", zap.String("address", addr))
	fmt.Printf("Server listening on %s\n", addr)
	http.ListenAndServe(addr, router)
}
