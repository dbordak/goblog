package golblog

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(writer http.ResponseWriter, req *http.Request) {
	context := appengine.NewContext(req)
	curr_user := user.Current(context)

	//Redirect to login page if not logged in.
	if curr_user == nil {
		url, err := user.LoginURL(context, req.URL.String())
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Location", url)
		writer.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprintf(writer, "Hello, %v!", curr_user)
}
