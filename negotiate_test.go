package main

import (
	"testing"

	contentnegotiation "gitlab.com/jamietanna/content-negotiation-go"
)

func TestNegotiate(t *testing.T) {
	t.Run("*/* -> any 👍", func(t *testing.T) {
		if !IsAcceptable(contentnegotiation.NewNegotiator("*/*"), "application/json") {
			t.Log("Should be acceptable")
			t.Fail()
		}
	})

	t.Run("application/* -> application/json", func(t *testing.T) {
		if !IsAcceptable(contentnegotiation.NewNegotiator("application/*"), "application/json") {
			t.Log("Should be acceptable")
			t.Fail()
		}
	})

	t.Run("application/*+json -> application/ld+json", func(t *testing.T) {
		if !IsAcceptable(contentnegotiation.NewNegotiator("application/*+json"), "application/ld+json") {
			t.Log("Should be acceptable")
			t.Fail()
		}
	})

	t.Run("application/ld+json -> application/*+json", func(t *testing.T) {
		if !IsAcceptable(contentnegotiation.NewNegotiator("application/ld+json"), "application/*+json") {
			t.Log("Should be acceptable")
			t.Fail()
		}
	})

	t.Run("text/plain -/-> application/json", func(t *testing.T) {
		if IsAcceptable(contentnegotiation.NewNegotiator("text/plain"), "application/json") {
			t.Log("Should be acceptable")
			t.Fail()
		}
	})

	t.Run("application/json -/-> application/ld+json", func(t *testing.T) {
		if IsAcceptable(contentnegotiation.NewNegotiator("application/json"), "application/ld+json") {
			t.Log("Should not be acceptable")
			t.Fail()
		}
	})
}
