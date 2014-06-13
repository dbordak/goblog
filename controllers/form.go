package controllers

import (
	"appengine/datastore"

	"models"
)

type FormController struct {
	BlogController
}

type form struct {
	Name     string
	Select   *selectList
	Textarea string
}

type selectList struct {
	Name   string
	DefOpt bool
	Items  map[int64]string
}

func (this *FormController) Prepare() {
	this.BlogController.Prepare()
	this.TplNames = "form.tpl"
}

func (this *FormController) AddCat() {
	this.Data["Title"] = "Add Category"
	sl := &selectList{
		Name:   "Parent",
		DefOpt: true,
		Items:  this.Data["Sidebar"].(*Sidebar).Categories,
	}

	this.Data["Form"] = &form{
		Name:   "Name",
		Select: sl,
	}
}

func (this *FormController) PostCat() {
	name := this.GetString("name")
	pkey, _ := this.GetInt("sel")
	parent := datastore.NewKey(this.AppEngineCtx, "Category", "", pkey, nil)
	cat := models.Category{
		Name:   name,
		Parent: parent,
	}

	key := datastore.NewIncompleteKey(this.AppEngineCtx, "Category", nil)
	_, err := datastore.Put(this.AppEngineCtx, key, &cat)
	if err != nil {
		//TODO: Error
		return
	}
}

func (this *FormController) AddEnt() {
	this.Data["Title"] = "Add Entry"

	this.Data["Form"] = &form{
		Name: "Title",
		//Select: &selectList{"Category", true, }
		Textarea: "Content",
	}

	//This sets the page to use TinyMCE for Textareas
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "tinymce_head.tpl"
}

func (this *FormController) PostEnt() {

}
