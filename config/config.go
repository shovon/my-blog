package config

import "os"

var host string

func init() {
	h, exists := os.LookupEnv("HOST")
	if !exists {
		panic("")
	}
	host = h
}

func Host() string {
	return host
}
