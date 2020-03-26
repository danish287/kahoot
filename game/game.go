package game

import (
	"fmt"
	"time"

	"github.com/danish287/kahoot/config"
)

var score int
var total int
var name string
var answer string

func BeginGame() {
	score = 0
	total = 0
	fmt.Print("Enter your UserName:\n")
	fmt.Scanln(&name)
	fmt.Print("\n")
	fmt.Print("Welcome " + name + "! \nYou will be given a series of questions. Try and answer as quickly and accurately as you can. Good luck!\n\n")

	time.Sleep(4 * time.Second)

	for k, v := range config.Questions {
		fmt.Print(k)
		fmt.Print("\n")
		fmt.Scanln(&answer)
		// fmt.Println(v)
		if answer == v {
			fmt.Println(("RIGHT!"))
			score++
		}
		total++
	}

	fmt.Print("\n")
	fmt.Println("Your score is: ", score)
	fmt.Println("Total points possible: ", total)
	// if !strings.Contains(reqHoroscope, "Please try") {
	// 	fmt.Print("\nWould you like to save this horoscope as a text file? (y/n)\n")
	// 	fmt.Scanln(&config.StoreHoroscope)

	// 	if config.StoreHoroscope == "y" {
	// 		savefile.SaveFile(reqHoroscope)
	// 	}
	// }
}
