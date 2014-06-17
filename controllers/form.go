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

func (this *FormController) getForm() {
	selkey, err := datastore.DecodeKey(this.GetString("sel"))
	if err != nil && this.GetString("sel") != "" {
		this.AppEngineCtx.Errorf("getting request form: %v", err)
		return
	}
	this.FReq = &formRequest{
		Name:      this.GetString("name"),
		Textarea:  this.GetString("ta"),
		SelectKey: selkey,
	}
}
