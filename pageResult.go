package csci150

import (
	// "fmt"
	"net/http"
	// "sort"
	// "strconv"
	// "strings"
	// "time"

	// "golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func pageResults(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	readCookie(res, req) // maintain user login / out state.apikey

	searchCmd := req.FormValue("cmd") // get possible search type.
	search := req.FormValue("search") // get possible title to seach for.
	log.Infof(ctx, "Cmd: %s\t\tSearch: %s", searchCmd, search)

	sm, _ := movieAPI.SearchMovie(ctx, search, nil)
	log.Infof(ctx, "Lists %v\n\n", sm)

	tpl.ExecuteTemplate(res, "results.html", webInformation)
}
