package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	if err := buildSite(cfg); err != nil {
		panic(err)
	}
}

func buildSite(cfg Config) error {
	fmt.Println("Building site:")
	pageSlugs, err := readPagesDir(cfg)
	if err != nil {
		return err
	}
	fmt.Println("Found pages:", pageSlugs)

	pages := make([]Page, 0, len(pageSlugs))
	menu := make([]MenuItem, 0, len(pageSlugs))

	for _, slug := range pageSlugs {
		meta, body, err := loadPage(cfg, slug)
		if err != nil {
			return err
		}
		if isArchived(meta) {
			fmt.Println("skipped (archived):", slug)
			continue
		}
		page := Page{
			Slug:  slug,
			Body:  body,
			Meta:  meta,
		}
		pages = append(pages, page)

		if isMenu(meta) {
			menu = append(menu, MenuItem{
				Title: pageTitle(meta, slug),
				URL:   pageURL(cfg.BaseURL, slug),
			})
		}
	}

	for _, page := range pages {
		if err := renderPage(cfg, page, menu); err != nil {
			return err
		}
		fmt.Println("rendered:", page.Slug)
	}

	return nil
}

func pageTitle(meta url.Values, slug string) string {
	title := meta.Get("title")
	if title == "" {
		return slug
	}
	return title
}

func pageURL(baseURL, slug string) string {
	base := strings.TrimSuffix(baseURL, "/")
	if slug == "index" || slug == "" {
		if base == "" {
			return "/"
		}
		return base + "/"
	}
	if base == "" {
		return "/" + slug + "/"
	}
	return base + "/" + slug + "/"
}
