package main

import (
	"github.com/astaxie/beegae"

	"controllers"
)

func init() {
	beegae.Router("/", &controllers.ListController{})
	beegae.Router("/about", &controllers.BlogController{}, "get:About")
	beegae.Router("/:year:int/:month:int/:title/:entsafe",
		&controllers.BlogController{}, "get:EntryPage")
	beegae.Router("/:catsafe", &controllers.ListController{})

	adminNamespace := beegae.NewNamespace("/admin").
		Router("/", &controllers.BlogController{}, "get:AdminNav")
	adminNamespace.Namespace(beegae.NewNamespace("/add").
		Router("/ent", &controllers.AddController{}, "get:AddEnt;post:PostEnt").
		Router("/cat", &controllers.AddController{}, "get:AddCat;post:PostCat"))
	adminNamespace.Namespace(beegae.NewNamespace("/del").
		Router("/ent", &controllers.DelController{}, "get:DelEnt;post:Delete").
		Router("/cat", &controllers.DelController{}, "get:DelCat;post:Delete"))
	beegae.AddNamespace(adminNamespace)

	beegae.Run()
}
