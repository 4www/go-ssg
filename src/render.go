package main

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func renderPage(cfg Config, page Page, menu []MenuItem) error {
	templatePath := filepath.Join(cfg.TemplatesFolder, "layout.html")
	funcs := template.FuncMap{
		"splitParagraphs": splitParagraphs,
	}
	tpl, err := template.New("layout.html").Funcs(funcs).ParseFiles(templatePath)
	if err != nil {
		return err
	}

	outputPath := outputHTMLPath(cfg, page.Slug)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	data := RenderData{
		Page: page,
		Menu: menu,
		SiteTitle: cfg.SiteTitle,
	}
	return tpl.Execute(f, data)
}

func outputHTMLPath(cfg Config, slug string) string {
	if slug == "index" || slug == "" {
		return filepath.Join(cfg.OutputFolder, "index.html")
	}
	return filepath.Join(cfg.OutputFolder, slug, "index.html")
}

func splitParagraphs(body string) []string {
	lines := strings.Split(body, "\n")
	paragraphs := make([]string, 0)
	var current []string

	flush := func() {
		if len(current) == 0 {
			return
		}
		paragraphs = append(paragraphs, strings.Join(current, " "))
		current = current[:0]
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			flush()
			continue
		}
		current = append(current, trimmed)
	}
	flush()

	return paragraphs
}
