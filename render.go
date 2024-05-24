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
	t := template.Must(template.ParseFiles("template.html"))
	buf := bytes.Buffer{}
	t.Execute(&buf, templData)
	return buf.String(), nil
}
