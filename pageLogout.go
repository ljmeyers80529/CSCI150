package csci150

import (
	"net/http"
)

func pageLogout(res http.ResponseWriter, req *http.Request) {
	userDefault()
    http.Redirect(res, req, "/", http.StatusSeeOther)
}
