package main

import (
	"math/rand"
	"fmt"
	"time"
	"os"
	"regexp"
)

func ElizaResponse (input string) string {

	pattern := regexp.MustCompile(`\b(?i)father`) // (?i) makes pattern lower case, \b to indicate word end.
	if pattern.MatchString(input) {
		return "Why don’t you tell me more about your father?"
	} else { // doesn't contain father
		iAmRe := regexp.MustCompile("^I am (.*)")
		if iAmRe.MatchString(input) {
			captured := iAmRe.FindStringSubmatch(input)[1]
			return fmt.Sprintf("How do you know you are %s?", captured)
		}
	}

	possibleResponses := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?", 
		"How does that make you feel?", 
		"Why do you say that?"}
	
	numResponses := len(possibleResponses) // total number of responses
	index := rand.Intn(numResponses) // index between 0 -> len - 1
	return possibleResponses[index] // choose the random element.
}

func printQuestionAndResponse(input string) {
	fmt.Fprintf(os.Stdout, "Input: %s - Response: %s\n", input, ElizaResponse(input))
}

func main(){
	rand.Seed(time.Now().UnixNano()) // ensure we don't get the same values each time
	
	inputs := []string{
		"I am happy",
		"I am not happy with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?"}

	for _, input := range inputs {
		printQuestionAndResponse(input)
	}

}