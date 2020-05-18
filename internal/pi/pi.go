package pi

import (
	piBlaster "github.com/ddrager/go-pi-blaster"
	database "github.com/vestor/watermyplant/internal/db"
	"github.com/vestor/watermyplant/internal/log"
	"time"
)

var l = log.Get()

type Pi struct {
	blaster piBlaster.Blaster
	waterPin int64
}

/*
Raspberry Pi Controls
letItRain
*/
var pi = Pi{}
var wateringChan = make(chan string, 10)

func GetWaterChan() chan string {
	return wateringChan
}

func SetupPi(pin int64) {
	pi.blaster = piBlaster.Blaster{}
	pi.waterPin = pin
	pi.blaster.Start([]int64{pi.waterPin})
}

func UpdateWater() {
	l.Println("Started watering loop")
	for {
		var userWatering = <-wateringChan
		setAngle(10, &pi)
		time.Sleep(1 * time.Second)
		setAngle(105, &pi)
		time.Sleep(3 * time.Second)
		setAngle(10, &pi)
		time.Sleep(1 * time.Second)
		l.Printf("%v has watered the plant \n", userWatering)
		database.LogWater(userWatering)
	}
}

// Move servo to requested angle
func setAngle(angle uint8, pi *Pi) {
	pi.blaster.Apply(pi.waterPin, angleToServo(angle))
}

// helper function - converts angle to servo value (aka 'map' function)
func angleToServo(val uint8) float64 {
	return (float64(val)-0)*(0.25-0.05)/(180-0) + 0.05
}
