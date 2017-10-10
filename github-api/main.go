package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var languages = []string{"JavaScript", "Java", "Python", "Ruby", "C", "PHP",
	"C#", "Clojure", "Haskell", "Perl", "C++", "CSS", "HTML", "Go", "Shell",
	"Objective C", "Swift", "Typescript", "Kotlin", "Rust", "Scala", "Lua",
	"R", "Elixir", "Pascal", "VimL", "CoffeeScript", "TeX", "Matlab", "Arduino",
	"Makefile", "Groovy", "Puppet", "PowerShell", "Erlang", "Visual Basic",
	"Assembly", "ActionScript", "ASP", "OCaml", "D", "Scheme", "Dart", "Julia",
	"F#", "FORTRAN", "Haxe", "Racket", "Logos", "Cobol", "Prolog"}

func CountProjects(language string, year int, month int) (int, error) {

	time.Sleep(7 * time.Second)

	const GhEndPoint = "https://api.github.com/search/repositories?q=size:>10+"

	u := GhEndPoint + "language:" + url.QueryEscape("\""+language+"\"")

	startMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endMonth := startMonth.AddDate(0, 1, -1)
	endMonth = time.Date(endMonth.Year(), endMonth.Month(), endMonth.Day(), 23, 59, 59, 0, time.UTC)

	u += "+pushed:" + fmt.Sprintf("%04d-%02d-01", year, month)
	u += ".." + fmt.Sprintf("%04d-%02d-%02d", year, month, endMonth.Day())
	log.Println(u)
	resp, err := http.Get(u)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("Wrong GH API status code: %d", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	type GHSearchResult struct {
		TotalCount int `json:"total_count"`
	}

	var result GHSearchResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return 0, err
	}

	return result.TotalCount, nil
}

func MainHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(http.StatusOK)

	for _, language := range languages {
		var md string = language + " |"
		fmt.Fprint(w, language)
		for year := 2017; year <= 2017; year++ {
			for month := 1; month <= 12; month++ {
				count, err := CountProjects(language, year, month)
				if err != nil {
					http.Error(w, "Error on language "+language+" "+err.Error(), http.StatusNotAcceptable)
					return
				}
				log.Printf("%04d-%02d\t%s\t%d", year, month, language, count)
				md += " " + fmt.Sprintf("%6s", strconv.Itoa(count)) + " |"
				fmt.Fprintf(w, "\t%d", count)
			}
		}
		log.Printf(md)
		fmt.Fprintln(w)
	}
}

func openWebBrowser() {
	http.Get("http://localhost:8000")
}

func main() {
	go openWebBrowser()

	r := mux.NewRouter()
	r.HandleFunc("/", MainHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
