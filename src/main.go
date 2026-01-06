package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	ContentFolder string
	PagesFolder string
	ContentExtension string
}

type Page struct {
	Slug string
	Title string
	Body string
}

func main() {
	cfg := Config{
		ContentFolder: "content",
		PagesFolder: "pages",
		ContentExtension: ".txt",
	}
	err := buildSite(cfg)
	if err != nil {
		panic(err)
	}
}

func buildSite(cfg Config) (error) {
	fmt.Println("Building site:")
	pageSlugs, err := readPagesDir(cfg)
	fmt.Println("Found pages:", pageSlugs)

	// @TODO: here should get page content and build each page output?
	for _, e := range pageSlugs {
		page, err := readPage(cfg, e)
		if err != nil {
			panic(err)
		}
		pageContent, err := parseTxt(page)
		if err != nil {
			panic(err)
		}
		fmt.Println(pageContent)
	}

	if err != nil {
		return err
	}
	return nil
}

func readPagesDir(cfg Config) ([]string, error) {
	dirPath := filepath.Join(cfg.ContentFolder, cfg.PagesFolder)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	
	// Allocate a slice with length 0 but capacity equal to the total entries.
	// This avoids extra allocations while still letting us append only the files we keep.
	slugs := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		ext := filepath.Ext(name)
		if ext != cfg.ContentExtension {
			continue
		}
		slug := strings.TrimSuffix(name, ext)
		slugs = append(slugs, slug)
	}
	return slugs, nil
}

func readPage(cfg Config, slug string) (string, error) {
	p := filepath.Join(cfg.ContentFolder, cfg.PagesFolder, slug + cfg.ContentExtension)
	file, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func parseTxt(raw string) (string, error){
	return raw, nil
}
