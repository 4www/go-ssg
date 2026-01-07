package main

import (
	"fmt"
)

func main() {
	cfg := Config{
		ContentFolder:    "content",
		PagesFolder:      "pages",
		ContentExtension: ".txt",
		TemplatesFolder:  "templates",
		OutputFolder:     "dist",
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

		if err := renderPage(cfg, page); err != nil {
			return err
		}
		fmt.Println("rendered:", slug)
	}

	return nil
}
