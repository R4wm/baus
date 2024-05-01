package baus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var (
	RED = "\033[0;31m"
	NC  = "\033[0m" // No Color"
)

// Bible
type Bible struct {
	Testament    string `json:"canonical_testament"` // canonical testament
	Book         string `json:"book"`                // book name
	Chapter      int    `json:"chapter"`             // chapter number
	Verse        int    `json:"verse"`               // verse number
	Text         string `json:"text"`                // actual verse
	OrdinalVerse int    `json:"ordinal_verse"`       // sequential verse number
	OrdinalBook  int    `json:"ordinal_book"`        // sequential book number
}

func (b *Bible) HighlightText(searchStr string) {
	// find the word in the text and
	//  put the color before it and the NC after it
	fmt.Println(">>> YOURE IN HIGHLIGHT")
	i := strings.Index(b.Text, searchStr)

	fmt.Printf(">>> index of %s is %d\n", searchStr, i)
	// b.Text = b.Text[:i] + RED + b.Text[i:] + NC
	var buffer bytes.Buffer
	buffer.WriteString(b.Text[:i] + RED)
	fmt.Println(">>> check where  am: ", b.Text[buffer.Len()+2:])
	// buffer.WriteString(NC)
	fmt.Println("this is bufferString after NC: ", buffer.String())
	buffer.WriteString(b.Text[buffer.Len():])
	// buffer.WriteString(b.Text[i+len(RED):] + NC)
	// buffer.WriteString(b.Text[buffer.Len():])
	// b.Text = b.Text[:i] + RED + b.Text[i:] + NC
	// fmt.Println("check where i am: ", b.Text)
	b.Text = buffer.String()
	// b.Text = b.Text[:i+2] + NC + b.Text[i+2:] + NC
	// afterWordIndex := i + len(searchStr)
	// b.Text = b.Text[:afterWordIndex] + NC + b.Text[i:]
	fmt.Println(">>> this is new string: ", b.Text)

}

// App todo
type App struct {
	Router *mux.Router
	Logger *log.Logger
	Bible  []Bible
}

// RemoveTrailingSlash guess what..
func RemoveTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}

// ConsumeJSONBibleFile slurp json file into memory
func (app *App) ConsumeJSONBibleFile(filePath string) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("failed to consume json file: %s\n", filePath)
		app.Logger.Fatal("failed to consume json file: %s\n", filePath)
	}

	// fix this later
	if err = json.Unmarshal(dat, &app.Bible); err != nil {
		fmt.Printf("failed to marshal bible file: %s", err)
		app.Logger.Fatalf("failed to marshal bible file: %s", err)
	}
}

// HandleFailure stuff
func (app *App) HandleFailure(w http.ResponseWriter, fm string) {
	w.Header().Set("content-type", "application/text")
	w.WriteHeader(400)
	w.Write([]byte(fm))
}

// VerseRange stuff
func VerseRange(r *http.Request) []int {
	u := r.URL.EscapedPath()
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("u: ", u)
	broken := strings.Split(u, "/")
	base := broken[len(broken)-1]
	log.Println(base)
	re := regexp.MustCompile(`\d{1,3}-\d{1,3}`)
	if re.MatchString(base) {
		rangeList := strings.Split(base, "-")
		s, err := strconv.Atoi(rangeList[0])
		if err != nil {
			fmt.Println("Failed to convert start range")
			return []int{}
		}
		e, err := strconv.Atoi(rangeList[1])
		if err != nil {
			fmt.Println("failed to convert end range")
			return []int{}
		}
		return []int{s, e}
	}
	return []int{}
}

