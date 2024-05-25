package main

import (
	"html/template"
)

var DATA TemplateData
var INPUT_FILE string
var OUTPUT_PATH string
var SELECTED_THEME string

type TemplateData struct {
	Title       string
	Langs       []string
	Body        template.HTML
	SyntaxTheme string
}
