package csci150

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

// main (top) web page.
func pageMain(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	if req.Method == "POST" {
		infoID := toInt(req, "cmdID")
		if infoID > 0 {
			switch itemType(infoID) {
			case 0:
			case 1:
				tvPost(ctx, req)
				if webInformation.MovieTvGame.ID != 0 { // no detail, search.
					http.Redirect(res, req, fmt.Sprintf("%s#tvmodal", req.URL.Path), http.StatusFound)
				}
			case 2:
			}
		}
		// 	fn := req.FormValue("cmdbutton")
		// 	switch fn {
		// 	case "UserLogin": // display user login dialog box.
		// 		http.Redirect(res, req, "/login", http.StatusFound)
		// 		// case "Register": // display user registration dialog box.
		// 		// 	http.Redirect(res, req, fmt.Sprintf("%s#openRegistration", req.URL.Path), http.StatusFound)
		// 		// case "OK": // new user registration.
		// 		// 	if WriteNewUserInformation(res, req) {
		// 		// 		http.Redirect(res, req, fmt.Sprintf("%s#openLogin", req.URL.Path), http.StatusFound)
		// 		// 	} else {
		// 		// 		http.Redirect(res, req, fmt.Sprintf("%s#openRegistration", req.URL.Path), http.StatusFound)
		// 		// 	}
		// 		// case "Login": // process user login.
		// 		// 	if checkUserLogin(res, req) {

		// 		// 		http.Redirect(res, req, "/counters", http.StatusFound) // change, /counters for testing only.
		// 		// 	} else {
		// 		// 		http.Redirect(res, req, fmt.Sprintf("%s#openLogin", req.URL.Path), http.StatusFound)
		// 		// 	}
		// 	}
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
		case 1:
			tvi, _ := movieAPI.GetTvInfo(ctx, wi.ID, nil)
			wi.TV = true
			wi.Title = tvi.Name
			wi.Rating = tvi.VoteAverage
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
