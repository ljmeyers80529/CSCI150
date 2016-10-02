package csci150

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func pageLogin(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		ctx := appengine.NewContext(req)
		fn := req.FormValue("cmdbutton")
		log.Infof(ctx, "Posting...%s", fn)
		switch fn {
		case "Register":
			http.Redirect(res, req, "/register", http.StatusSeeOther)
		case "Login":
			if checkUserLogin(res, req) {
				http.Redirect(res, req, "/count", http.StatusSeeOther)
			}
		}
	}
	tpl.ExecuteTemplate(res, "login.html", nil)
}

// check if user has successfully logged in.
// returns true if success.
func checkUserLogin(res http.ResponseWriter, req *http.Request) bool {
	var uuidKey string

	ctx := appengine.NewContext(req)
	user := req.FormValue("username")
	pass := EncryptPassword(req.FormValue("password"))

	log.Infof(ctx, "EPass: %v", pass)
	userInformation = userInformationType{"", "", "", "", -8, true, false, nil, nil} // defaults.
	if uuidKey = SearchUser(ctx, user); uuidKey != "" {
		ReadUserInformation(ctx, req, uuidKey)
		userInformation.LoggedIn = userInformation.Password == pass
	}
	return userInformation.LoggedIn
}
