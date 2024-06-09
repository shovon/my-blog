package main

import contentnegotiation "gitlab.com/jamietanna/content-negotiation-go"

func IsAcceptable(negotiator contentnegotiation.Negotiator, accepts string) bool {
	_, _, err := negotiator.Negotiate(accepts)
	return err == nil
}
