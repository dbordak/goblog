package models

import (
	"time"
)

type Category struct {
	Name string `datastore:"n"`
}

type Entry struct {
	Title   string    `datastore:"t,noindex"`
	Content string    `datastore:"c,noindex"`
	Date    time.Time `datastore:"d"`
	Url     string    `datastore:"-"`
}

func NewEntry(title string, content string) *Entry {
	return &Entry{
		Title:   title,
		Content: content,
		Date:    time.Now(),
	}
}
