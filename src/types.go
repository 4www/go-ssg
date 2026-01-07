package main

import "net/url"

type Page struct {
	Slug  string
	Body  string
	Meta  url.Values
}
