package csci150

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func pageMovies(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	if req.Method == "POST" {
		moviePost(ctx, req)
		if webInformation.MovieTvGame.ID != 0 {		// no detail, search.
			http.Redirect(res, req, fmt.Sprintf("%s#moviemodal", req.URL.Path), http.StatusFound)
		}
	}
	webInformation.Counters = upcomingReleases(ctx) // upcomming moview releases.
	webInformation.Top = topRatedMovies(ctx)        // overall most popular movies.
	webInformation.Pop = popularMovies(ctx)         // current most popular movies.
	sort.Sort(sort.Reverse(webInformation.Pop))
	sort.Sort(webInformation.Counters)
	tpl.ExecuteTemplate(res, "movies.html", webInformation)
}

// get list of top rated movies.
func topRatedMovies(ctx context.Context) topPopRated {
	var tops topPopRated
	var rated topRatedPop

	tr, _ := movieAPI.GetMovieTopRated(ctx, nil)
	for _, val := range tr.Results {
		rated.Title = val.Title
		rated.ID = val.ID
		rated.Rating = val.VoteAverage
		tops = append(tops, rated)
	}
	return tops
}

// get list of top rated movies.
func popularMovies(ctx context.Context) topPopRated {
	var tops topPopRated
	var rated topRatedPop

	pop, _ := movieAPI.GetMoviePopular(ctx, nil)
	for _, val := range pop.Results {
		rated.Title = val.Title
		rated.ID = val.ID
		rated.Rating = val.VoteAverage
		tops = append(tops, rated)
	}
	return tops
}

// get list of movies that are upcoming for release. Retrieve first 8.
func upcomingReleases(ctx context.Context) cdUpcomming {
	var cnts cdUpcomming
	var movie upComming
	pop, _ := movieAPI.GetMovieUpcoming(ctx, nil)

	for _, val := range pop.Results {

		rd, b := movieRelease(ctx, val.ID)
		s := strings.Split(rd, "-")
		if b {
			movie.Title = val.OriginalTitle
			movie.ID = val.ID
			movie.Year, _ = strconv.Atoi(s[0])
			movie.Month, _ = strconv.Atoi(s[1])
			movie.Day, _ = strconv.Atoi(s[2])
			movie.Hours = 0
			movie.Minutes = 0

			cnts = append(cnts, movie)
		}
	}
	return cnts
}

// get US release date for specified movie.
func movieRelease(ctx context.Context, id int) (string, bool) {
	var dRelease = time.Now().String()
	var found = false

	pop1, _ := movieAPI.GetMovieReleases(ctx, id, nil)
	pop2 := pop1.Countries
	for _, rd := range pop2 {
		if rd.Iso3166_1 == "US" {
			found = dRelease < rd.ReleaseDate
			dRelease = rd.ReleaseDate
			break
		}
	}
	return dRelease, found
}

/**************************************************  sort by rating  **************************************************/

func (t topPopRated) Len() int           { return len(t) }
func (t topPopRated) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t topPopRated) Less(i, j int) bool { return t[i].Rating < t[j].Rating }

/**************************************************  sort by release data  **************************************************/

func (t cdUpcomming) Len() int      { return len(t) }
func (t cdUpcomming) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t cdUpcomming) Less(i, j int) bool {
	return time.Date(t[i].Year, time.Month(t[i].Month), t[i].Day, 0, 0, 0, 0, time.UTC).String() < time.Date(t[j].Year, time.Month(t[j].Month), t[j].Day, 0, 0, 0, 0, time.UTC).String()
}
