package main

import (
	"bytes"
	_ "embed"
	"html/template"
	"os"

	bf "github.com/russross/blackfriday/v2"
)

//go:embed template.html
var TEMPLATE string
var compiledTempl *template.Template

func initTemplate() error {
	compiledTempl = template.New("maredo")
	_, err := compiledTempl.Parse(TEMPLATE)
	return err
}

func renderPage(filename string, templData TemplData) (out string, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	html := bf.Run(file)
	templData.Body = template.HTML(html)
	buf := bytes.Buffer{}
	compiledTempl.Execute(&buf, templData)
	return buf.String(), nil
}

// func renderDir(dir string, templData TemplData)
