package main

import "fmt"

// "github.com/adrg/xdg"

const VERSION = "0.0.1"

func main() {
	app := initCli()
	_ = app
	html, _ := renderPage(sourceFile, data)
	fmt.Printf("%s\n", html)
}
