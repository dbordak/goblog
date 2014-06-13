package controllers

// import (
// 	"models"
// )

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
	//Items
}

func (this *FormController) Prepare() {
	this.BlogController.Prepare()
	this.TplNames = "form.tpl"
}

func (this *FormController) AddEnt() {
	this.Data["Title"] = "A Form"
	this.Data["Form"] = &form{
		Name: "Title",
		//Select: "Category",
		Textarea: "Content",
	}

	//This sets the page to use TinyMCE for Textareas
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "tinymce_head.tpl"
}
