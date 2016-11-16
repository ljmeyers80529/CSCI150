package csci150

import (
	// "fmt"
	"fmt"
	"net/http"

	"github.com/Henry-Sarabia/igdbgo"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	// "google.golang.org/appengine/log"
)

func pageResults(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	readCookie(res, req) // maintain user login / out state.apikey

	search := req.FormValue("srch") // get possible title to seach for.
	webInformation.MovieTvGame.Search = search

	if req.Method == "POST" {
		movieTvPost(ctx, res, req)
		if webInformation.MovieTvGame.ID != 0 { // no detail, search.
			it := itemType(webInformation.MovieTvGame.ID)
			if it == 0 {
				http.Redirect(res, req, fmt.Sprintf("%s?srch=%s#moviemodal", req.URL.Path, search), http.StatusFound)
			} else if it == 1 {
				http.Redirect(res, req, fmt.Sprintf("%s?srch=%s#tvmodal", req.URL.Path, search), http.StatusFound)
			} else {
				http.Redirect(res, req, fmt.Sprintf("%s?srch=%s#gamemodal", req.URL.Path, search), http.StatusFound)
			}
		}

	}

	webInformation.Top = searchMovies(ctx, search) // search for movies.
	webInformation.Pop = searchTV(ctx, search)     // search for tv.
	webInformation.Game = searchGames(ctx, search) // search for games

	tpl.ExecuteTemplate(res, "result.html", webInformation)
}

func searchGames(ctx context.Context, search string) topPopRated {
	var tops topPopRated
	var rated topRatedPop

	list, err := igdbgo.GetGames(ctx, search, 15, 0, 0, "")
	if err != nil {
		return nil
	}

	for _, val := range list {
		rated.Title = val.Name
		rated.ID = val.ID
		rated.Rating = float32(setPrecision(val.Rating, 1))
		tops = append(tops, rated)
	}
	return tops
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
