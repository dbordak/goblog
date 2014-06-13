package main

import (
	"github.com/astaxie/beegae"

	"controllers"
)

func init() {
	beegae.Router("/", &controllers.ListController{})
	beegae.Router("/about", &controllers.BlogController{}, "get:About")
	beegae.Router("/:year:int/:month:int/:title/:entid:int",
		&controllers.BlogController{}, "get:EntryPage")
	beegae.Router("/cat/:catid:int", &controllers.ListController{})

	adminNamespace := beegae.NewNamespace("/admin").
		Router("/", &controllers.BlogController{}, "get:AdminNav")
	adminNamespace.Namespace(beegae.NewNamespace("/add").
		Router("/ent", &controllers.FormController{}, "get:AddEnt").
		Router("/cat", &controllers.FormController{}))
	adminNamespace.Namespace(beegae.NewNamespace("/del").
		Router("/ent", &controllers.FormController{}).
		Router("/cat", &controllers.FormController{}))
	beegae.AddNamespace(adminNamespace)

	beegae.Run()
}
