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
	"time"
)

// json is structured so we can just store text + author
type Quote struct {
	Text   string
	Author string
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

	// // iterate through each quote and print them
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("\"%s\" - %s\n", quoteList[i].Text, quoteList[i].Author)
	// }

	// set the seed for rng
	rand.Seed(time.Now().UnixNano())
	quoteNum := rand.Intn(len(quoteList))
	randQuote := quoteList[quoteNum]
	fmt.Printf("\"%s\" - %s\n", randQuote.Text, randQuote.Author)
}
