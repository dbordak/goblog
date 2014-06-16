package controllers

import (
	"appengine/datastore"

	"models"
	"util"
)

type DelController struct {
	FormController
}

func (this *DelController) getForm(selectType string) {
	selkey, err := datastore.DecodeKey(this.GetString("sel"))
	if err != nil {
		this.AppEngineCtx.Errorf("getting request form: %v", err)
		return
	}
	this.FReq = &formRequest{SelectKey: selkey}
}

func (this *DelController) GetDelCat() {
	this.Data["Title"] = "Delete Category"
	this.Data["Form"] = &form{
		Select: &selectList{
			Name:  "Category",
			Items: this.Data["Sidebar"].(*sidebar).Categories,
		},
	}
}

func (this *DelController) PostDelCat() {
	this.getForm("Category")
	datastore.Delete(this.AppEngineCtx, this.FReq.SelectKey)
}

func (this *DelController) GetDelEnt() {
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

func (this *DelController) PostDelEnt() {
	this.getForm("Entry")
	datastore.Delete(this.AppEngineCtx, this.FReq.SelectKey)
}
