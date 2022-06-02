package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	baus "github.com/r4wm/baus/internal"
)

const (
	url = "http://localhost:9664/v1"
	//url = "https://mintz5.duckdns.org/bible"
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// Query the thing for client
type Query struct {
	Book        string
	Chapter     int
	VerseStart  int
	VerseEnd    int
	SearchQuery string
}

func (q *Query) Search() {
	// check book is correct
	url := url + "/search"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("cant talk to %s", url)
	}
	// format query
	qu := req.URL.Query()
	qu.Add("q", q.SearchQuery)
	req.URL.RawQuery = qu.Encode()
	// print result
	resp, err := http.Get(req.URL.String())
	if err != nil {
		log.Fatalf("Failed to get %s", req.URL.String())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response from %s", req.URL.String())
	}
	result := []baus.Bible{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Failed to json body")
	}
	for _, v := range result {
		fmt.Printf("%s %d:%d\n", v.Book, v.Chapter, v.Verse)
		fmt.Printf("%s\n\n", v.Text)
	}
}

func (q *Query) GetChapter() {
	requestURL := fmt.Sprintf("%s/%s/%d", url, q.Book, q.Chapter)
	resp, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("Failed to get chapter: ")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read chapter body")
	}
	result := []baus.Bible{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalf("Failed to json body")
	}
	if len(result) == 0 {
		return
	}
	// print result
	fmt.Printf("%s %d\n", strings.ToUpper(q.Book), q.Chapter)
	for _, v := range result {
		fmt.Printf("%d. %s\n\n", v.Verse, v.Text)
	}
}

// GetVerse get single verse or verse range
func (q *Query) GetVerse() {
	var requestURL string
	switch q.VerseEnd {
	case 0:
		requestURL = fmt.Sprintf("%s/%s/%d/%d", url, q.Book, q.Chapter, q.VerseStart)
	default:
		requestURL = fmt.Sprintf("%s/%s/%d/%d-%d", url, q.Book, q.Chapter, q.VerseStart, q.VerseEnd)
	}
	resp, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("Failed to get verse: ")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read verse body")
	}
	result := []baus.Bible{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalf("Failed to json body")
	}
	if len(result) == 0 {
		return
	}
	// print result
	fmt.Printf("%s %d\n", strings.ToUpper(q.Book), q.Chapter)
	for _, v := range result {
		fmt.Printf("%d. %s\n\n", v.Verse, v.Text)
	}
}

func main() {
	m := Query{}
	switch len(os.Args) {
	case 1:
		fmt.Println("nothing")
		return
	case 2:
		fmt.Println("just book")
	case 3:
		if strings.ToLower(os.Args[1]) == "search" {
			m.SearchQuery = os.Args[2]
			m.Search()
		} else {
			m.Book = os.Args[1]
			chapter, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal("Cant parse chapter %s", os.Args[2])
			}
			m.Chapter = chapter
			m.GetChapter()
		}
	case 4:
		fmt.Println("book chapter verse")
		m.Book = os.Args[1]
		chapterCandidate, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("failed to parse chapter")
			return
		}
		m.Chapter = chapterCandidate
		m.VerseStart, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("failed to parse start verse")
			return
		}
		m.VerseEnd = 0
		m.GetVerse()
	case 5:
		fmt.Println("book chapter verseStart verseEnd")
		m.Book = os.Args[1]
		chapterCandidate, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("failed to parse chapter")
			return
		}
		m.Chapter = chapterCandidate
		m.VerseStart, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("failed to parse start verse")
			return
		}
		m.VerseEnd, err = strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println("failed to parse end verse")
			return
		}
		m.GetVerse()
	}
}
