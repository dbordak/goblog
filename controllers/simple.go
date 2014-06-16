package controllers

type SimpleController struct {
	BlogController
}

func (this *SimpleController) Prepare() {
	this.BlogController.Prepare()
	this.TplNames = "simple.tpl"
}
