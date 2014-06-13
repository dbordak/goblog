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

	adminNamespace := beegae.NewNamespace("/admin").
		Router("/", &controllers.SimpleController{})
	adminNamespace.Namespace(beegae.NewNamespace("/add").
		Router("/ent", &controllers.FormController{}).
		Router("/cat", &controllers.FormController{}))
	adminNamespace.Namespace(beegae.NewNamespace("/del").
		Router("/ent", &controllers.FormController{}).
		Router("/cat", &controllers.FormController{}))
	beegae.AddNamespace(adminNamespace)

	beegae.Run()
}
