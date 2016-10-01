package csci150

import (
	"net/http"
	"time"
)

func pageTest(res http.ResponseWriter, req *http.Request) {

	var counters = []cdDT{{"Day shopping left", 2016, 12, 24, 23, 59},
		{"CSCI 144 Mid term", 2016, 10, 19, 14, 0},
		{"CSCI 150 Mid Term", 2016, 10, 27, 15, 30},
	}
	tod := time.Now()
	counters = append(counters, cdDT{"Today''s Class", tod.Year(), int(tod.Month()), 4, 15, 30})

	userInformation.Counters = counters
	tpl.ExecuteTemplate(res, "test.html", userInformation)
}
