package csci150

import (
	// "fmt"
	"fmt"
	"net/http"
	// "sort"
	// "strconv"
	// "strings"
	// "time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func pageResults(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	readCookie(res, req) // maintain user login / out state.apikey

	search := req.FormValue("search") // get possible title to seach for.

	if req.Method == "POST" {
		moviePost(ctx, res, req)
		if webInformation.MovieTvGame.ID != 0 { // no detail, search.
			http.Redirect(res, req, fmt.Sprintf("%s#moviemodal", req.URL.Path), http.StatusFound)
		}
	}

	webInformation.Top = searchMovies(ctx, search) // search for movies.
	webInformation.Pop = searchTV(ctx, search)     // search for tv.

	tpl.ExecuteTemplate(res, "result.html", webInformation)
}

// get list of top rated movies.
func searchMovies(ctx context.Context, search string) topPopRated {
	var tops topPopRated
	var rated topRatedPop

	tr, _ := movieAPI.SearchMovie(ctx, search, nil)

	for _, val := range tr.Results {
		rated.Title = val.Title
		rated.ID = val.ID
		rated.Rating = val.VoteAverage
		tops = append(tops, rated)
	}
	return tops
}

// get list of top rated movies.
func searchTV(ctx context.Context, search string) topPopRated {
	var tops topPopRated
	var rated topRatedPop

	tr, _ := movieAPI.SearchTv(ctx, search, nil)

	for _, val := range tr.Results {
		rated.Title = val.Name
		rated.ID = val.ID
		rated.Rating = val.VoteAverage
		tops = append(tops, rated)
	}
	return tops
}
