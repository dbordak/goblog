package controllers

import (
	"appengine/datastore"
)

type FormController struct {
	BlogController
}

func (this *FormController) Prepare() {
	this.BlogController.Prepare()
	this.TplNames = "form.tpl"
}

type selectList struct {
	Name   string
	DefOpt bool
	Items  map[int64]string
}

type form struct {
	Name     string
	Select   *selectList
	Textarea string
}

type formRequest struct {
	Name     string
	Select   *datastore.Key
	Textarea string
}
