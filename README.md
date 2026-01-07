This project is to learn more [Go](https://go.dev) by building a minimal,
experimental static site generator.

> Just some personal tests, by no mean supposed to be "useful" to anyone (else).

# What it does
- Reads pages from `content/pages/*.txt`
- Parses metadata lines starting with `&` into key/value pairs
- Renders each page using `templates/layout.html`
- Outputs HTML into `dist/`

# Run locally
```bash
go run ./src 
```

# Content format
Files in `content/pages` use a simple format:

```text
&title=Home page
&isArchived=false
This is the body.

Blank lines create new paragraphs.
```

# Output
- `index.txt` becomes `dist/index.html`
- `about.txt` becomes `dist/about/index.html`

# Notes
- Metadata is available in templates as `{{ .Meta }}`.
- The current body parser splits paragraphs by blank lines.
