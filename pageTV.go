package csci150

import (
	"net/http"
	"sort"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func pageTV(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	userDefault()

	if req.Method == "POST" {
		info := toInt(req, "cmdID")									// get possible movie id to show detail.
		searchCmd := req.FormValue("cmdSearch")						// get possible search type. 
		search := req.FormValue("search")							// get possible title to seach for.

		log.Infof(ctx, "Info %7d\tCmd: %-12s\tSearch: %s", info, searchCmd, search)
		if info != 0 {
			burl, _ := movieAPI.GetConfiguration(ctx)
			tvi, _ := movieAPI.GetTvInfo(ctx, info, nil)
			userInformation.MovieTvGame.ID = info
			userInformation.MovieTvGame.Image = fmt.Sprintf("%s%s%s", burl.Images.BaseURL, burl.Images.PosterSizes[1], tvi.PosterPath)
			userInformation.MovieTvGame.Description = tvi.Overview;
		}
	}
	userInformation.Top = topRatedTV(ctx) 							// overall most popular movies.
	userInformation.Pop = popularTV(ctx) 							 // current most popular movies.
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
