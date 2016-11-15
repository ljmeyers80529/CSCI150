package csci150

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/Henry-Sarabia/igdbgo"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func pageGames(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	readCookie(res, req)
	if req.Method == "POST" {
		gamePost(ctx, res, req)
		if webInformation.MovieTvGame.ID != 0 {
			http.Redirect(res, req, fmt.Sprintf("%s#gamemodal", req.URL.Path), http.StatusFound)
		}
	}

	webInformation.Counters = upcomingGames(ctx)
	webInformation.Top = topRatedGames(ctx)
	webInformation.Pop = popularGames(ctx)
	sort.Sort(sort.Reverse(webInformation.Pop))
	sort.Sort(webInformation.Counters)
	tpl.ExecuteTemplate(res, "games.html", webInformation)
}

//upcomingGames creates a slice of upComming types representing the 10 next upcoming games to use in the template
func upcomingGames(ctx context.Context) cdUpcomming {
	var cnts cdUpcomming
	var game upComming
	list, err := igdbgo.GetUpcoming(ctx)
	if err != nil {
		log.Infof(ctx, "\n\nError: %v\n\n", err)
		return nil
	}

	log.Infof(ctx, "\n\nGames: %v\n\n", list)

	for _, val := range list {
		game.Title = val.Name
		game.ID = val.ID
		game.Year, game.Month, game.Day = val.GetDate()
		game.Hours = 0
		game.Minutes = 0

		cnts = append(cnts, game)
	}
	return cnts
}

//topRatedGames creates a slice of topPopRated types representing the 10 top rated games to use in the template
func topRatedGames(ctx context.Context) topPopRated {
	var tops topPopRated
	var rated topRatedPop

	list, err := igdbgo.GetTop(ctx)
	if err != nil {
		log.Infof(ctx, "\n\nError: %v\n\n", err)
		return nil
	}

	log.Infof(ctx, "\n\nGames: %v\n\n", list)

	for _, val := range list {
		rated.ID = val.ID
		rated.Title = val.Name
		rated.Rating = float32(setPrecision(val.Rating, 1))
		tops = append(tops, rated)
	}
	return tops
}

//popularGames creates a slice a topPopRated types representing the 10 most popular games to use in the template
func popularGames(ctx context.Context) topPopRated {
	var tops topPopRated
	var rated topRatedPop

	// Deprecated with the SetOptions code
	/*
		opt, err := igdbgo.SetOptions("", 10, 2, 1) //any title, 10 listings, by popularity, descending
		if err != nil {
			return nil
		}
	*/

	//list, err := igdbgo.GetGames(ctx, "", 10, 2, 1, "") //any title, 10 listings, by popularity, descending
	list, err := igdbgo.GetPop(ctx)
	if err != nil {
		log.Infof(ctx, "\n\nError: %v\n\n", err)
		return nil
	}

	log.Infof(ctx, "\n\nGames: %v\n\n", list)

	for _, val := range list {
		rated.ID = val.ID
		rated.Title = val.Name
		rated.Rating = float32(setPrecision(val.Rating, 1))
		tops = append(tops, rated)
	}
	return tops
}
