package controllers

import (
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
