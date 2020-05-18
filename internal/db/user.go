package db

import (
	"encoding/json"
	"github.com/vestor/watermyplant/internal/constants"
	"github.com/vestor/watermyplant/internal/pogos"
)

func WriteUser(user *pogos.User) error {
	l.Println("Writing user", user.Username)
	err := database.Write(constants.USER_DB_PREFIX, user.Username, user)
	return err
}

func GetUser(username string) (pogos.User, error) {
	oneUser := pogos.User{}
	err := database.Read(constants.USER_DB_PREFIX, username, &oneUser)
	return oneUser, err
}

func GetAllUsers() ([]pogos.User, error) {
	var users []pogos.User
	rawUsers, err := database.ReadAll(constants.USER_DB_PREFIX)
	for _, rawUser := range rawUsers {
		user := pogos.User{}
		if err := json.Unmarshal([]byte(rawUser), &user); err != nil {
			l.Printf("Unable to read the json. With error %v", err)
		}
		user.Pass = ""
		users = append(users, user)
	}
	return users, err
}
