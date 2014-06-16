package controllers

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
