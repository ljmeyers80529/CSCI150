package csci150

import (
	"net/http"
	"sort"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"fmt"
)

func pageTV(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	userDefault()

	if req.Method == "POST" {
		tvPost(ctx, req)
		if userInformation.MovieTvGame.ID != 0 {		// no detail, search.
			http.Redirect(res, req, fmt.Sprintf("%s#tvmodal", req.URL.Path), http.StatusFound)
		}
	}
	userInformation.Top = topRatedTV(ctx) // overall most popular movies.
	userInformation.Pop = popularTV(ctx)  // current most popular movies.
	sort.Sort(sort.Reverse(userInformation.Pop))
	tpl.ExecuteTemplate(res, "tv.html", userInformation)
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
