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
	beegae.Router("/cat/:catsafe", &controllers.ListController{})

	adminNamespace := beegae.NewNamespace("/admin").
		Router("/", &controllers.BlogController{}, "get:AdminNav")
	adminNamespace.Namespace(beegae.NewNamespace("/add").
		Router("/ent", &controllers.AddController{}, "get:AddEnt;post:PostEnt").
		Router("/cat", &controllers.AddController{}, "get:AddCat;post:PostCat"))
	adminNamespace.Namespace(beegae.NewNamespace("/del").
		Router("/ent", &controllers.DelController{}, "get:GetDelEnt;post:PostDelEnt").
		Router("/cat", &controllers.DelController{}, "get:GetDelCat;post:PostDelCat"))
	beegae.AddNamespace(adminNamespace)

	beegae.Run()
}
