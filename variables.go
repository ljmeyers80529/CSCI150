package csci150

import (
    "html/template"
)

var tpl *template.Template                  // html web page processing / parsing object
var userInformation userInformationType     // logged in user's information and preferences.
