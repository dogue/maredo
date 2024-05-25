package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func run() {
	// get current dir for default rendered page title
	cwd, _ := os.Getwd()
	cwd = filepath.Base(cwd)

	app := cli.App{
		Name:      "maredo",
		Usage:     "render a markdown file to static HTML",
		Version:   VERSION,
		Copyright: "©️ 2024 dogue <https://github.com/dogue>",
		UsageText: "maredo [options]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "render markdown `FILE` to HTML",
				Value:       "README.md",
				Destination: &INPUT_FILE,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "save output to `PATH`",
				Value:       "docs",
				Destination: &OUTPUT_PATH,
			},
			&cli.StringFlag{
				Name:        "title",
				Aliases:     []string{"t"},
				Usage:       "set output page title to `TITLE`",
				Value:       cwd,
				Destination: &DATA.Title,
			},
			&cli.StringSliceFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Usage:   "include syntax highlighting stylesheet for `LANG` (see Highlight.js for available languages)",
				Action: func(ctx *cli.Context, s []string) error {
					DATA.Langs = s
					return nil
				},
			},
			&cli.StringFlag{
				Name:        "page-theme",
				Usage:       "set CSS theme (see https://github.com/dogue/maredo/themes)",
				Value:       "default",
				Destination: &SELECTED_THEME,
			},
			&cli.StringFlag{
				Name:        "syntax-theme",
				Usage:       "set syntax highlighting theme (see Highlight.js for available themes)",
				Value:       "github-dark",
				Destination: &DATA.SyntaxTheme,
			},
		},
		Action: func(ctx *cli.Context) error {
			if err := initTemplate(); err != nil {
				return err
			}

			if err := renderPage(); err != nil {
				if errors.Is(err, os.ErrNotExist) {
					fmt.Printf("Could not locate file `%s`\n", INPUT_FILE)
					os.Exit(1)
				}

				return err
			}

			return exportCSS()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
