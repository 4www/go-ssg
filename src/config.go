package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	ContentFolder    string `json:"contentFolder"`
	PagesFolder      string `json:"pagesFolder"`
	ContentExtension string `json:"contentExtension"`
	TemplatesFolder  string `json:"templatesFolder"`
	OutputFolder     string `json:"outputFolder"`
	BaseURL          string `json:"baseURL"`
	SiteTitle        string `json:"siteTitle"`
}

func DefaultConfig() Config {
	return Config{
		ContentFolder:    "content",
		PagesFolder:      "pages",
		ContentExtension: ".txt",
		TemplatesFolder:  "templates",
		OutputFolder:     "dist",
		BaseURL:          "",
		SiteTitle:        "",
	}
}

func LoadConfig(path string) (Config, error) {
	cfg := DefaultConfig()
	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	if err := json.Unmarshal(b, &cfg); err != nil {
		return Config{}, err
	}
	if cfg.ContentFolder == "" {
		cfg.ContentFolder = "content"
	}
	if cfg.PagesFolder == "" {
		cfg.PagesFolder = "pages"
	}
	if cfg.ContentExtension == "" {
		cfg.ContentExtension = ".txt"
	}
	if cfg.TemplatesFolder == "" {
		cfg.TemplatesFolder = "templates"
	}
	if cfg.OutputFolder == "" {
		cfg.OutputFolder = "dist"
	}
	return cfg, nil
}
