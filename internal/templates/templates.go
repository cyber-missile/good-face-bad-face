package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/cyber-missile/good-face-bad-face/web"
)

type TemplateError struct {
	TemplateName string
	err          error
}

func (e TemplateError) Error() string {
	// I am quite dissatisfied with this. Actually `TemplateError.err` should be used in the f-string with a %w,
	// but `fmt.Sprintf` does not allow to use %w.
	return fmt.Sprintf("error rendering template %s: %s", e.TemplateName, e.err.Error())
}

func (e TemplateError) Unwrap() error {
	return e.err
}

type templateCache map[string]*template.Template

type Templates struct {
	cache templateCache
}

func New() (Templates, error) {
	pagesDir, err := web.GetAllTemplateFiles()
	if err != nil {
		return Templates{}, err
	}

	cache, err := buildCache(pagesDir)
	if err != nil {
		return Templates{}, err
	}

	return Templates{cache: *cache}, nil
}

func (t *Templates) RenderTemplate(w io.Writer, name string, data any) error {
	pageTemplate, ok := t.cache[name]
	if !ok {
		return fmt.Errorf("the template %s does not exist", name)
	}

	err := renderTemplate(pageTemplate, w, name, data)
	if err != nil {
		return TemplateError{err: err, TemplateName: name}
	}

	return nil
}

func buildCache(templateDirEntries *[]fs.DirEntry) (*templateCache, error) {
	var cache templateCache = templateCache{}

	for _, entry := range *templateDirEntries {
		if !entry.Type().IsRegular() {
			continue
		}

		fileName := filepath.Base(entry.Name())
		if !strings.HasSuffix(entry.Name(), ".tmpl") {
			continue
		}

		renderedTemplate, err := initTemplate(entry)
		if err != nil {
			return &cache, fmt.Errorf("error rendering %s: %w", entry.Name(), err)
		}

		cache[fileName] = renderedTemplate
	}

	return &cache, nil
}

// initTemplate create a template with all necessary files for a page.
func initTemplate(entry fs.DirEntry) (*template.Template, error) {
	files := []string{
		"template/base.tmpl",
		fmt.Sprintf("template/pages/%s", entry.Name()),
	}

	ts, err := template.New(entry.Name()).
		Funcs(sprig.FuncMap()).
		ParseFS(web.TemplatesDir, files...)

	if err != nil {
		return nil, fmt.Errorf("building template failed: %w", err)
	}

	return ts, nil
}

func renderTemplate(tmpl *template.Template, w io.Writer, name string, data any) error {
	var buf bytes.Buffer = bytes.Buffer{}

	err := tmpl.ExecuteTemplate(&buf, name, data)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error executing template: %w", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error writing template buffer to writer: %w", err)
	}

	return nil
}
