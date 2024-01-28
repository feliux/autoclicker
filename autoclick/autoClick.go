package autoclick

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

var Delay int

func delay() {
	fmt.Printf("-> Delaying %d seconds\n", Delay)
	time.Sleep(time.Duration(Delay) * time.Second)
}

func GetLocation() (int, int) {
	x, y := robotgo.Location()
	fmt.Println("Current mouse location:", x, y)
	return x, y
}

func Move(x, y int) {
	delay()
	fmt.Printf("Moving mouse to %d, %d\n", x, y)
	robotgo.Move(x, y)
}

func Click() {
	delay()
	fmt.Println("Just clicking")
	robotgo.Click()
}

func MoveAndClick(x, y int) {
	delay()
	fmt.Printf("Moving mouse to %d, %d and clicking\n", x, y)
	robotgo.MoveClick(x, y)
}
