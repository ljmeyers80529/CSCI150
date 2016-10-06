package csci150

import (
	"net/http"
)

// main (top) web page.
func pageMain(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
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
	tpl.ExecuteTemplate(res, "index.html", userInformation)
}
