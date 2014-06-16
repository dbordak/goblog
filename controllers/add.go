package controllers

import (
	"appengine/datastore"

	"models"
)

type AddController struct {
	FormController
}

type addForm struct {
	Name     string
	Select   *selectList
	Textarea string
}

type addFormRequest struct {
	Name     string
	Select   *datastore.Key
	Textarea string
}

func (this *AddController) getAddForm() *addFormRequest {
	fReq := &addFormRequest{
		Name:     this.GetString("name"),
		Textarea: this.GetString("ta"),
	}
	sel, _ := this.GetInt("sel")
	if sel != 0 {
		fReq.Select = datastore.NewKey(this.AppEngineCtx, "Category", "", sel, nil)
	}
	return fReq
}

func (this *AddController) AddCat() {
	this.Data["Title"] = "Add Category"
	sl := &selectList{
		Name:   "Parent",
		DefOpt: true,
		Items:  this.Data["Sidebar"].(*Sidebar).Categories,
	}

	this.Data["Form"] = &addForm{
		Name:   "Name",
		Select: sl,
	}
}

func (this *AddController) PostCat() {
	fReq := this.getAddForm("Category")
	cat := &models.Category{fReq.Name, fReq.Select}

	key := datastore.NewIncompleteKey(this.AppEngineCtx, "Category", nil)
	_, err := datastore.Put(this.AppEngineCtx, key, cat)
	if err != nil {
		//TODO: Error
		return
	}
}

func (this *AddController) AddEnt() {
	this.Data["Title"] = "Add Entry"
	sl := &selectList{
		Name:   "Category",
		DefOpt: true,
		Items:  this.Data["Sidebar"].(*Sidebar).Categories,
	}

	this.Data["Form"] = &addForm{
		Name:     "Title",
		Select:   sl,
		Textarea: "Content",
	}

	//This sets the page to use TinyMCE for Textareas
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "tinymce_head.tpl"
}

func (this *AddController) PostEnt() {
	fReq := this.getAddForm("Category")
	ent := models.NewEntry(fReq.Name, fReq.Select, fReq.Textarea)

	key := datastore.NewIncompleteKey(this.AppEngineCtx, "Entry", nil)
	_, err := datastore.Put(this.AppEngineCtx, key, ent)
	if err != nil {
		//TODO: Error
		return
	}
}
