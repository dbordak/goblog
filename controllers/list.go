package controllers

import (
	"fmt"

	"appengine/datastore"

	"models"
	"util"
)

//Struct for all list pages; fills in looped content and next/prev buttons
type ListController struct {
	BlogController
}

func (this *ListController) Get() {
	this.TplNames = "index.tpl"
	que := datastore.NewQuery("Entry").Order("-d").Limit(11)

	catsafe := this.Ctx.Input.Param(":catsafe")
	if catsafe != "" {
		akey, _ := datastore.DecodeKey(catsafe)
		que = datastore.NewQuery("Entry").Ancestor(akey).Order("-d").Limit(11)
	}

	cursorsafe := this.GetString("page")
	if cursorsafe != "" {
		cursor, err := datastore.DecodeCursor(string(cursorsafe))
		if err == nil {
			que = que.Start(cursor)
		}
	}

	t := que.Run(this.AppEngineCtx)
	var arr = [10]models.Entry{}
	i := 0
	for i < 10 {
		i++
		var ent models.Entry
		key, err := t.Next(&ent)
		if util.QueryErrHandler(err, this.AppEngineCtx, "Entry") {
			break
		}
		// ent.Url = this.UrlFor(
		// 	"BlogController.EntryPage", ":year", strconv.Itoa(ent.Date.Year()),
		// 	":month", strconv.Itoa(int(ent.Date.Month())), ":title", ent.Title,
		// 	":entid", something)

		ent.Url = fmt.Sprintf("/%d/%d/%s/%s", ent.Date.Year(),
			int(ent.Date.Month()), ent.Title, key.Encode())
		arr[i-1] = ent
	}
	this.Data["Entries"] = arr

	if cursor, err := t.Cursor(); err == nil {
		var ent models.Entry
		_, err := t.Next(&ent)
		if i == 10 && err == nil {
			this.Data["Footer"] = &footer{"", cursor.String()}
		}
	}
}
