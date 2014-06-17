package controllers

import (
	"html/template"
	"strconv"

	"appengine/datastore"
	"github.com/astaxie/beegae"

	"models"
	"util"
)

type sidebar struct {
	Categories map[string]string
	IsAdmin    bool
}

type footer struct {
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
	moop := make(map[string]string)
	for {
		var cat models.Category
		key, err := t.Next(&cat)
		if util.QueryErrHandler(err, this.AppEngineCtx, "Category") {
			break
		}
		moop[cat.Name] = key.Encode()
	}

	//Sets Category dropdown and "Admin Page" button
	this.Data["Sidebar"] = &sidebar{moop, true}
	//Sets Newer and Older buttons
	//this.Data["Footer"] = &Footer{"", ""}

	this.Layout = "layout.html"
}

func (this *BlogController) About() {
	this.TplNames = "about.tpl"
}

func (this *BlogController) AdminNav() {
	this.TplNames = "admin.tpl"
}

func (this *BlogController) EntryPage() {
	this.TplNames = "simple.tpl"
	year, err := strconv.Atoi(this.Ctx.Input.Param(":year"))
	if err != nil {
		return
	}
	month, err := strconv.Atoi(this.Ctx.Input.Param(":month"))
	if err != nil {
		return
	}
	title := this.Ctx.Input.Param(":title")
	entkey, err := datastore.DecodeKey(this.Ctx.Input.Param(":entsafe"))
	if err != nil {
		return
	}
	var ent models.Entry
	datastore.Get(this.AppEngineCtx, entkey, &ent)
	if ent.Date.Year() == year && int(ent.Date.Month()) == month && title == ent.Title {
		this.Data["Title"] = ent.Title
		this.Data["Content"] = template.HTML(ent.Content)
	}
}
