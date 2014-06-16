package controllers

import (
	"appengine/datastore"
	"github.com/astaxie/beegae"

	"models"
)

type Sidebar struct {
	Categories map[string]string
	IsAdmin    bool
}

type Footer struct {
	PrevButton string
	NextButton string
}

type BlogController struct {
	beegae.Controller
}

//TODO: Add check for IsAdmin
func (this *BlogController) Prepare() {
	q := datastore.NewQuery("Category").Order("n")
	t := q.Run(this.AppEngineCtx)
	m := make(map[string]string)
	for {
		var cat models.Category
		key, err := t.Next(&cat)
		if err == datastore.Done {
			break
		}
		if err != nil {
			this.AppEngineCtx.Errorf("fetching next Category: %v", err)
			break
		}
		m[key.Encode()] = cat.Name
	}

	//Sets Category dropdown and "Admin Page" button
	this.Data["Sidebar"] = &Sidebar{m, true}
	//Sets Newer and Older buttons
	//this.Data["Footer"] = &Footer{"", ""}

	this.Layout = "layout.html"
}

//Struct for all list pages; fills in looped content and next/prev buttons
type ListController struct {
	BlogController
}

func (this *ListController) Get() {
	this.TplNames = "index.tpl"
	q := datastore.NewQuery("Entry").Order("d").Limit(10)
	t := q.Run(this.AppEngineCtx)
	m := make(map[string]string)
	for {
		var ent models.Entry
		_, err := t.Next(&ent)
		if err == datastore.Done {
			break
		}
		if err != nil {
			this.AppEngineCtx.Errorf("fetching next Entry: %v", err)
			break
		}
		m[ent.Title] = ent.Content
	}
	this.Data["Entries"] = m
}

type SimpleController struct {
	BlogController
}

func (this *SimpleController) Prepare() {
	this.BlogController.Prepare()
	this.TplNames = "simple.tpl"
}

func (this *BlogController) About() {
	this.TplNames = "about.tpl"
}

func (this *BlogController) AdminNav() {
	this.TplNames = "admin.tpl"
}

func (this *BlogController) EntryPage() {
	this.TplNames = "simple.tpl"
}
