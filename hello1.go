package hello_modules2

import (
	"fmt"
	"math/rand/v2"
)

func Hello1(name string) string {
	message := fmt.Sprintf("Hello mr. %v from function hello1", name)
	return message
}

func RandomHello() string {
	greetings := []string{"Hello %v", "Hi %v", "Get out %v", "Major league bullies %v", "Skin tone chicken bone leave me alone here %v"}
	return greetings[rand.IntN(len(greetings))]
}
