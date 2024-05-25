package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

// "github.com/adrg/xdg"

const VERSION = "0.0.1"

func main() {
	app := initCli()
	_ = app
	initTemplate()
	html, _ := renderPage(sourceFile, data)
	out := filepath.Join(outputFile, "index.html")
	os.WriteFile(out, html, fs.ModePerm)
	outCss := filepath.Join(outputFile, "style.css")
	os.WriteFile(outCss, []byte(DEFAULT_CSS), fs.ModePerm)
}
