package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Hello returns a greeting string for the named person when called.
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//Hellos returns a map of named people and their associated greetings.
func Hellos(names []string) (map[string]string, error) { // Notice that this is names not name as parameter to this function
	messages := make(map[string]string) //Create a map of messages
	for _, name := range names {
		message, err := Hello(name) // for each name in map, get freeting message
		if err != nil {             //If error is not nil, return nil and error
			return nil, err
		}
		messages[name] = message // Add greeting received to the messages map
	}
	return messages, nil //Return messages map
}

func init() {
	rand.Seed(time.Now().UnixNano()) //Intialize random number generator
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you here, %v!",
		"Wello %v...",
	}
	return formats[rand.Intn(len(formats))]
}
