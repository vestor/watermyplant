package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/vestor/watermyplant/internal/db"
	"github.com/vestor/watermyplant/internal/log"
	"github.com/vestor/watermyplant/internal/pi"
	"github.com/vestor/watermyplant/internal/routes"
	"strconv"

	"net/http"
	"os"
)

var l = log.Get()

func main() {

	db.SetupDB()

	if os.Getenv("ENV") == "pi" {
		pipin := os.Getenv("PI_PIN")

		pin, err := strconv.Atoi(pipin)
		if err != nil {
			l.Fatalf("Invalid PIN specified %v \n", pipin)
		}

		l.Println("Setting up Pi")
		pi.SetupPi(int64(pin))
		l.Println("Starting subroutine to water")
		go pi.UpdateWater()
	} else {
		l.Println("Not setting up pi")
	}
	port := os.Getenv("PORT")

	l.Println("Setting up Handlers")
	http.Handle("/", routes.Handlers())

	l.Printf("Starting server on port %v\n", port)
	l.Fatal(http.ListenAndServe(":"+port, nil))
}