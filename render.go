package main

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	bf "github.com/russross/blackfriday/v2"
)

//go:embed template.html
var TEMPLATE_FILE string
var TEMPLATE *template.Template

//go:embed themes/*
var THEMES embed.FS

func initTemplate() error {
	TEMPLATE = template.New("maredo")
	_, err := TEMPLATE.Parse(TEMPLATE_FILE)
	return err
}

func renderPage() error {
	file, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		return err
	}

	html := bf.Run(file)
	DATA.Body = template.HTML(html)
	buf := bytes.Buffer{}
	if err = TEMPLATE.Execute(&buf, DATA); err != nil {
		return err
	}

	outFile := filepath.Join(OUTPUT_PATH, "index.html")
	return os.WriteFile(outFile, buf.Bytes(), fs.ModePerm)
}

func exportCSS() (err error) {
	var themeBytes []byte

	if strings.HasSuffix(SELECTED_THEME, ".css") {
		// user-supplied CSS file
		if themeBytes, err = os.ReadFile(SELECTED_THEME); err != nil {
			return err
		}
	} else {
		// built-in theme
		themePath := filepath.Join("themes", SELECTED_THEME+".css")
		if themeBytes, err = THEMES.ReadFile(themePath); err != nil {
			return err
		}
	}

	themeOut := filepath.Join(OUTPUT_PATH, "style.css")
	return os.WriteFile(themeOut, themeBytes, fs.ModePerm)
}
