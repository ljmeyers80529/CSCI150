package csci150

import (
	"html/template"

	"github.com/ljmeyers80529/tmdb-go-gae"
)

var tpl *template.Template              // html web page processing / parsing object
var userInformation userInformationType // logged in user's information and preferences.
var movieAPI *tmdbgae.TMDb              // movie / tv database access object instance.
