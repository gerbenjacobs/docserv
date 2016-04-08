package docserv

import (
	"fmt"
	"html/template"
)

func getTemplate() error {
	h, err := Asset("etc/doc.html")
	if err != nil {
		return fmt.Errorf("Failed to read HTML: %v\n", err)
	}
	parsedTemplate, err = template.New("docs").Parse(string(h))
	if err != nil {
		return fmt.Errorf("Failed to parse HTML: %v\n", err)
	}
	return nil
}
