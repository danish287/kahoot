package main

import (
	// "fmt"

	// "fmt"

	"github.com/danish287/kahoot/config"
	"github.com/danish287/kahoot/game"
	"github.com/danish287/kahoot/questions"
)

func main() {

	questions.Kahoot(config.Topic)

	if len(config.Topic) != 0 {
		game.BeginGame()
	}
	

}
