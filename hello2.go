package hello_modules2

import "fmt"

func Hello2(name string) string {
	message := fmt.Sprintf("Hello mr. %v from function hello2", name)
	return message
}
