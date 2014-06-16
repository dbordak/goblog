package util

import (
	"appengine"
	"appengine/datastore"
)

func QueryErrHandler(err error, ctx appengine.Context, qtype string) bool {
	if err == datastore.Done {
		return true
	}
	if err != nil {
		ctx.Errorf("fetching next %s: %v", qtype, err)
		return true
	}
	return false
}
