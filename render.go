package main

import (
	"bytes"
	_ "embed"
	bf "github.com/russross/blackfriday/v2"
	"html/template"
	"os"
)

//go:embed template.html
var TEMPLATE string

func renderPage(filename string, templData TemplData) (out string, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	html := bf.Run(file)
	templData.Body = template.HTML(html)
	t := template.New("main")
	t.Parse(TEMPLATE)
	buf := bytes.Buffer{}
	t.Execute(&buf, templData)
	return buf.String(), nil
}
