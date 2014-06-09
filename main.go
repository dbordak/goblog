package main

import (
	"github.com/astaxie/beegae"

	"controllers"
)

func init() {
	beegae.Router("/", &controllers.ListController{})
	beegae.Router("/about", controllers.NewSimpleController("about.tpl"))
	beegae.Router("/:year:int/:month:int/:title/:entid:int",
		&controllers.SimpleController{})
	beegae.Router("/cat/:catid:int", &controllers.ListController{})
	beegae.Router("/admin/", &controllers.SimpleController{})
	beegae.Router("/admin/add/ent", &controllers.FormController{})
	beegae.Router("/admin/del/ent", &controllers.FormController{})
	beegae.Router("/admin/add/cat", &controllers.FormController{})
	beegae.Router("/admin/del/cat", &controllers.FormController{})
	beegae.Run()
}
