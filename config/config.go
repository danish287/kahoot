package config

import (
	"flag"
)

var (
	Topic     string
	Questions = make(map[string]string)
)

func init() {

	flag.StringVar(&Topic, "topic", "", "determines the type of questions for the kahoot game")
	flag.Parse()

}
