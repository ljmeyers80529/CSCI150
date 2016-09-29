package csci150

import (
    "net/http"
    "time"
)

type cdDT struct {
    Label string
    Year, Month, Day, Hours, Minutes int    
}

var counters = []cdDT {{"Day shopping left", 2016, 12, 24, 23, 59}, 
                       {"CSCI 144 Mid term", 2016, 10, 19, 14, 0},
                       {"CSCI 150 Mid Term", 2016, 10, 27, 15, 30},
                    }

func pageTest(res http.ResponseWriter, req *http.Request) {
    tod := time.Now()
    counters = append(counters, cdDT {"Today''s Class", tod.Year(), int(tod.Month()), tod.Day(), 15, 30})
    tpl.ExecuteTemplate(res, "test.html", counters)
}
