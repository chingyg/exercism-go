package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, theGreeter Greeter) string {
	return fmt.Sprintf("I can speak %v: %v", theGreeter.LanguageName(), theGreeter.Greet(name))
}

type Italian struct {
}

type Portuguese struct {
}

func (l Italian) LanguageName() string {
	return "Italian"
}

func (l Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (l Portuguese) LanguageName() string {
	return "Portuguese"
}

func (l Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
