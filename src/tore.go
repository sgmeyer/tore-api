package main

import "time"

// Tore (ToRe) represents an article, blog, or other internet resource to be read
// or viewed at a future point.  ToRe: TO REad.
type Tore struct {
	URL       string    `json:"url"`
	Completed bool      `json:"completed"`
	Time      time.Time `json:"time"`
}

// Tores (TOREs) is a list of all the unread or unviewed resouces.  TOREs TO REads.
type Tores []Tore
