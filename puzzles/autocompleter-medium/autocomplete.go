package main

import "fmt"

/**
Implement an autocomplete system.
That is, given a query string s and a set of all possible query strings,
return all strings in the set that have s as a prefix.

For example, given the query string de and the set of strings
[dog, deer, deal], return [deer, deal].

Hint: Try preprocessing the dictionary into a more
efficient data structure to speed up queries.
*/
func main() {
	dict := []string{"dog", "deer", "deal", "danish", "elephant", "edelweiss"}
	fmt.Printf("Dictionary is %v \n", dict)
	p := newPrefixer()
	for _, word := range dict {
		p.insert(word)
	}
	prefix1 := "de"
	results1 := p.searchForPrefix(prefix1)
	fmt.Printf("Searching for prefix %s has found %v \n", prefix1, results1)
	prefix2 := "f"
	results2 := p.searchForPrefix(prefix2)
	fmt.Printf("Searching for prefix %s has found %v \n", prefix2, results2)
}
