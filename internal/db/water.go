package db

import (
	"encoding/json"
	"github.com/vestor/watermyplant/internal/constants"
	"github.com/vestor/watermyplant/internal/pogos"
	"strconv"
	"time"
)

// We log water everyday. Leaderboards are displayed per day. And overall
func LogWater(userWatering string) {
	var waterLog = pogos.WaterLog{}
	waterLog.Username = userWatering
	waterLog.TimeStamp = time.Now().Unix()
	day := strconv.Itoa(time.Now().YearDay())
	l.Printf("Day current %v \n", day)
	GetDb().Write(constants.WATER_LOG_DB_PREFIX+ day, strconv.Itoa(int(waterLog.TimeStamp)), waterLog)
	l.Printf("Added water l for user: %v\n", userWatering)
}

//Returns unsorted right now
func MostWateredToday() (map[string]int, error) {
	day := strconv.Itoa(time.Now().YearDay())
	var leaderboard = make(map[string]int)
	var records, e = GetDb().ReadAll(constants.WATER_LOG_DB_PREFIX + day)
	if e != nil {
		return nil, e
	}

	for _, f := range records {
		waterLog := pogos.WaterLog{}
		if err := json.Unmarshal([]byte(f), &waterLog); err != nil {
			l.Println("Error", err)
		}
		leaderboard[waterLog.Username] = leaderboard[waterLog.Username] + 1
	}
	return leaderboard, nil
}
