package models

import (
	"time"

	"appengine/datastore"
)

type Category struct {
	Name   string         `datastore:"n"`
	Parent *datastore.Key `datastore:"p"`
}

type Entry struct {
	Title    string         `datastore:"t,noindex"`
	Category *datastore.Key `datastore:"cat"`
	Content  string         `datastore:"c,noindex"`
	Date     time.Time      `datastore:"d"`
	Url      string         `datastore:"-"`
	Name     string         `datastore:"-"`
}
