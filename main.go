package csci150

import (
	"html/template"
	"net/http"
)

func init() {
	configureResourceLocation("images", "img")
	configureResourceLocation("css", "css")
	configureResourceLocation("images", "js/images")
	configureResourceLocation("js", "js")
	http.Handle("/favicon.ico", http.NotFoundHandler())           // ignore favicon request (error 404)
	http.HandleFunc("/", pageMain)                                // main page.
	http.HandleFunc("/login", pageLogin)
	http.HandleFunc("/logout", pageLogout)
	http.HandleFunc("/register", pageRegister)
	http.HandleFunc("/username/check", pageRegisterUsernameCheck) // verify username is unique.
	http.HandleFunc("/movies", pageMovies)
	http.HandleFunc("/games", pageGames)
	http.HandleFunc("/tv", pageTV)
	// http.HandleFunc("/result", pageResult)

	// http.HandleFunc("/count", pageTest)
	tpl = template.Must(template.ParseGlob("html/*.html"))
}

// map resource physical location to href relative location.
// phyDir : resource files physical location relative to html file.
// hrefDir: resource location as defined withing the href tag.
func configureResourceLocation(phyDir, hrefDir string) {
	fs := http.FileServer(http.Dir(phyDir))
	fs = http.StripPrefix("/"+hrefDir, fs)
	http.Handle("/"+hrefDir+"/", fs)
}
