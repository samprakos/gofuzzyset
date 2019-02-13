package gofuzzyset

import (
"context"
"sort"
"testing"
)

func TestMin(t *testing.T) {
	things := []int{5, 2, 7, 4, 7, 3, 1, -45, 7, -1}

	minThing := min(things...)

	if minThing != -45 {
		t.Fatalf("nope")
	}
}

func TestLevenshtein(t *testing.T) {
	ctx := context.Background()
	a := "hello"
	b := "hello"

	// distance should be zero I would think
	if levenshtein(ctx, a, b) != 0 {
		t.Fatalf("nah")
	}

	// "kitten" and "sitting" is 3
	a = "kitten"
	b = "sitting"

	if levenshtein(ctx, a, b) != 3 {
		t.Fatalf("uh oh....%v", levenshtein(ctx, a, b))
	}

	// "sit" and "sitting" is 4
	a = "sit"
	b = "sitting"

	if levenshtein(ctx, a, b) != 4 {
		t.Fatalf("uh oh....%v", levenshtein(ctx, a, b))
	}

	// "hello and "goodbye" is 7 (there is no overlap)
	a = "hello"
	b = "goodbye"

	if levenshtein(ctx, a, b) != 7 {
		t.Fatalf("uh oh....%v", levenshtein(ctx, a, b))
	}

	// "flaw" and "lawn" is 2 (delete the f and add an n to the end)
	a = "flaw"
	b = "lawn"

	if levenshtein(ctx, a, b) != 2 {
		t.Fatalf("uh oh....%v", levenshtein(ctx, a, b))
	}
}

func TestIterateGrams(t *testing.T) {
	ctx := context.Background()
	a := "test the gram thing"
	gramSize := 3

	grams := iterateGrams(ctx, a, gramSize)

	for i := range grams {
		t.Logf(grams[i])
	}
}

func TestGramCounter(t *testing.T) {
	ctx := context.Background()
	a := "test the gram thing test"
	gramSize := 2

	gramsByCount := gramCounter(ctx, a, gramSize)

	for k, v := range gramsByCount {
		t.Logf("%v = %v", k, v)
	}
}

func TestIterateGramsEdgeCases(t *testing.T) {
	ctx := context.Background()
	a := "testing"
	gramSize := 15

	grams := iterateGrams(ctx, a, gramSize)

	for i := range grams {
		t.Logf(grams[i])
	}
}

func TestNew(t *testing.T) {
	ctx := context.Background()
	data := []string{"Hello", "Hell", "hEllo"}
	lowerGramSize := 3
	upperGramSize := 3

	f := New(ctx, data, true, lowerGramSize, upperGramSize, 0.33)

	t.Logf("exactSet %v", f.exactSet)
	t.Logf("itemsByGramSize %v", f.itemsByGramSize)
	t.Logf("matchDict %v", f.matchDict)
}

func TestSortByScore(t *testing.T) {
	ms := []Match{
		{
			Word: "should be last",
			Score: 0.01,
		},
		{
			Word: "should be second to last",
			Score: 0.02,
		},
		{
			Word: "should be second",
			Score: 0.90,
		},
		{
			Word: "should be first",
			Score: 0.99,
		},
	}

	sort.Sort(byScore(ms))

	if ms[0].Word != "should be first" {
		t.Fatalf("nope")
	}
	if ms[1].Word != "should be second" {
		t.Fatalf("nope")
	}
	if ms[2].Word != "should be second to last" {
		t.Fatalf("nope")
	}
	if ms[3].Word != "should be last" {
		t.Fatalf("nope")
	}
}

func TestFindMatchesForGramSize(t *testing.T) {
	ctx := context.Background()
	data := []string{"Hello", "Hell", "hEllo", "hollow"}
	lowerGramSize := 3
	upperGramSize := 3

	f := New(ctx, data, true, lowerGramSize, upperGramSize, 0.33)

	results := f.findMatchesForGramSize(ctx, "holl", 3)

	t.Logf("Results = %v", results)
}

func TestFull(t *testing.T) {
	ctx := context.Background()

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

	f := New(ctx, data, true, 2, 3, 0.33)

	results := f.Get(ctx, "mossisippi")

	t.Logf("Results = %v", results)

	f.useLevenshtein = false

	results = f.Get(ctx, "mossisippi")

	t.Logf("Results = %v", results)
}