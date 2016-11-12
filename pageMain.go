package csci150

import (
	"fmt"
	"net/http"
	"strings"

	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// main (top) web page.
func pageMain(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	readCookie(res, req) // maintain user login / out state.apikey

	if req.Method == "POST" {
		infoID := toInt(req, "cmdID")
		removeID := toInt(req, "cmdRM")
		if infoID > 0 {
			switch itemType(infoID) {
			case 0:
				moviePost(ctx, res, req)
				if webInformation.MovieTvGame.ID != 0 { // no detail, search.
					http.Redirect(res, req, fmt.Sprintf("%s#moviemodal", req.URL.Path), http.StatusFound)
				}
			case 1:
				tvPost(ctx, res, req)
				if webInformation.MovieTvGame.ID != 0 { // no detail, search.
					http.Redirect(res, req, fmt.Sprintf("%s#tvmodal", req.URL.Path), http.StatusFound)
				}
			case 2:
			}
		}
		if removeID > 0 {
			log.Infof(ctx, "ID: %8d\tLocation: %d\n", removeID, removeItem(removeID))
		}
	}
	popWatch(ctx)
	tpl.ExecuteTemplate(res, "index.html", webInformation)
}

// get watch items.
func popWatch(ctx context.Context) {
	var wi watchedType
	var wis []watchedType

	for _, wat := range userInformation.Watched {
		wi.ID = int(wat.ID)
		wi.Movie = false
		wi.TV = false
		wi.Game = false
		switch wat.MTGType {
		case 0:
			mvi, _ := movieAPI.GetMovieInfo(ctx, wi.ID, nil)
			wi.Movie = true
			wi.Title = mvi.Title
			wi.Rating = mvi.VoteAverage
			wi.Release = mvi.ReleaseDate
			dr, ok := movieRelease(ctx, wi.ID)
			wi.Future = ok
			if wi.Future {
				s := strings.Split(dr, "-")
				wi.Year, _ = strconv.Atoi(s[0])
				wi.Month, _ = strconv.Atoi(s[1])
				wi.Day, _ = strconv.Atoi(s[2])
				wi.Hours = 0
				wi.Minutes = 0
			}
		case 1:
			tvi, _ := movieAPI.GetTvInfo(ctx, wi.ID, nil)
			wi.TV = true
			wi.Future = false
			wi.Title = tvi.Name
			wi.Rating = tvi.VoteAverage
			wi.Release = ""
		case 2:
		}
		wis = append(wis, wi)
	}
	webInformation.Watched = wis
}

func itemType(ID int) int {
	var t = -1

	for _, val := range webInformation.Watched {
		if val.ID == ID {
			if val.Movie {
				t = 0
			} else if val.TV {
				t = 1
			} else {
				t = 2
			}
		}
	}
	return t
}

// remove an item from the watch list.
func removeItem(ID int) int {
	var location = 0

	for _, val := range webInformation.Watched {
		if val.ID == ID {
			break
		}
		location++
	}
	return location
}
