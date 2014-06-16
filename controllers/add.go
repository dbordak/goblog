package controllers

import (
	"appengine/datastore"

	"models"
)

type AddController struct {
	FormController
}

func (this *AddController) getForm() *formRequest {
	selkey, _ := datastore.DecodeKey(this.GetString("sel"))
	return &formRequest{
		Name:      this.GetString("name"),
		Textarea:  this.GetString("ta"),
		SelectKey: selkey,
	}
}

func (this *AddController) AddCat() {
	this.Data["Title"] = "Add Category"
	this.Data["Form"] = &form{
		Name: "Name",
		Select: &selectList{
			Name:   "Parent",
			DefOpt: true,
			Items:  this.Data["Sidebar"].(*Sidebar).Categories,
		},
	}
}

func (this *AddController) PostCat() {
	fReq := this.getForm()
	cat := &models.Category{fReq.Name}

	key := datastore.NewIncompleteKey(this.AppEngineCtx, "Category", fReq.SelectKey)
	_, err := datastore.Put(this.AppEngineCtx, key, cat)
	if err != nil {
		//TODO: Error
		return
	}
}

func (this *AddController) AddEnt() {
	this.Data["Title"] = "Add Entry"
	this.Data["Form"] = &form{
		Name:     "Title",
		Textarea: "Content",
		Select: &selectList{
			Name:   "Category",
			DefOpt: true,
			Items:  this.Data["Sidebar"].(*Sidebar).Categories,
		},
	}

	//This sets the page to use TinyMCE for Textareas
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "tinymce_head.tpl"
}

func (this *AddController) PostEnt() {
	fReq := this.getForm()
	ent := models.NewEntry(fReq.Name, fReq.Textarea)

	key := datastore.NewIncompleteKey(this.AppEngineCtx, "Entry", fReq.SelectKey)
	_, err := datastore.Put(this.AppEngineCtx, key, ent)
	if err != nil {
		//TODO: Error
		return
	}
}
