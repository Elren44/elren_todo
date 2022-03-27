package utils

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/Elren44/elren_todo/internal/config"
	"github.com/Elren44/elren_todo/internal/model"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var appTemplates *config.Config

func NewAppTemplates(a *config.Config) {
	appTemplates = a
}

func AddDefaultData(td *model.TemplateData, r *http.Request) *model.TemplateData {
	td.Flash = appTemplates.Session.PopString(r.Context(), "flash")
	td.Error = appTemplates.Session.PopString(r.Context(), "error")
	td.Warning = appTemplates.Session.PopString(r.Context(), "warning")

	td.SCRFToken = nosurf.Token(r)
	return td
}

//RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data *model.TemplateData) error {
	var tc map[string]*template.Template

	if appTemplates.UseCache {
		//get the template cache from app config
		tc = appTemplates.TemplatesCache
	} else {
		tc, _ = TemplateCache()

	}

	t, ok := tc[tmpl]
	if !ok {
		return errors.New("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data, r)

	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		return err
	}
	return nil

}

//TemplateCache creates a template cashe as a map
func TemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
