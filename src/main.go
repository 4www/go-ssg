package main

import (
	"fmt"
)

func main() {
	cfg := Config{
		ContentFolder:    "content",
		PagesFolder:      "pages",
		ContentExtension: ".txt",
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
		fmt.Println("loaded:", slug, "meta:", meta, "body:", body)
	}

	return nil
}


