# !Work In Progress!

# docserv

docserv is a Golang package that you can integrate into your project to generate a documentation browser.

## Install

`go get -u github.com/gerbenjacobs/docserv`

## Usage

```go
ds := docserv.NewDocServ([]string{"README.md"})
ds.Run()
```

## Config

```
type DocServConfig struct {
	Docs               []string
	Host               string
	Port               string
	SyntaxHighlighting bool
	SyntaxStyle        string
}
```

### Syntax Highlighting

We are using [highlight.js](https://highlightjs.org) to add syntax highlighting to your documents.

If you want you can set the `SyntaxStyle` in the `DocServeConfig` to any string
that will match any of the styles showcased [here](https://highlightjs.org/static/demo/)
and hosted [here](https://cdnjs.com/libraries/highlight.js) (Lowercase and spaces replaced with dashes)