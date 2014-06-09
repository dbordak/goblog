package models

import (
	"time"

	"appengine/datastore"
)

type Category struct {
	Name   string         `datastore:"n"`
	Parent *datastore.Key `datastore:"p" form:",select"`
}

type Entry struct {
	Title    string         `datastore:"t,noindex"`
	Category *datastore.Key `datastore:"cat" form:",select"`
	Content  string         `datastore:"c,noindex" form:",textarea"`
	Date     time.Time      `datastore:"d" form:"-"`
	Url      string         `datastore:"-" form:"-"`
}
