package csci150

import (
	"net/http"
)

func pageTV(res http.ResponseWriter, req *http.Request) {
   	tpl.ExecuteTemplate(res, "tv.html", userInformation)
}
