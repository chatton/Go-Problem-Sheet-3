package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

type Response struct {
	re      *regexp.Regexp
	answers []string
}

func (resp *Response) randomAnswer() string {
	numResponses := len(resp.answers) // total number of responses
	index := rand.Intn(numResponses)  // index between 0 -> len - 1
	return resp.answers[index]        // choose the random element.
}

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
	allResponses := []*Response{}
	patterns := []string{"I like (.*)", "My name is (.*)", "I feel (.*)", `(?i)^(i am|im|i'm) (.*)[\.?!]`, `\b(?i)father`, "(.*)"}
	answers := [][]string{[]string{"Why do you like %s?", "Are you sure you like %s?"},
		[]string{"Hello %s, how are you?", "It's nice to meet you %s."},
		[]string{"Why do you feel %s?", "Are you sure you feel %s?"},
		[]string{"How do you know you are %s?"},
		[]string{"Tell me more about your father."},
		[]string{"I’m not sure what you’re trying to say. Could you explain it to me?",
			"How does that make you feel?",
			"Why do you say that?"}}

	for i := 0; i < len(patterns); i++ {
		resp := makeResponse(patterns[i], answers[i])
		allResponses = append(allResponses, resp)
	}

	for _, response := range allResponses {
		if response.re.MatchString(input) {
			ans := response.randomAnswer()   // choose any valid answer for this pattern
			if strings.Contains(ans, "%s") { // if we want to sub in something from the user input
				match := response.re.FindStringSubmatch(input) // get the match
				capture := match[len(match)-1]                 // the last element is the specific capture group we want
				capture = applyReflections(capture)            // process the words and swap them if required.
				ans = fmt.Sprintf(ans, capture)                // sub into answer
			}
			return ans // the answer to the question
		}
	}
	panic("Uh oh, the programmer made a mistake, the catch all should have caught any response!")
}

func printQuestionAndResponse(input string) {
	fmt.Fprintf(os.Stdout, "Input: %s - Response: %s\n", input, ElizaResponse(input))
}

func makeResponse(pattern string, answers []string) *Response {
	resp := new(Response)
	resp.re = regexp.MustCompile(pattern)
	resp.answers = answers
	return resp
}

func main() {
	rand.Seed(time.Now().UnixNano()) // ensure we don't get the same values each time

	inputs := []string{
		"Im happy.",
		"I'm not happy with your responses.",
		"im not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?",
		"I like waffles.",
		"My name is Bob."}

	for _, input := range inputs {
		printQuestionAndResponse(input)
	}

}
