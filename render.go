package main

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"

	bf "github.com/russross/blackfriday/v2"
)

//go:embed template.html
var TEMPLATE string
var compiledTempl *template.Template

//go:embed themes/*
var THEMES_FS embed.FS

func initTemplate() error {
	compiledTempl = template.New("maredo")
	_, err := compiledTempl.Parse(TEMPLATE)
	return err
}

func renderPage() error {
	file, err := os.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	html := bf.Run(file)
	data.Body = template.HTML(html)
	buf := bytes.Buffer{}
	if err = compiledTempl.Execute(&buf, data); err != nil {
		return err
	}

	outFile := filepath.Join(outputPath, "index.html")
	return os.WriteFile(outFile, buf.Bytes(), fs.ModePerm)
}

func exportCSS() error {
	themeFile := filepath.Join("themes", data.PageTheme+".css")
	themeData, err := THEMES_FS.ReadFile(themeFile)
	if err != nil {
		return err
	}
	themeOut := filepath.Join(outputPath, "style.css")
	return os.WriteFile(themeOut, themeData, fs.ModePerm)
}
