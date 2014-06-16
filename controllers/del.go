package controllers

import (
	"appengine/datastore"

	"models"
)

type DelController struct {
	FormController
}

func (this *DelController) getForm(selectType string) *formRequest {
	sel, _ := this.GetInt("sel")
	return &formRequest{
		Select: datastore.NewKey(this.AppEngineCtx, selectType, "", sel, nil),
	}
}

func (this *DelController) GetDelCat() {
	this.Data["Title"] = "Remove Category"
	this.Data["Form"] = &form{
		Select: &selectList{
			Name:  "Category",
			Items: this.Data["Sidebar"].(*Sidebar).Categories,
		},
	}
}

func (this *DelController) PostDelCat() {
	fReq := this.getForm("Category")
	datastore.Delete(this.AppEngineCtx, fReq.Select)
}

func (this *DelController) GetDelEnt() {
	q := datastore.NewQuery("Entry").Order("d")
	t := q.Run(this.AppEngineCtx)
	m := make(map[int64]string)
	for {
		var ent models.Entry
		key, err := t.Next(&ent)
		if err == datastore.Done {
			break
		}
		if err != nil {
			this.AppEngineCtx.Errorf("fetching next Entry: %v", err)
			break
		}
		m[key.IntID()] = ent.Title
	}

	this.Data["Title"] = "Remove Entry"
	this.Data["Form"] = &form{
		Select: &selectList{
			Name:  "Entry",
			Items: m,
		},
	}
}

func (this *DelController) PostDelEnt() {
	fReq := this.getForm("Entry")
	datastore.Delete(this.AppEngineCtx, fReq.Select)
}
