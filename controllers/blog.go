package controllers

import (
	"appengine/datastore"
	"github.com/astaxie/beegae"

	"models"
)

type BlogController struct {
	beegae.Controller
}

type ListController struct {
	BlogController
}

type SimpleController struct {
	BlogController
}

type FormController struct {
	BlogController
}

type HtmlHead struct {
	Title string
	Head  string
}

type Sidebar struct {
	Categories []string
	IsAdmin    bool
}

type Footer struct {
	PrevButton string
	NextButton string
}

func (this *BlogController) Prepare() {
	q := datastore.NewQuery("Category").Order("Name")
	var cats []models.Category
	keys, err := q.GetAll(this.AppEngineCtx, &cats)
	if err != nil {
		this.AppEngineCtx.Errorf("fetching categories: %v", err)
		return
	}

	//Sets Category dropdown and "Admin Page" button
	this.Data["Sidebar"] = &Sidebar{[]string{"a", "b"}, true}
	//Sets Newer and Older buttons
	this.Data["Footer"] = &Footer{"", ""}
	//Sets page title and adds additional scripts
	this.Data["HtmlHead"] = &HtmlHead{"", ""}

	this.Data["Entries"] = keys
	this.Data["Categories"] = keys

	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
}

func (this *ListController) Get() {
	this.TplNames = "index.tpl"
}
