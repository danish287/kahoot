package questions

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/danish287/kahoot/config"
)


// Kahoot will
func Kahoot(topic string) {
	if topic == "gobatch" {
		GenerateKahoot("questions/gobatch.txt")
	} else if topic == "go" {
		GenerateKahoot("questions/go.txt")
	} else if topic == "docker" {
		GenerateKahoot("questions/docker.txt")
	} else if topic == "go" {
		GenerateKahoot("questions/kubernates.txt")
	} else if topic == "go" {
		GenerateKahoot("questions/networking.txt")
	} else {
		fmt.Println("You did not declare a valid topic. Please try again using a valid topic.")
		fmt.Println("i.e -topic=kubernates, -topic=go, -topic=docker, -topic=gobatch, -topic=netowrking")
	}
}

func GenerateKahoot(fname string) {

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		temp := scanner.Text()
		scanner.Scan()
		config.Questions[temp] = scanner.Text()

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
