package main

import (
	"math/rand"
	"fmt"
	"time"
	"os"
)

func ElizaResponse (input string) string {

	possibleResponses := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?", 
		"How does that make you feel?", 
		"Why do you say that?"}

	numResponses := len(possibleResponses) // total number of responses
	index := rand.Intn(numResponses) // index between 0 -> len - 1
	return possibleResponses[index] // choose the random element.
}

func printQuestionAndResponse(input string) {
	fmt.Fprintf(os.Stdout, "Input: %s, Response: %s\n", input, ElizaResponse(input))
}

func main(){
	rand.Seed(time.Now().UnixNano()) // ensure we don't get the same values each time

	inputs := []string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I’m looking forward to the weekend.",
		"My grandfather was French!"}

	for _, input := range inputs {
		printQuestionAndResponse(input)
	}

}