package main

import "net/url"

type Page struct {
	Slug  string
	Body  string
	Meta  url.Values
}

type MenuItem struct {
	Title string
	URL   string
}

type RenderData struct {
	Page Page
	Menu []MenuItem
	SiteTitle string
}
