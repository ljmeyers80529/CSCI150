package csci150

import (
	"net/http"
)

func pageMovies(res http.ResponseWriter, req *http.Request) {
   	tpl.ExecuteTemplate(res, "movies.html", userInformation)
}
