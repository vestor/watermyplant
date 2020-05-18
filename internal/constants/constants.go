package constants

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var SECRET = os.Getenv("SECRET")
const SESSION_COOKIE_NAME = "WMPSESSION"
const WATER_LOG_DB_PREFIX = "water_log_"
const USER_DB_PREFIX = "users"
