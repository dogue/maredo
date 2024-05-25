Maredo (pronounced like Laredo) is a very simple static site generator for rendering single-page project docs.

#### Features:

* defaults to rendering `README.md` if no input file given
* configurable page title which defaults to the current directory from which Maredo is run
* [HighlightJS](https://github.com/highlightjs/highlight.js/) support with configurable syntax theme and language support (see HighlightJS link for theme/language strings)
* basic built-in CSS themes; custom CSS files are also supported
* simplicity

#### Demo:

Maredo renders its own docs. Check out the online manpage [here](https://dogue.github.io/maredo).

#### Non-features:

Maredo is not a fully-featured static site generator. It has a very narrow scope by design. It only supports a single input markdown file and a single output HTML file. My goal is to provide a very simple markdown renderer that makes hosting docs for simpler apps quick and convenient (particularly with Github Pages).

#### Contributing

Issues and PRs are always welcome. If you have a theme you'd like to add as a built-in, please open an issue and include a screenshot.
