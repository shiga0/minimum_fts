package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/kljensen/snowball"
)

type Doc struct {
	Documents []document `json:"documents"`
}
type document struct {
	ID   int    `json:"ID"`
	Text string `json:"Text"`
}

type index map[string][]int

func (idx index) add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func loadDocuments(path string) ([]document, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var doc Doc
	json.Unmarshal(jsonData, &doc)

	return doc.Documents, nil
}

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i], _ = snowball.Stem(token, "english", true)
	}
	return r
}

func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}

func intersection(a []int, b []int) []int {
	var result []int
	for _, v1 := range a {
		for _, v2 := range b {
			if v1 == v2 {
				result = append(result, v1)
			}
		}
	}
	return result
}

func (idx index) search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}

func main() {
	var path string
	var search_query string

	path = "sample.json"
	search_query = "deep reason"

	log.Printf("=========================")
	log.Printf("🚀    minimum_fts     🚀")
	log.Printf("=========================")

	docs, err := loadDocuments(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("📕 Loaded %d", len(docs))

	idx := make(index)
	idx.add(docs)
	log.Printf("🛠  Indexed %d", len(docs))

	log.Printf("🤓 search_word : %s", search_query)

	matchedIDs := idx.search(search_query)
	log.Printf("🔎 Search found %d", len(matchedIDs))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("\t%d %s\n", id, doc.Text)
	}

	log.Printf("========================")
	log.Printf("🎊🥳🎉 Congrats!  🎊🥳🎉")
	log.Printf("========================")
}
