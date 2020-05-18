package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	controller "github.com/vestor/watermyplant/internal/controllers"
	auth "github.com/vestor/watermyplant/internal/utils"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.HandleFunc("/", controller.TestAPI).Methods("GET")
	r.HandleFunc("/api", controller.TestAPI).Methods("GET")
	r.HandleFunc("/register", controller.CreateUser).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/leaderboard", controller.GetWaterLeaderBoard).Methods("GET")

	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	//s.HandleFunc("/users", controller.FetchUsers).Methods("GET")
	s.HandleFunc("/user", controller.GetUser).Methods("GET")
	s.HandleFunc("/user", controller.UpdateUser).Methods("PUT")
	s.HandleFunc("/user", controller.DeleteUser).Methods("DELETE")
	s.HandleFunc("/water", controller.WaterDaWood).Methods("POST")
	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, " +
			"Content-Type, " +
			"Content-Length, " +
			"Accept-Encoding, " +
			"X-CSRF-Token, " +
			"Authorization, " +
			"Access-Control-Request-Headers, " +
			"Access-Control-Request-Method, " +
			"Connection, " +
			"Host, " +
			"Origin, " +
			"User-Agent, " +
			"Referer, " +
			"Cache-Control, " +
			"X-header")
		next.ServeHTTP(w, r)
	})
}
