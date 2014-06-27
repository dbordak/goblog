package main

import (
	"github.com/astaxie/beegae"

	"controllers"
)

func init() {
	beegae.EnableXSRF = true

	beegae.Router("/", &controllers.ListController{})
	beegae.Router("/about", &controllers.BlogController{}, "get:About")
	beegae.Router("/:year:int/:month:int/:title/:entsafe",
		&controllers.BlogController{}, "get:EntryPage")
	beegae.Router("/:catsafe", &controllers.ListController{})

	adminNamespace := beegae.NewNamespace("/admin",
		beegae.NSRouter("/", &controllers.BlogController{}, "get:AdminNav"),
		beegae.NSNamespace("/add",
			beegae.NSRouter("/ent", &controllers.AddController{}, "get:AddEnt;post:PostEnt"),
			beegae.NSRouter("/cat", &controllers.AddController{}, "get:AddCat;post:PostCat")),
		beegae.NSNamespace("/del".
			beegae.NSRouter("/ent", &controllers.DelController{}, "get:DelEnt;post:Delete"),
			beegae.NSRouter("/cat", &controllers.DelController{}, "get:DelCat;post:Delete")))
	beegae.AddNamespace(adminNamespace)

	beegae.Run()
}
