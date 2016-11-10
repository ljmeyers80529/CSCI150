package csci150

import (
	"net/http"
)

func pageLogout(res http.ResponseWriter, req *http.Request) {
	setUserDefault()
    http.Redirect(res, req, "/", http.StatusSeeOther)
}
