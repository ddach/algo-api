package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sort"
	"strings"
)

func response(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("Hello world"))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/AlgoIndex", AlgoIndex)
	router.HandleFunc("/Sort{nums}", Sort)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func AlgoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sort\n")
	fmt.Fprintf(w, "Jasonize\n") //todo
}

func Sort(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numbers := vars["nums"]
	parsedNumbers := make(map[string][]int)
	err := json.NewDecoder(strings.NewReader(numbers)).Decode(&parsedNumbers)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	//todo: return as json
	for k, v := range parsedNumbers {
		sort.Ints(v)
		fmt.Fprintf(w, "%s sorted: %v", k, v)
	}
}
