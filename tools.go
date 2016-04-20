package docserv

import (
	"fmt"
	"github.com/gerbenjacobs/docserv/bindata"
	"html/template"
	"strings"
)

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
