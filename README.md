# gofuzzyset
A go implementation of the [Javascript fuzzyset library](https://github.com/Glench/fuzzyset.js), which is itself a port from a [python library](https://github.com/axiak/fuzzyset).

There is EXCELLENT documentation by the [author](https://github.com/Glench) of the Javascript library [here](http://glench.github.io/fuzzyset.js/ui/).

Usage
-----

The usage is simple. Just initialize a new fuzzyset, and ask for matches
by using ``.Get``:
```go
package main

import (
	"github.com/samprakos/gofuzzyset"
	"log"
)

func main() {
	lowerGramSize := 2
	upperGramSize := 3
	minScore := 0.33

	data := []string{
		"Alabama",
		"Alaska",
		"Arizona",
		"Arkansas",
		"California",
		"Colorado",
		"Connecticut",
		"Delaware",
		"Florida",
		"Georgia",
		"Hawaii",
		"Idaho",
		"Illinois",
		"Indiana",
		"Iowa",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Maine",
		"Maryland",
		"Massachusetts",
		"Michigan",
		"Minnesota",
		"Mississippi",
		"Missouri",
		"Montana",
		"Nebraska",
		"Nevada",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York",
		"North Carolina",
		"North Dakota",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"Pennsylvania",
		"Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Vermont",
		"Virginia",
		"Washington",
		"West Virginia",
		"Wisconsin",
		"Wyoming",
	}

	f := gofuzzyset.New(data, false, lowerGramSize, upperGramSize, minScore)

	results := f.Get("mossisippi")

	log.Printf("%v results found", len(results))
}
```
The result will ``[]Match``.
The score is between 0 and 1, with 1 being a perfect match.

```go
type Match struct {
	Word string
	Score float64
}
```
