# NAME

maredo - render a markdown file to static HTML

# SYNOPSIS

maredo

```
[--help|-h]
[--input|-i]=[value]
[--lang|-l]=[value]
[--list-themes]
[--output|-o]=[value]
[--page-theme]=[value]
[--syntax-theme]=[value]
[--title|-t]=[value]
[--version|-v]
```

**Usage**:

```
maredo [options]
```

# GLOBAL OPTIONS

**\--help, -h**: show help

**\--input, -i FILE**: render markdown `FILE` to HTML (default: "README.md")

**\--lang, -l LANG**: include syntax highlighting stylesheet for `LANG` (see Highlight.js for available languages)

**\--list-themes**: list built-in CSS themes

**\--output, -o PATH**: save output to `PATH` (default: "docs")

**\--page-theme [BUILT_IN_THEME | CSS_FILE]**: set CSS theme (see \--list-themes) (default: "default")

**\--syntax-theme HJS_THEME**: set syntax highlighting theme (see Highlight.js for available themes) (default: "github-dark")

**\--title, -t TITLE**: set output page title to `TITLE` (default: "maredo")

**\--version, -v**: print the version


# COMMANDS

## help, h

Shows a list of commands or help for one command

