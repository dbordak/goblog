package controllers

import (
	"appengine/datastore"

	"models"
	"util"
)

type DelController struct {
	FormController
}

func (this *DelController) DelCat() {
	this.Data["Title"] = "Delete Category"
	this.Data["Form"] = &form{
		Select: &selectList{
			Name:  "Category",
			Items: this.Data["Sidebar"].(*sidebar).Categories,
		},
	}
}

func (this *DelController) DelEnt() {
	q := datastore.NewQuery("Entry")
	t := q.Run(this.AppEngineCtx)
	moop := make(map[string]string)
	for {
		var ent models.Entry
		key, err := t.Next(&ent)
		this.AppEngineCtx.Errorf("%v", err)
		if util.QueryErrHandler(err, this.AppEngineCtx, "Entry") {
			break
		}
		moop[ent.Title] = key.Encode()
	}

	this.Data["Title"] = "Delete Entry"
	this.Data["Form"] = &form{
		Select: &selectList{
			Name:  "Entry",
			Items: moop,
		},
	}
}

func (this *DelController) Delete() {
	this.getForm()
	datastore.Delete(this.AppEngineCtx, this.FReq.SelectKey)
}
