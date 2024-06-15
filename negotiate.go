package main

import (
	"fmt"
	"strings"

	contentnegotiation "gitlab.com/jamietanna/content-negotiation-go"
)

func IsAcceptable(negotiator contentnegotiation.Negotiator, accept string) bool {
	accept = strings.TrimSpace(accept)
	if accept == "" {
		accept = "*/*"
	}
	fmt.Println(accept)
	server, client, err := negotiator.Negotiate(accept)
	fmt.Println(server)
	fmt.Println(client)
	return err == nil
}
