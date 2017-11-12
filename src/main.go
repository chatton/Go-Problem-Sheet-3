package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

func applyReflections(original string) string {
	// reflection map adapted from https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/

	reflections := map[string]string{
		"am":     "are",
		"was":    "were",
		"i":      "you",
		"i'd":    "you would",
		"i've":   "you have",
		"i'll":   "you will",
		"my":     "your",
		"are":    "am",
		"you've": "I have",
		"you'll": "I will",
		"your":   "my",
		"yours":  "mine",
		"you":    "I",
		"me":     "you",
	}

	words := strings.Split(original, " ") // break the string into words

	for index, word := range words {
		// if the word is in the map, it needs to be changed
		if val, ok := reflections[strings.ToLower(word)]; ok {
			words[index] = val // update the current index to be the mapped word
		}
	}

	return strings.Join(words, " ") // re-construct the string with the edited words
}

func ElizaResponse(input string) string {
	pattern := regexp.MustCompile(`\b(?i)father`) // (?i) makes pattern lower case, \b to indicate word end.
	if pattern.MatchString(input) {
		return "Why don’t you tell me more about your father?"
	}
	// doesn't contain father
	iAmRe := regexp.MustCompile(`(?i)^(i am|im|i'm) (.*)[\.?!]`) // requires that the input end with some form of sentence terminator.
	if iAmRe.MatchString(input) {                                // capture groups are ["fullMatch", "I'm", "topic"]
		captured := iAmRe.FindStringSubmatch(input)[2] // 2 is what was after the *I'm* variant.
		captured = applyReflections(captured)
		return fmt.Sprintf("How do you know you are %s?", captured)
	}

	possibleResponses := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?"}

	numResponses := len(possibleResponses) // total number of responses
	index := rand.Intn(numResponses)       // index between 0 -> len - 1
	return possibleResponses[index]        // choose the random element.
}

func printQuestionAndResponse(input string) {
	fmt.Fprintf(os.Stdout, "Input: %s - Response: %s\n", input, ElizaResponse(input))
}

func main() {
	rand.Seed(time.Now().UnixNano()) // ensure we don't get the same values each time

	inputs := []string{
		"Im happy.",
		"I'm not happy with your responses.",
		"im not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?"}

	for _, input := range inputs {
		printQuestionAndResponse(input)
	}

}
