package docserv

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/gerbenjacobs/docserv/bindata"
)

var identifierRegex = regexp.MustCompile("[^A-Za-z0-9]+")

func parseTemplate() error {
	h, err := bindata.Asset("resources/template.html")
	if err != nil {
		return fmt.Errorf("Failed to read HTML: %v\n", err)
	}

	f := template.FuncMap{
		"ident": parseAsIdentifier,
	}

	parsedTemplate, err = template.New(TEMPLATE).Funcs(f).Parse(string(h))
	if err != nil {
		return fmt.Errorf("Failed to parse HTML: %v\n", err)
	}
	return nil
}

func parseAsIdentifier(input string) string {
	r := identifierRegex.ReplaceAllString(input, "")
	return strings.ToLower("i-" + r)
}
