package controllers

import (
	"encoding/json"
	"fmt"
	database "github.com/vestor/watermyplant/internal/db"
	"github.com/vestor/watermyplant/internal/pi"
	"net/http"
)

func WaterDaWood(w http.ResponseWriter, r *http.Request) {
	var username = fmt.Sprintf("%v",r.Context().Value("user"))
	select {
	case pi.GetWaterChan() <- username:
		return
	default:
		http.Error(w, "Can't water anymore", http.StatusServiceUnavailable)
		return
	}
}

func GetWaterLeaderBoard(w http.ResponseWriter, r *http.Request) {
	var leaderboard, err = database.MostWateredToday()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(&leaderboard)
	}
}
