package csci150

import (
	"encoding/base64"
	"encoding/json"
	// "fmt"
	"net/http"
)

/**************************************************  constants, types and variables  **************************************************/

const cookieSessionName string = "cookieReleaseIt"

// readCookie reads current state
// create a new cookie if it does not exists or expired.
func readCookie(res http.ResponseWriter, req *http.Request) {
	cookie := readCreateCookie(req)
	http.SetCookie(res, cookie)                  // set cookie into browser.
	userInformation = cookieInformationDecoding(cookie.Value) // decode and set user state into page variable.
}

// read an existing cookie or create a new one.
// returns the cookie.
func readCreateCookie(req *http.Request) (cookie *http.Cookie) {
	cookie, err := req.Cookie(cookieSessionName) // get if a cookie already exists (had not expired)
	if err == http.ErrNoCookie {
		cookie = newCookie() // need a new cookie.
	}
	return
}

// create a new cookie, set value fields to default values, JSON / base 64 processed.
func newCookie() (cookie *http.Cookie) {
	cookie = &http.Cookie{
		Name:     cookieSessionName,
		Value:    cookieInformationEncoding(),
		HttpOnly: true,
		//Secure: false,
	}
	return
}



// // CookieDataType contain fields to render user login information.
// type cookieDataType struct {
// 	LoggedIn bool
// 	Ldap     ldapUserInformationType
// 	User     userDataRow
// }

// // initialize logged on user information.
// var cookieData cookieDataType

// var lastPage string

// /**************************************************  user cookie information functions  **************************************************/

// // CheckLoggedIn checks if an user is already logged in.
// // redirect to the login page if not.
// // all http page handler function must call this function to load cookie
// // returns reference to a cookie structure and a boolean that returns true if first visit to page,o therwise false.
// func checkLoggedIn(res http.ResponseWriter, req *http.Request) (*cookieDataType, bool) {
// 	var lv bool

// 	ci := readCookie(res, req) // get current user information from cookie, ignore cookie object.
// 	if !ci.LoggedIn {
// 		http.Redirect(res, req, "/logout", http.StatusSeeOther)
// 	} else {
// 		if len(dsTeam.OptionList) == 0 {
// 			fmt.Println("TODO: Add code to set team select option list to teams only available to logged in user (cookie.go).")
// 			dsTeam.setOptionListTeams(ci.User)
// 		}
// 	}
// 	if lv = lastPage != req.URL.Path; lv {
// 		lastPage = req.URL.Path
// 	}
// 	return &ci, lv
// }

// // ClearCookie set cookie information fields to default, unlogged in state.
// func clearCookie() (ci cookieDataType) {
// 	ci = cookieDataType{
// 		LoggedIn: false,
// 		Ldap:     ldapUserInformationType{"", "", "", ""},
// 		User:     dsUser.clearRow(),
// 	}
// 	return
// }

// UpdateCookie updates cookie information.
// read existing cookie or re-sreate the cookie if expired.
func updateCookie(res http.ResponseWriter, req *http.Request) {
	cookie := readCreateCookie(req)
	cookie.Value = cookieInformationEncoding()
	http.SetCookie(res, cookie) // set cookie into browser.
}

// // readCookie reads current state
// // create a new cookie if it does not exists or expired.
// func readCookie(res http.ResponseWriter, req *http.Request) (ci cookieDataType) {
// 	cookie := readCreateCookie(req)
// 	http.SetCookie(res, cookie)                  // set cookie into browser.
// 	ci = cookieInformationDecoding(cookie.Value) // decode and set user state into page variable.
// 	cookieData = ci
// 	return
// }

// // read an existing cookie or create a new one.
// // returns the cookie.
// func readCreateCookie(req *http.Request) (cookie *http.Cookie) {
// 	cookie, err := req.Cookie(cookieSessionName) // get if a cookie already exists (had not expired)
// 	if err == http.ErrNoCookie {
// 		cookie = newCookie() // need a new cookie.
// 	}
// 	return
// }

// // create a new cookie, set value fields to default values, JSON / base 64 processed.
// func newCookie() (cookie *http.Cookie) {
// 	cookie = &http.Cookie{
// 		Name:     cookieSessionName,
// 		Value:    userStateDefaults(),
// 		HttpOnly: true,
// 		//Secure: false,
// 	}
// 	return
// }

// // Set user state information to default values.
// // returns state encodes using JSON / base 64.
// func userStateDefaults() string {
// 	return cookieInformationEncoding(clearCookie())
// }

// encode user state information marshalled using JSON and then converted into base 64.
// returns JSON / base 64 encoded string.
func cookieInformationEncoding() (encoded string) {
	j, err := json.Marshal(userInformation) // encode using JSON and base 64.
	if err == nil {
		encoded = base64.URLEncoding.EncodeToString(j)
	}
	return
}

// retrieve information from the cookie and decode from base 64 then unmarshal JSON.
// returns user information data.
func cookieInformationDecoding(userInfo string) (ci userInformationType) {
	decode, _ := base64.URLEncoding.DecodeString(userInfo)
	json.Unmarshal(decode, &ci)
	return
}
