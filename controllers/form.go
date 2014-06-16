package controllers

import (
	"appengine/datastore"
)

type FormController struct {
	BlogController
	FReq *formRequest
}

func (this *FormController) Prepare() {
	this.BlogController.Prepare()
	this.TplNames = "form.tpl"
}

type selectList struct {
	Name   string
	DefOpt bool
	Items  map[string]string
}

type form struct {
	Name     string
	Select   *selectList
	Textarea string
}

type formRequest struct {
	Name      string
	SelectKey *datastore.Key
	Textarea  string
}
