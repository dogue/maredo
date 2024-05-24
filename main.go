package main

import (
	// "fmt"
	// "bytes"
	// "fmt"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	// bf "github.com/russross/blackfriday/v2"
	"github.com/urfave/cli/v2"
)

const VERSION = "0.0.1"

var data TemplData
var sourceFile string

type TemplData struct {
	Title string
	Langs []string
	Body  template.HTML
}

func main() {
	cwd, _ := os.Getwd()
	cwd = filepath.Base(cwd)

	app := cli.App{
		Name:      "maredo",
		Usage:     "render a markdown file to static HTML",
		Version:   VERSION,
		Copyright: "©️ 2024 dogue <https://github.com/dogue>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "render `FILE` to HTML",
				Value:       "README.md",
				Destination: &sourceFile,
			},
			&cli.StringFlag{
				Name:        "title",
				Aliases:     []string{"t"},
				Usage:       "set output page title to `TITLE`",
				Value:       cwd,
				Destination: &data.Title,
			},
			&cli.StringSliceFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Usage:   "set highlighting languages to be included (see Highlight.js for available languages)",
				Action: func(ctx *cli.Context, s []string) error {
					data.Langs = s
					return nil
				},
			},
		},
		// Action: func(ctx *cli.Context) error {
		// 	return nil
		// },
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", data)
	fmt.Println(sourceFile)

	// docs, _ := os.ReadFile("sample.md")
	// out := bf.Run(docs)
	// t := template.Must(template.ParseFiles("template.html"))
	// t.Execute(os.Stdout, Data{Title: "test", Langs: []string{"javascript", "go", "rust"}, Body: template.HTML(out)})
}
