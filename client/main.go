package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	baus "github.com/r4wm/baus/internal"
)

var (
	BookChapterLimit = map[string]int{
		"GENESIS":         50,
		"EXODUS":          40,
		"LEVITICUS":       27,
		"NUMBERS":         36,
		"DEUTERONOMY":     34,
		"JOSHUA":          24,
		"JUDGES":          21,
		"RUTH":            4,
		"1SAMUEL":         31,
		"2SAMUEL":         24,
		"1KINGS":          22,
		"2KINGS":          25,
		"1CHRONICLES":     29,
		"2CHRONICLES":     36,
		"EZRA":            10,
		"NEHEMIAH":        13,
		"ESTHER":          10,
		"JOB":             42,
		"PSALMS":          150,
		"PROVERBS":        31,
		"ECCLESIASTES":    12,
		"SONG OF SOLOMON": 8,
		"ISAIAH":          66,
		"JEREMIAH":        52,
		"LAMENTATIONS":    5,
		"EZEKIEL":         48,
		"DANIEL":          12,
		"HOSEA":           14,
		"JOEL":            3,
		"AMOS":            9,
		"OBADIAH":         1,
		"JONAH":           4,
		"MICAH":           7,
		"NAHUM":           3,
		"HABAKKUK":        3,
		"ZEPHANIAH":       3,
		"HAGGAI":          2,
		"ZECHARIAH":       14,
		"MALACHI":         4,
		"MATTHEW":         28,
		"MARK":            16,
		"LUKE":            24,
		"JOHN":            21,
		"ACTS":            28,
		"ROMANS":          16,
		"1CORINTHIANS":    16,
		"2CORINTHIANS":    13,
		"GALATIANS":       6,
		"EPHESIANS":       6,
		"PHILIPPIANS":     4,
		"COLOSSIANS":      4,
		"1THESSALONIANS":  5,
		"2THESSALONIANS":  3,
		"1TIMOTHY":        6,
		"2TIMOTHY":        4,
		"TITUS":           3,
		"PHILEMON":        1,
		"HEBREWS":         13,
		"JAMES":           5,
		"1PETER":          5,
		"2PETER":          3,
		"1JOHN":           5,
		"2JOHN":           1,
		"3JOHN":           1,
		"JUDE":            1,
		"REVELATION":      22,
	}
)

const (
	url = "http://localhost:9664/v1"
	// url         = "https://mintz5.duckdns.org/bible"
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

// lazyBook only adds book to string if pattern is met
func lazyBook(shortName string) (book string, err error) {
	// TODO: write simple test
	shortName = strings.ToUpper(shortName)
	var possibleBooks []string

	// iterate the BookList and add if pattern meets
	for bookCandidate, _ := range BookChapterLimit {
		// fmt.Println("this is candidate: ", bookCandidate)
		// TODO CHange to regex /
		if strings.HasPrefix(bookCandidate, shortName) {
			// fmt.Println("Found : ", bookCandidate)
			possibleBooks = append(possibleBooks, bookCandidate)
		}
	}

	// Handle non existing book name
	if len(possibleBooks) == 0 {
		errMsg := fmt.Sprintf("no books with string combination of: %s\n", shortName)
		err = errors.New(errMsg)
		return "", err
	}

	fmt.Println("possibleBooks: ", possibleBooks)
	if len(possibleBooks) > 1 {
		errMsg := fmt.Sprintf("more than one possible choice: %s", possibleBooks)
		err = errors.New(errMsg)
		return "", err
	} else {
		book = possibleBooks[0]
		return book, nil
	}

	return book, err
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
			var err error
			m.Book, err = lazyBook(os.Args[1])
			if err != nil {
				fmt.Printf("failed to use lazyBook: %s", err)
			}
			// m.Book = os.Args[1]
			chapter, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal("Cant parse chapter %s", os.Args[2])
			}
			m.Chapter = chapter
			m.GetChapter()
		}
	case 4:
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
