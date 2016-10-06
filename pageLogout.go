package csci150

import (
	"net/http"
)

func pageLogout(res http.ResponseWriter, req *http.Request) {
    userInformation = userInformationType{"", "", "", "", -8, true, false, nil, nil} // defaults.
    http.Redirect(res, req, "/", http.StatusSeeOther)
}
