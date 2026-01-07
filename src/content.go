package main

import (
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func readPagesDir(cfg Config) ([]string, error) {
	dirPath := filepath.Join(cfg.ContentFolder, cfg.PagesFolder)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

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

func loadPage(cfg Config, slug string) (meta url.Values, body string, err error) {
	raw, err := readPage(cfg, slug)
	if err != nil {
		return nil, "", err
	}
	meta, body, err = parseTxt(raw)
	if err != nil {
		return nil, "", err
	}
	return meta, body, nil
}

func readPage(cfg Config, slug string) (string, error) {
	p := filepath.Join(cfg.ContentFolder, cfg.PagesFolder, slug+cfg.ContentExtension)
	b, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func parseTxt(raw string) (meta url.Values, body string, err error) {
	lines := strings.Split(raw, "\n")

	var parts []string
	i := 0

	for i < len(lines) {
		line := strings.TrimRight(lines[i], "\r") // handle CRLF
		if strings.HasPrefix(line, "&") {
			frag := strings.TrimSpace(line[1:])
			if frag == "" {
				return nil, "", errors.New("empty metadata line")
			}
			parts = append(parts, frag)
			i++
			continue
		}
		break
	}

	meta = url.Values{}
	if len(parts) > 0 {
		query := strings.Join(parts, "&")
		v, perr := url.ParseQuery(query)
		if perr != nil {
			return nil, "", perr
		}
		meta = v
	}

	body = strings.Join(lines[i:], "\n")
	return meta, body, nil
}

func isArchived(meta url.Values) bool {
	value := strings.TrimSpace(strings.ToLower(meta.Get("isArchived")))
	return value == "true" || value == "1" || value == "yes"
}

func isMenu(meta url.Values) bool {
	value := strings.TrimSpace(strings.ToLower(meta.Get("isMenu")))
	if value == "" {
		return true
	}
	return !(value == "false" || value == "0" || value == "no")
}
