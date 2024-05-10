package airportrobot

import "fmt"

// Write your code here.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct {
}

func (p Italian) LanguageName() string {
	return "Italian"
}
func (p Italian) Greet(name string) string {
	return "Ciao " + name
}

type Portuguese struct {
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func (p Portuguese) Greet(name string) string {
	return "Olá " + name
}

// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
