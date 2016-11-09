package csci150

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	// "google.golang.org/appengine/log"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
)

/**************************************************  public functions  **************************************************/

// UsernameExists checks the user name dictionary if an
func UsernameExists(req *http.Request) (bool, error) {
	var names dictionaryUserName

	ctx := appengine.NewContext(req)    // generate an appengine context.
	bs, err := ioutil.ReadAll(req.Body) // read username as it is typed in.
	if err != nil {
		return false, err // exit if error.
	}
	names.Name = string(bs) // convert.
	return readDataStore(ctx, nameDict, names.Name, &names) != nil, nil
}

// SearchUser searches if a username has been registered.
// return true and the user's uuid if found otherwise returns false and the uuid is empty.
func SearchUser(ctx context.Context, user string) (ukey string) {
	var ui dictionaryUserName

	user = strings.ToLower(user)
	dsQuery := datastore.NewQuery(nameDict).Run(ctx)
	for {
		_, err := dsQuery.Next(&ui)
		if err == datastore.Done {
			break
		}
		if found := strings.ToLower(ui.Name) == user; found {
			ukey = ui.UUID
			break
		}
	}
	return
}

// WriteNewUserInformation writes newly register user information and preferences to datastore / memcache.
// req contains all received information.
func WriteNewUserInformation(res http.ResponseWriter, req *http.Request) (registered bool) {
	var names dictionaryUserName
	var err error

	ctx := appengine.NewContext(req)
	pass := req.FormValue("newpassword")
	conf := req.FormValue("confirm")
	fn := req.FormValue("fullname")
	un := req.FormValue("newusername")
	names.Name = un

	if pass == conf && fn != "" && un != "" {
		uid, _ := generateUUID()
		names = dictionaryUserName{
			Name: un,
			UUID: uid,
		}
		userInformation = userInformationType{
			UserID:   uid,
			Name:     fn,
			Password: EncryptPassword(pass),
			Username: un,
			Timezone: toInt(req, "timezone"),
			DST:      req.FormValue("dst") == "1",
			LoggedIn: true,
		}
		if err = writeDataStore(ctx, nameDict, un, &names); err == nil {
			err = WriteUserInformation(ctx, req)
		}
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		registered = true
		log.Infof(ctx, "User: %v", userInformation)
	}
	return
}

// WriteUserInformation write user data to both memcache and datastore.
func WriteUserInformation(ctx context.Context, req *http.Request) error {
	var err error

	err = writeDataStore(ctx, dataKind, userInformation.UserID, &userInformation)
	if err == nil {
		err = writeMemcache(ctx, req)
	}
	return err
}

// ReadUserInformation reads user information first from memcache,
// and if not present read from datastore and write that data back into memcache.
func ReadUserInformation(ctx context.Context, req *http.Request, userID string) error {
	var err error
	if !readMemcache(ctx, req, userID) {
		err := readDataStore(ctx, dataKind, userID, &userInformation)
		if err == nil {
			err = writeMemcache(ctx, req)
		}
	}
	return err
}

