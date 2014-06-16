package models

import (
	"time"

	"appengine/datastore"
)

type Category struct {
	Name string `datastore:"n"`
}

type Entry struct {
	Title   string    `datastore:"t,noindex"`
	Content string    `datastore:"c,noindex"`
	Date    time.Time `datastore:"d"`
}

func NewEntry(title string, content string) *Entry {
	return &Entry{
		Title:   title,
		Content: content,
		Date:    time.Now(),
	}
}

func (this *Entry) Url() {
	return
}

func (this *Entry) DateQuery() *datastore.Query {
	return datastore.NewQuery("Entry").Order("d")
}
