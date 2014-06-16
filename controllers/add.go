package controllers

import (
	"appengine/datastore"

	"models"
)

type AddController struct {
	FormController
}

func (this *AddController) getForm() {
	selkey, err := datastore.DecodeKey(this.GetString("sel"))
	if err != nil {
		this.AppEngineCtx.Errorf("getting request form: %v", err)
		return
	}
	this.FReq = &formRequest{
		Name:      this.GetString("name"),
		Textarea:  this.GetString("ta"),
		SelectKey: selkey,
	}
}

func (this *AddController) putModel(modType string, model interface{}) {
	key := datastore.NewIncompleteKey(this.AppEngineCtx, modType, this.FReq.SelectKey)
	_, err := datastore.Put(this.AppEngineCtx, key, model)
	if err != nil {
		this.AppEngineCtx.Errorf("adding %s: %v", modType, err)
		this.Data["Label"] = "Uh oh"
		return
	}
}

func (this *AddController) AddCat() {
	this.Data["Title"] = "Add Category"
	this.Data["Form"] = &form{
		Name: "Name",
		Select: &selectList{
			Name:   "Parent",
			DefOpt: true,
			Items:  this.Data["Sidebar"].(*sidebar).Categories,
		},
	}
}

func (this *AddController) PostCat() {
	this.getForm()
	cat := &models.Category{this.FReq.Name}
	this.putModel("Category", cat)
}

func (this *AddController) AddEnt() {
	this.Data["Title"] = "Add Entry"
	this.Data["Form"] = &form{
		Name:     "Title",
		Textarea: "Content",
		Select: &selectList{
			Name:   "Category",
			DefOpt: true,
			Items:  this.Data["Sidebar"].(*sidebar).Categories,
		},
	}

	//This sets the page to use TinyMCE for Textareas
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "tinymce_head.tpl"
}

func (this *AddController) PostEnt() {
	this.getForm()
	ent := models.NewEntry(this.FReq.Name, this.FReq.Textarea)
	this.putModel("Entry", ent)
}
