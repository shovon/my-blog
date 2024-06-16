package main

import (
	"io"
	"net/http"
	"sus/ports/followersport"
	"sus/prettyld"
)

func Inbox(followersAdapter followersport.FollowersManager) http.Handler {
	return ExactAcceptHandler{
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
	}
}
