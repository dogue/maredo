package main

import (
	"html/template"
)

var data TemplData
var sourceFile string

type TemplData struct {
	Title string
	Langs []string
	Body  template.HTML
	Theme string
}
