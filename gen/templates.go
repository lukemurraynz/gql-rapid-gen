// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package gen

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"strings"
	"text/template"
)

// templateCache stores all available templates
var templateCache = make(map[string]*template.Template, 128)

// templateFiles includes core and plugin templates
//
//go:embed templates/*.tmpl plugins/*/templates/*.tmpl
var templateFiles embed.FS

// init loads all template files from the embedded directories
func init() {
	err := fs.WalkDir(templateFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if path == "" || (d != nil && d.IsDir()) {
			return nil
		}
		log.Printf("Loading template: %s", path)
		rawTemplate, err := fs.ReadFile(templateFiles, path)
		if err != nil {
			return fmt.Errorf("failed loading %s: %w", path, err)
		}
		tl, err := template.New(path).Parse(string(rawTemplate))
		if err != nil {
			return fmt.Errorf("failed parsing %s: %w", path, err)
		}
		templateCache[path] = tl
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// MustRegisterTemplate manually registers a new template, if defined outside the normal file structure
func MustRegisterTemplate(plugin Plugin, name string, raw string) {
	key := plugin.Name() + "__" + name
	_, pres := templateCache[key]
	if pres {
		panic("Template " + key + " already registered")
	}
	tmpl, err := template.New(key).Parse(raw)
	if err != nil {
		panic("Failed initializing template " + key + ": " + err.Error())
	}
	templateCache[key] = tmpl
}

func ExecuteTemplate(templateName string, data any) (output string, err error) {
	tmpl, ok := templateCache[templateName]
	if !ok || tmpl == nil {
		return "", fmt.Errorf("invalid template name: %s", templateName)
	}

	outBuf := &strings.Builder{}

	err = tmpl.Execute(outBuf, data)
	if err != nil {
		return "", fmt.Errorf("failed generating template %s: %w", templateName, err)
	}

	return outBuf.String(), nil
}
