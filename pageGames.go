package csci150

import (
	"net/http"
)

func pageGames(res http.ResponseWriter, req *http.Request) {
	userDefault()
   	tpl.ExecuteTemplate(res, "games.html", userInformation)
}
