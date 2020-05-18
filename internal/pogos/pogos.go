package pogos

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

type WaterLog struct {
	TimeStamp int64  `json:"time"`
	Username  string `json:"username"`
}

type LeaderboardEntry struct {
	Username string `json:"username"`
	TimesWatered int `json:"times-watered"`
}

type Exception struct {
	Message string `json:"message"`
}

//Token struct declaration
type Token struct {
	Username string `json:"username"`
	*jwt.StandardClaims
}