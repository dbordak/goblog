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
}

func NewEntry(title string, cat *datastore.Key, content string) *Entry {
	return &Entry{
		Title:    title,
		Category: cat,
		Content:  content,
		Date:     time.Now(),
	}
}

func (this *Entry) Url() {
	return
}
