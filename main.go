// Description: fortune(ish) but using Golang
// Resources:
//	https://flaviocopes.com/go-tutorial-fortune/
//	https://github.com/dmgk/faker/blob/master/hacker.go
//	https://forum.freecodecamp.org/t/free-api-inspirational-quotes-json-with-code-examples/311373

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// json is structured so we can just store text + author
type Quote struct {
	Text   string
	Author string
}

func formatQuote(myQuote Quote) string {
	words := strings.Split(myQuote.Text, " ")
	count := 0
	result := ""
	for i := 0; i < len(words); i++ {
		if count+len(words[i]) < 60 {
			result += words[i] + " "
			count += len(words[i]) + 1
		} else if count+len(words[i]) == 60 {
			count = 0
			result += words[i] + "\n"
		} else {
			count = 0
			result += "\n" + words[i] + " "
		}
	}

	// add formatted author string
	myAuthor := myQuote.Author
	if len(myQuote.Author) == 0 {
		myAuthor = "Unknown"
	}
	result += "\n" + strings.Repeat(" ", 60-len(myAuthor)-2) + "- " + myAuthor

	return result
}

func main() {
	// open the json file (which is pulled from link in references)
	jsonFile, err := os.Open("quotes.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	// read the json file as bytes
	fileBytes, _ := ioutil.ReadAll(jsonFile)

	// unmarshal json contents into quoteList slice
	var quoteList []Quote
	json.Unmarshal(fileBytes, &quoteList)

	// set the seed for rng
	rand.Seed(time.Now().UnixNano())
	quoteNum := rand.Intn(len(quoteList))
	randQuote := quoteList[quoteNum]

	// format the quote
	fmt.Println(formatQuote(randQuote))
	// fmt.Printf("\"%s\" - %s\n", randQuote.Text, randQuote.Author)
}
