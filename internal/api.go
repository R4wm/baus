package baus

import (
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

// ConsumeJsonBibleFile slurp json file into memory
func (app *App) ConsumeJsonBibleFile(filePath string) {
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

// GetVerse get verse(s)
func (app *App) GetVerse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := strings.ToUpper(params["book"])

	ordinalBookNum := getOrdinalBookNumber(book)
	fmt.Println("ordinalBookNum: ", ordinalBookNum)

	chapter, err := strconv.Atoi(params["chapter"])
	if err != nil {
		log.Fatalf("problems with verse: %s\n", err)
	}
	verse, err := strconv.Atoi(params["verse"])
	if err != nil {
		log.Fatalf("problems with verse: %s\n", err)
	}
	response := []Bible{}
	for _, v := range app.Bible {
		if v.Book == book {
			if v.Chapter == chapter {
				if v.Verse == verse {
					response = append(response, v)
				}

			}
		}
	}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("failed to marshal response: %s\n", err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(b)
}

// GetChapter get chapter
func (app *App) GetChapter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := strings.ToUpper(params["book"])

	// TODO skip to ordinal verse
	ordinalBookNum := getOrdinalBookNumber(book)
	fmt.Println("ordinalBookNum: ", ordinalBookNum)

	chapter, err := strconv.Atoi(params["chapter"])
	if err != nil {
		log.Fatalf("problems with chapter: %s\n", err)
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
	w.Header().Set("content-type", "application/json")
	w.Write(b)
}

// SearchResult your results
type SearchResult struct {
	Count int            `json:"count"`
	Query string         `json:"query"`
	Stats map[string]int `json:"stats"`
}

// Search get chapter
func (app *App) Search(w http.ResponseWriter, r *http.Request) {
	searchString := r.URL.Query().Get("q")
	fmt.Println("searchString: ", searchString)

	response := []Bible{}
	searchResult := SearchResult{Query: searchString}
	overallCount := make(map[string]int)
	re := regexp.MustCompile(searchString)
	for _, v := range app.Bible {
		if re.Match([]byte(v.Text)) {
			response = append(response, v)
			searchResult.Count++
			overallCount[v.Book]++

		}
	}
	b, err := json.Marshal(response)
	searchResult.Stats = overallCount
	c, err := json.Marshal(searchResult)
	for _, v := range c {
		b = append(b, v)
	}
	if err != nil {
		fmt.Printf("failed to marshal response: %s\n", err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(b)
}

// Hi test
func (app *App) Hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi\n"))
}

// SetupRouter where the fun begins
func (app *App) SetupRouter() {
	subrouter := app.Router.PathPrefix("/v1").Subrouter()
	subrouter.HandleFunc("/hi", app.Hi).Methods("GET")
	subrouter.HandleFunc("/{book}/{chapter}", app.GetChapter).Methods("GET")
	subrouter.HandleFunc("/{book}/{chapter}/{verse}", app.GetVerse).Methods("GET")
	subrouter.HandleFunc("/search", app.Search).Methods("GET")

}
