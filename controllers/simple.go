package controllers

type SimpleController struct {
	BlogController
}

func (this *SimpleController) Prepare() {
	this.BlogController.Prepare()
	this.TplNames = "simple.tpl"
}

func (this *BlogController) About() {
	this.TplNames = "about.tpl"
}

func (this *BlogController) AdminNav() {
	this.TplNames = "admin.tpl"
}

func (this *BlogController) EntryPage() {
	this.TplNames = "simple.tpl"
}
