package controllers

import (
	"appengine/datastore"
	"github.com/astaxie/beegae"

	"models"
)

type HtmlHead struct {
	Title string
	Head  string
}

type Sidebar struct {
	Categories []models.Category
	IsAdmin    bool
}

type Footer struct {
	PrevButton string
	NextButton string
}

//Master struct for all pages; fills in sidebar and basic info
type BlogController struct {
	beegae.Controller
}

//TODO: Add admin check
func (this *BlogController) Prepare() {
	q := datastore.NewQuery("Category").Order("Name")
	var cats []models.Category
	_, err := q.GetAll(this.AppEngineCtx, &cats)
	if err != nil {
		this.AppEngineCtx.Errorf("fetching categories: %v", err)
		return
	}

	//Sets Category dropdown and "Admin Page" button
	this.Data["Sidebar"] = &Sidebar{cats, true}
	//Sets Newer and Older buttons
	//this.Data["Footer"] = &Footer{"", ""}
	//Sets page title and adds additional scripts
	//this.Data["HtmlHead"] = &HtmlHead{"", ""}

	this.Layout = "layout.html"
}

//Struct for all list pages; fills in looped content and next/prev buttons
type ListController struct {
	BlogController
}

func (this *ListController) Get() {
	this.TplNames = "index.tpl"
	q := datastore.NewQuery("Entry").Order("Date")
	var ents []models.Entry
	_, err := q.GetAll(this.AppEngineCtx, &ents)
	if err != nil {
		this.AppEngineCtx.Errorf("fetching entries: %v", err)
		return
	}

	this.Data["Entries"] = ents
}

type SimpleController struct {
	BlogController
}

func NewSimpleController(template string) *SimpleController {
	sc := &SimpleController{}
	sc.TplNames = "about.tpl"
	//sc.Get = renderSimpleTemplate
	return sc
}

func (this *SimpleController) Get() {
	this.Data["Footer"] = &Footer{this.TplNames, "dwadaw"}
	this.TplNames = "about.tpl"
}

func (this *SimpleController) About() {
	this.TplNames = "about.tpl"
}

type FormController struct {
	BlogController
}
