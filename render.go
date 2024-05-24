package main

import (
	"bytes"
	bf "github.com/russross/blackfriday/v2"
	"html/template"
	"os"
)

func renderPage(filename string, templData TemplData) (out string, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	html := bf.Run(file)
	templData.Body = template.HTML(html)
	t := template.New("main")
	t.Parse(RENDER_TEMPL)
	buf := bytes.Buffer{}
	t.Execute(&buf, templData)
	return buf.String(), nil
}

const (
	RENDER_TEMPL = `<!DOCTYPE html>
<html>

<head>
	<title>{{ .Title }}</title>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/github-dark.css">
	<!-- <link rel="stylesheet" href="./tomorrow-night-eighties.css"> -->
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js"></script>
	{{ range .Langs -}}
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/languages/{{ . }}.min.js"></script>
	{{ end }}
</head>

<body>
	{{ .Body }}
	<script>hljs.highlightAll();</script>
</body>

</html>
`
)
