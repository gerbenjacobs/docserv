package docserv

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	endpoint        = "/docs"
	parsedTemplate  *template.Template
	identifierRegex = regexp.MustCompile("[^A-Za-z0-9]+")
)

type DocServConfig struct {
	Docs               []string
	Host               string
	Port               string
	SyntaxHighlighting bool
	SyntaxStyle        string
}

type DocServ struct {
	Config   DocServConfig
	docFiles map[string]template.HTML
}

func NewDocServ(docs []string) *DocServ {
	d := DocServ{
		Config: DocServConfig{
			Docs:               docs,
			Host:               "",
			Port:               "9000",
			SyntaxHighlighting: true,
			SyntaxStyle:        "darkula",
		},
		docFiles: map[string]template.HTML{},
	}
	return &d
}

func (d *DocServ) Run() error {
	// get template
	if err := getTemplate(); err != nil {
		return err
	}

	// validate docs
	for _, doc := range d.Config.Docs {
		c, err := ioutil.ReadFile(doc)
		if err != nil {
			return fmt.Errorf("Failed to read file: %v", err)
		}
		md := blackfriday.MarkdownCommon(c)
		d.docFiles[doc] = template.HTML(md)
	}

	// serve docs
	fmt.Printf("Serving your docs at %s%s", d.getAddress(), endpoint)

	http.HandleFunc(endpoint, d.Handler)
	return http.ListenAndServe(d.getAddress(), nil)
}

func (d *DocServ) Handler(w http.ResponseWriter, r *http.Request) {
	err := parsedTemplate.ExecuteTemplate(w, "docs", d)
	if err != nil {
		fmt.Fprintf(w, "Failed to execute template: %v", err)
	}
}

func (d *DocServ) getAddress() string {
	return fmt.Sprintf("%s:%s", d.Config.Host, d.Config.Port)
}

func (d *DocServ) DocFiles() map[string]template.HTML {
	return d.docFiles
}
