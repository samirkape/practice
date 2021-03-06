// Trying JSON encoding with Go

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct { // Any field with non-capital first letter will not be marshalled
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` // No space allowed in `` qoutes
	Actors []string
}

var movies = []Movie{
	{Title: "Casanova", Year: 2005, Color: true, Actors: []string{"Heath Ledger"}},
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
}

func main() {
	Marshal()
	MarshalIndent()
}

func Marshal() {
	data, err := json.Marshal(movies) // No white spaces
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func MarshalIndent() {
	data, err := json.MarshalIndent(movies, "", " ") // Human readable indentation
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
