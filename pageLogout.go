package csci150

import (
	"net/http"
)

func pageLogout(res http.ResponseWriter, req *http.Request) {
	setUserDefault()
	updateCookie(res, req)
    http.Redirect(res, req, "/", http.StatusSeeOther)
}