// GetVerse get verse(s)
func (app *App) GetVerse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := strings.ToUpper(params["book"])
	if book == "SONGOFSOLOMON" {
		book = "SONG OF SOLOMON"
	}
	verseRange := VerseRange(r)
	log.Println("verserange: ", verseRange)

	ordinalBookNum := Books[book]
	fmt.Println("ordinalBookNum: ", ordinalBookNum)

	chapter, err := strconv.Atoi(params["chapter"])
	if err != nil {
		w.Header().Set("content-type", "application/text")
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("problems with verse: %s\n", err)))
		return
	}
	// no verse range
	response := []Bible{}
	if len(verseRange) == 0 {
		verse, err := strconv.Atoi(params["verse"])
		if err != nil {
			log.Fatalf("problems with verse: %s\n", err)
		}
		for _, v := range app.Bible {
			if v.Book == book {
				if v.Chapter == chapter {
					if v.Verse == verse {
						response = append(response, v)
						break
					}
				}
			}
		}
	} else {
		for _, v := range app.Bible {
			if v.Book == book {
				if v.Chapter == chapter {
					if v.Verse >= verseRange[0] && v.Verse <= verseRange[1] {
						response = append(response, v)
					}
					if v.Verse == verseRange[1] {
						break
					}
				}
			}
		}
	}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("failed to marshal response: %s\n", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.Write(b)
}

// GetChapter get chapter
func (app *App) GetChapter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := strings.ToUpper(params["book"])
	if book == "SONGOFSOLOMON" {
		book = "SONG OF SOLOMON"
	}
	// TODO skip to ordinal verse
	ordinalBookNum := Books[book]
	log.Println("ordinalBookNum: ", ordinalBookNum)
	chapter, err := strconv.Atoi(params["chapter"])
	if err != nil {
		log.Printf("problems with chapter: %s\n", err)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(fmt.Sprintf("Chapters %s invalid: %s", chapter, err)))
		return
	}
	response := []Bible{}
	for _, v := range app.Bible {
		if v.Book == book {
			if v.Chapter == chapter {
				response = append(response, v)
			}
		}
	}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("failed to marshal response: %s\n", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.Write(b)
}

// SearchResult your results
type SearchResult struct {
	Count int            `json:"count"`
	Query string         `json:"query"`
	Stats map[string]int `json:"bucket_stats"`
}

// Search get chapter
func (app *App) Search(w http.ResponseWriter, r *http.Request) {
	searchString := r.URL.Query().Get("q")
	if searchString == "" {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("empty search parameter: %s\n", searchString)))
		return
	}
	fmt.Printf("searchString: %#v\n", searchString)
	response := []Bible{}
	searchIt := SearchResult{Query: searchString}

	overallCount := make(map[string]int)
	re := regexp.MustCompile(searchString)
	for _, v := range app.Bible {
		// remove italics indicator for search
		noItal := strings.Replace(v.Text, "[", "", -1)
		noItal = strings.Replace(noItal, "]", "", -1)
		if re.Match([]byte(noItal)) {
			// add highlight here
			fmt.Println("this is v: ", v.Text)
			v.HighlightText(searchString)
			fmt.Println("this is v after highlight: ", v.Text)
			response = append(response, v)
			searchIt.Count++
			overallCount[v.Book]++
		}
	}
	b, err := json.Marshal(response)
	// searchIt.Stats = overallCount
	// c, err := json.Marshal(searchIt)
	// for _, v := range c {
	// 	b = append(b, v)
	// }
	if err != nil {
		fmt.Printf("failed to marshal response: %s\n", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.Write(b)

}

// ListBooks lists books
func (app *App) ListBooks(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(orderedBooks)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.Write(b)
}

// SetupRouter where the fun begins
func (app *App) SetupRouter() {
	subrouter := app.Router.PathPrefix("/v1").Subrouter()
	subrouter.HandleFunc("/{book}/{chapter}", app.GetChapter).Methods("GET")
	subrouter.HandleFunc("/{book}/{chapter}/{verse}", app.GetVerse).Methods("GET") // can handle range ex: 1-5
	subrouter.HandleFunc("/search", app.Search).Methods("GET")
	subrouter.HandleFunc("/books", app.ListBooks).Methods("GET")

}
