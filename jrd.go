package main

import "time"

type Link struct {
	Rel        string             `json:"rel"`
	Href       string             `json:"href"`
	Type       *string            `json:"type,omitempty"`
	Titles     map[string]string  `json:"titles,omitempty"`
	Properties map[string]*string `json:"properties,omitempty"`
}

type JRD struct {
	Expires    *time.Time         `json:"expires,omitempty"`
	Subject    *string            `json:"subject,omitempty"`
	Aliases    []string           `json:"aliases,omitempty"`
	Properties map[string]*string `json:"properties,omitempty"`
	Links      []Link             `json:"links,omitempty"`
}
