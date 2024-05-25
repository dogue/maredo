package main

import (
	"html/template"
)

var data TemplData
var sourceFile string
var outputFile string

type TemplData struct {
	Title       string
	Langs       []string
	Body        template.HTML
	PageTheme   string
	SyntaxTheme string
}
