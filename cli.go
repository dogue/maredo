package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
)

func run() {
	// get current dir for default rendered page title
	cwd, _ := os.Getwd()
	cwd = filepath.Base(cwd)

	app := &cli.Command{
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
				Action: func(ctx context.Context, c *cli.Command, s []string) error {
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
				Action: func(ctx context.Context, cmd *cli.Command, s string) error {
					DATA.SyntaxTheme = strings.TrimSuffix(s, ".css")
					return nil
				},
			},
			&cli.BoolFlag{
				Name:  "list-themes",
				Usage: "list built-in CSS themes",
				Action: func(ctx context.Context, cmd *cli.Command, b bool) error {
					themeList, err := THEMES.ReadDir("themes")
					if err != nil {
						return err
					}

					fmt.Printf("Available built-in themes:\n\n")

					for _, theme := range themeList {
						t := strings.TrimSuffix(theme.Name(), ".css")
						fmt.Printf("  - %s\n", t)
					}

					os.Exit(0)
					return nil
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
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

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
