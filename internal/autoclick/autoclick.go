package autoclick

import (
	"log/slog"
	"time"

	"github.com/go-vgo/robotgo"
)

var Interval int

func delay() {
	slog.Info("delay action", "seconds", Interval)
	time.Sleep(time.Duration(Interval) * time.Second)
}

func GetLocation() (int, int) {
	x, y := robotgo.Location()
	slog.Info("current mouse location", "x", x, "y", y)
	return x, y
}

func Move(x, y int) {
	delay()
	slog.Info("moving pointer to a new location", "x", x, "y", y)
	robotgo.Move(x, y)
}

func Click() {
	delay()
	slog.Info("just clicking")
	robotgo.Click()
}

func MoveAndClick(x, y int) {
	delay()
	slog.Info("moving mouse to a new location and clicking", "x", x, "y", y)
	robotgo.MoveClick(x, y)
}
