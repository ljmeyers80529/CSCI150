package csci150

import (
	"net/http"
	"sort"

	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func pageTV(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	readCookie(res, req) // maintain user login / out state.apikey

	if req.Method == "POST" {
		tvPost(ctx, res, req)
		if webInformation.MovieTvGame.ID != 0 { // no detail, search.
			http.Redirect(res, req, fmt.Sprintf("%s#tvmodal", req.URL.Path), http.StatusFound)
		}
	}
	webInformation.Top = topRatedTV(ctx) // overall most popular movies.
	webInformation.Pop = popularTV(ctx)  // current most popular movies.
	sort.Sort(sort.Reverse(webInformation.Pop))
	tpl.ExecuteTemplate(res, "tv.html", webInformation)
}

// get list of top rated movies.
func topRatedTV(ctx context.Context) topPopRated {
	var tops topPopRated
	var rated topRatedPop

	tr, _ := movieAPI.GetTvTopRated(ctx, nil)
	for _, val := range tr.Results {
		rated.Title = val.Name
		rated.ID = val.ID
		rated.Rating = val.VoteAverage
		tops = append(tops, rated)
	}
	return tops
}

// get list of top rated movies.
func popularTV(ctx context.Context) topPopRated {
	var tops topPopRated
	var rated topRatedPop

	pop, _ := movieAPI.GetTvPopular(ctx, nil)
	for _, val := range pop.Results {
		rated.Title = val.Name
		rated.ID = val.ID
		rated.Rating = val.VoteAverage
		tops = append(tops, rated)
	}
	return tops
}