// EncryptPassword encrypts user password using prefix and suffix salt value and using sha256 hashing.
func EncryptPassword(pass string) string {
	h := sha256.New()
	io.WriteString(h, passwordPrefix)
	io.WriteString(h, pass)
	io.WriteString(h, passwordSuffix)
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**************************************************  private functions  **************************************************/

// read designated data from the desired datastore.
// returns if there was an error.
// data is returned within the data interface parameter.
func readDataStore(ctx context.Context, kind, key string, data interface{}) error {
	dsKey := datastore.NewKey(ctx, kind, key, 0, nil)
	err := datastore.Get(ctx, dsKey, data)
	return err
}

// write data to the datastore.
//returns if there was an error, reports an error 500.
func writeDataStore(ctx context.Context, kind, key string, data interface{}) error {
	dsKey := datastore.NewKey(ctx, kind, key, 0, nil)
	_, err := datastore.Put(ctx, dsKey, data)
	return err
}

// read information based on logged in user's uuid.
// returns true if data read successfully.
func readMemcache(ctx context.Context, req *http.Request, userID string) bool {
	item, err := memcache.Get(ctx, userID)
	if err == nil {
		err = json.Unmarshal(item.Value, &userInformation)
	}
	return err == nil
}

// write user data to memcache.
// data is to be defined within the userInformation variable.
func writeMemcache(ctx context.Context, req *http.Request) error {

	bs, err := json.Marshal(userInformation)
	if err != nil {
		return err
	}
	memData := memcache.Item{
		Key:   userInformation.UserID,
		Value: bs,
	}
	err = memcache.Set(ctx, &memData)
	if err != nil {
		return err
	}
	return nil
}

// get an UUID for user.
func generateUUID() (string, error) {
	uuid, err := uuid.NewV4()
	return uuid.String(), err
}

// ToInt converts returned form value data to integer
// req: http request containing data to be converted.
// key: field key / name of data control.
// returns converted value, if error, returns 0
func toInt(req *http.Request, key string) (val int) {
	tv, err := strconv.Atoi(req.FormValue(key))
	if err == nil {
		val = int(tv)
	}
	return
}

// set defaults.
func userDefault() {
	if userInformation.Timezone == 0 {
		mtg := movieTvGameInformation{0, "", "", "", 0, 0, nil, 0}
		userInformation = userInformationType{"", "", "", "", -8, true, false, nil, nil, nil, nil, mtg} // defaults.
	}
}

// tv information and search results.
func tvPost(ctx context.Context, req *http.Request) {
	var g []string

	info := toInt(req, "cmdID")             // get possible movie id to show detail.
	searchCmd := req.FormValue("cmdSearch") // get possible search type.
	search := req.FormValue("search")       // get possible title to seach for.

	log.Infof(ctx, "Info %7d\tCmd: %-12s\tSearch: %s", info, searchCmd, search)
	
	userInformation.MovieTvGame.ID = 0		// no detail, search.
	if info != 0 {
		burl, _ := movieAPI.GetConfiguration(ctx)
		tvi, _ := movieAPI.GetTvInfo(ctx, info, nil)
		userInformation.MovieTvGame.ID = info
		userInformation.MovieTvGame.Image = fmt.Sprintf("%s%s%s", burl.Images.BaseURL, burl.Images.PosterSizes[1], tvi.PosterPath)
		userInformation.MovieTvGame.Description = tvi.Overview
		userInformation.MovieTvGame.TVSeasons = tvi.NumberOfSeasons
		userInformation.MovieTvGame.TVEpisodes = tvi.NumberOfEpisodes
		for _, gn := range tvi.Genres {
			g = append(g, gn.Name)
		}
		userInformation.MovieTvGame.Genres = g
	}
}

// movie information and search results.
func moviePost(ctx context.Context, req *http.Request) {
	var g []string

	info := toInt(req, "cmdID")            // get possible movie id to show detail.
	searchCmd := req.FormValue("cmdSearch") // get possible search type.
	search := req.FormValue("search")       // get possible title to seach for.

	log.Infof(ctx, "Info %7d\tCmd: %-12s\tSearch: %s", info, searchCmd, search)

	userInformation.MovieTvGame.ID = 0		// no detail, search.
	if info != 0 {
		burl, _ := movieAPI.GetConfiguration(ctx)
		mvi, _ := movieAPI.GetMovieInfo(ctx, info, nil)
		log.Infof(ctx, "\n\nInfo %v\n\n", mvi)
		log.Infof(ctx, "\n\nReleased: %v\n\n", mvi.ReleaseDate)
		log.Infof(ctx, "\n\n%s%s%s\n\n", burl.Images.BaseURL, burl.Images.PosterSizes[1], mvi.PosterPath)

		userInformation.MovieTvGame.ID = info
		userInformation.MovieTvGame.Description = mvi.Overview
		if mvi.ReleaseDate != "" {
			userInformation.MovieTvGame.ReleaseDate = mvi.ReleaseDate;
		} else {
			userInformation.MovieTvGame.ReleaseDate = "Future";
		}
		userInformation.MovieTvGame.Image = fmt.Sprintf("%s%s%s", burl.Images.BaseURL, burl.Images.PosterSizes[1], mvi.PosterPath)
		for _, gn := range mvi.Genres {
			g = append(g, gn.Name)
		}
		userInformation.MovieTvGame.Genres = g
	}
}
