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

//go:embed themes/default.css
var DEFAULT_CSS string

func initTemplate() error {
	compiledTempl = template.New("maredo")
	_, err := compiledTempl.Parse(TEMPLATE)
	return err
}

func renderPage(filename string, templData TemplData) (out []byte, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	html := bf.Run(file)
	templData.Body = template.HTML(html)
	buf := bytes.Buffer{}
	compiledTempl.Execute(&buf, templData)
	return buf.Bytes(), nil
}

// func renderDir(dir string, templData TemplData)
