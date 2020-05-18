package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/vestor/watermyplant/internal/constants"
	"github.com/vestor/watermyplant/internal/log"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"

	database "github.com/vestor/watermyplant/internal/db"
	"github.com/vestor/watermyplant/internal/pogos"
)

var l = log.Get()

func TestAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API live and kicking"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	userFromReq := &pogos.User{}
	err := json.NewDecoder(r.Body).Decode(userFromReq)
	if err != nil {
		http.Error(w,"Invalid Request", http.StatusBadRequest)
		return
	}

	var userFromDb, err2 = database.GetUser(userFromReq.Username)
	if err2 != nil {
		http.Error(w, "User not found.", http.StatusBadRequest)
		return
	}

	errf := bcrypt.CompareHashAndPassword([]byte(userFromDb.Pass), []byte(userFromReq.Pass))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		http.Error(w,"Invalid login credentials. Please try again", http.StatusUnauthorized)
		return
	}

	var cookie, done = setToken(w, userFromDb)
	if !done {
		return
	}
	http.SetCookie(w, &cookie)
}


//CreateUser function -- create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := pogos.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "" {
		http.Error(w, "Empty username not permitted", http.StatusBadRequest)
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
	if err != nil {
		l.Println(err)
		http.Error(w, "Error trying to encrypt password", http.StatusInternalServerError)
	}

	user.Pass = string(pass)

	err = database.WriteUser(&user)

	if err != nil {
		l.Println("Error while writing user", err)
		http.Error(w, "Error while writing user", http.StatusInternalServerError)
		return
	}

	cookie, done := setToken(w, user)
	if !done {
		return
	}
	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode(user.Username)
}

func setToken(w http.ResponseWriter, user pogos.User) (http.Cookie, bool) {
	expiresAt := time.Now().Add(time.Minute * 30)
	tk := &pogos.Token{
		Username: user.Username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(constants.SECRET))
	if err != nil {
		l.Println(err)
		http.Error(w, "Error while trying to sign the string", http.StatusInternalServerError)
		return http.Cookie{}, false
	}

	l.Printf("Successfully logged in user: %v\n", user.Username)
	cookie := http.Cookie{
		Name:    constants.SESSION_COOKIE_NAME,
		Value:   tokenString,
		Expires: expiresAt}
	return cookie, true
}

//FetchUser function
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	var users, err = database.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := pogos.User{}
	var username = fmt.Sprintf("%v",r.Context().Value("user"))

	user, _ = database.GetUser(username)
	json.NewDecoder(r.Body).Decode(&user)

	err := database.WriteUser(&user)
	if err != nil {
		l.Println(err)
		http.Error(w, "Unable to update the user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not implemented. TooDaLoo", http.StatusMethodNotAllowed)
	return

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var username = fmt.Sprintf("%v",r.Context().Value("user"))
	var user, _ = database.GetUser(username)
	json.NewEncoder(w).Encode(&user)
}
