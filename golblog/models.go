package golblog

import (
	"time"

	"appengine/datastore"
)

type Category struct {
	Name string
	Parent *datastore.Key
}

type Entry struct {
	Title string
	Content string
	Date time.Time
	Category *datastore.Key
}
