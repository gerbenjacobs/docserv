# docserv

docserv is a Golang package that you can integrate into your project to generate a documentation browser.
See [godoc](https://godoc.org/github.com/gerbenjacobs/docserv) for more information.

It has two modes; reading the documents from filesystem or generating static copies with
[go-bindata](https://github.com/jteeuwen/go-bindata)

## Install

`go get -u github.com/gerbenjacobs/docserv`

If you plan on using the static abilities, also make sure to install go-bindata.

`go get -u github.com/jteeuwen/go-bindata/...`

## Usage

**Filesystem mode**

```go
ds := docserv.NewDocServ([]string{"README.md"})
ds.Run()
```

**Static mode**

`./static.sh [filepath]...`

The above script will convert your documents to a go-bindata Go script that will imitate a local filesystem when you compile your original program.

You can copy this script to your project (you can change the package name and output file) and use it. __Make sure to import this package.__

```go
import path/to/static

ds := docserv.NewStaticDocServ(map[string][]byte{"README.md": static.MustAsset("README.md")})
ds.Run()
```

## Config

```go
type DocServConfig struct {
	Endpoint        string   // Endpoint to listen to (default: /docs)
	UseStatic       bool     // Whether to use Static mode (default: false)
	Docs            []string // The filepaths for documentation
	Host            string   // The host part of the HTTP listener (default: "")
	Port            string   // The port part of the HTTP listener (default: 9000)
	UseHighlighting bool     // Whether to use syntax highlighting (default: true)
	SyntaxStyle     string   // The style to use for syntax highlighting (default: darkula)
}
```

## Template

![screenshot](https://raw.githubusercontent.com/gerbenjacobs/docserv/master/screenshot.png)

DocServ comes with a default template in `resources/template.html`. You can change this template if you wish.

It is using `go-bindata` so you can run the `build.sh` script.

### Syntax Highlighting

We are using [highlight.js](https://highlightjs.org) to add syntax highlighting to your documents.

If you want you can set the `SyntaxStyle` in the `DocServeConfig` to any string
that will match any of the styles showcased [here](https://highlightjs.org/static/demo/)
and hosted [here](https://cdnjs.com/libraries/highlight.js) (Note: Lowercase and spaces replaced with dashes)
