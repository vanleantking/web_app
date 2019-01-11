package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"../route"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func EditHandler(router *route.Router, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(router.Response, "edit", p)
}

func ViewHandler(router *route.Router, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(router.Response, router.Request, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(router.Response, "view", p)
}

func SaveHandler(router *route.Router, title string) {
	body := router.Request.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(router.Response, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(router.Response, router.Request, "/view/"+title, http.StatusFound)
}
