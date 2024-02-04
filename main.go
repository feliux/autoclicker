package main

import (
	//"bufio"
	//"os"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/feliux/autoclicker/autoclick"
)

var (
	duration  int
	delay     int
	random    bool
	userInput int
)

func init() {
	const (
		defaultDuration int    = 0
		defaultDelay    int    = 0
		usageDuration   string = "Duration to run the program (in seconds). If not set the program will ask for it."
		usageDelay      string = "Time between pulsation (in seconds). If not set the program will ask for it."
		usageRandom     string = "If true then choose a random value between 1 and delay argument."
	)
	flag.IntVar(&duration, "duration", defaultDuration, usageDuration)
	flag.IntVar(&duration, "du", defaultDuration, usageDuration+" (shorthand)")
	flag.IntVar(&delay, "delay", defaultDelay, usageDelay)
	flag.IntVar(&delay, "de", defaultDelay, usageDelay+" (shorthand)")
	flag.BoolVar(&random, "random", false, usageRandom+" (shorthand)")
	flag.BoolVar(&random, "r", false, usageRandom+" (shorthand)")
}

func main() {
	flag.Parse()
	if duration == 0 || delay == 0 {
		fmt.Println("Enter a integer value for duration...")
		_, err := fmt.Scanf("%d", &duration)
		if err != nil {
			panic("Please enter a integer value...")
		}
		fmt.Println("Enter a integer value for interval...")
		_, err = fmt.Scanf("%d", &delay)
		if err != nil {
			panic("Please enter a integer value...")
		}
	}
	x, y := autoclick.GetLocation()
	fmt.Println("Select an option below. Press enter to confirm. Ctrl+C to quit the program.")
	fmt.Println("[1] Force the mouse to the current location.")
	fmt.Println("[2] Just click where the mouse is located.")
	fmt.Println("[3] Move to location and click.")
	//bufio.NewReader(os.Stdin).ReadBytes('\n') // press enter whithout option
	_, err := fmt.Scanf("%d", &userInput)
	if err != nil {
		panic("Please enter an option...")
	}
	tf := time.Now().Add(time.Second * time.Duration(duration))
	for tf.After(time.Now()) {
		if random {
			autoclick.Delay = 1 + rand.Intn(delay)
		} else {
			autoclick.Delay = delay
		}
		switch userInput {
		case 1:
			autoclick.Move(x, y)
		case 2:
			autoclick.Click()
		case 3:
			autoclick.MoveAndClick(x, y)
		}
	}
}
