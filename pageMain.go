package csci150

import (
	"fmt"
	"io"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// main (top) web page.
func pageMain(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fn := req.FormValue("cmdbutton")
		switch fn {
		case "UserLogin", "Cancel": // display user login dialog box.
			http.Redirect(res, req, fmt.Sprintf("%s#openLogin", req.URL.Path), http.StatusFound)
		case "Register": // display user registration dialog box.
			http.Redirect(res, req, fmt.Sprintf("%s#openRegistration", req.URL.Path), http.StatusFound)
		case "OK": // new user registration.
			if WriteNewUserInformation(res, req) {
				http.Redirect(res, req, fmt.Sprintf("%s#openLogin", req.URL.Path), http.StatusFound)
			} else {
				http.Redirect(res, req, fmt.Sprintf("%s#openRegistration", req.URL.Path), http.StatusFound)
			}
		case "Login": // process user login.
			if checkUserLogin(res, req) {

				http.Redirect(res, req, "/counters", http.StatusFound) // change, /counters for testing only.
			} else {
				http.Redirect(res, req, fmt.Sprintf("%s#openLogin", req.URL.Path), http.StatusFound)
			}
		}
	}
	tpl.ExecuteTemplate(res, "main.html", userInformation)
}

// check if username exists.
func pageMainUsernameCheck(res http.ResponseWriter, req *http.Request) {
	if ex, _ := UsernameExists(req); ex {
		io.WriteString(res, "false")
		return
	}
	io.WriteString(res, "true")
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
