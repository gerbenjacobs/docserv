package docserv

import (
	"fmt"
	"html/template"
	"strings"
)

func getTemplate() error {
	h, err := Asset("etc/doc.html")
	if err != nil {
		return fmt.Errorf("Failed to read HTML: %v\n", err)
	}

	f := template.FuncMap{
		"ident": parseAsIdentifier,
	}

	parsedTemplate, err = template.New("docs").Funcs(f).Parse(string(h))
	if err != nil {
		return fmt.Errorf("Failed to parse HTML: %v\n", err)
	}
	return nil
}

func parseAsIdentifier(input string) string {
	r := identifierRegex.ReplaceAllString(input, "")
	return strings.ToLower("i-" + r)
}
