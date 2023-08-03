package main

import (
	"fmt"
	"sort"
)

// Build a Markov chain of n characters from a list of words.
// The chain is represented as a map of strings to slices of strings.
// The key is a string of n characters, and the value is a slice of all
// the characters that follow that string in the source text.
// The chain terminates if a STOP token is reached; this is represented
// by an empty string in the slice of following characters.

var EOW = ""

type MarkovChain struct {
	m        map[string][]string
	initVals []string
}

func buildChain(words []string, n int) MarkovChain {
	chain := make(map[string][]string)
	for _, word := range words {
		if len(word) < n {
			continue
		}
		for i := 0; i < len(word)-n; i++ {
			key := word[i : i+n]
			if _, ok := chain[key]; !ok {
				chain[key] = []string{}
			}
			chain[key] = append(chain[key], string(word[i+n]))
		}
		// Add termination marker for the last ngram in each word
		key := word[len(word)-n:]
		if _, ok := chain[key]; !ok {
			chain[key] = []string{}
		}
		chain[key] = append(chain[key], EOW)
	}
	var initVals = []string{}
	for k := range chain {
		initVals = append(initVals, k)
	}
	return MarkovChain{chain, initVals}
}

func (chain MarkovChain) String() string {
	s := "{\n"
	for k, v := range chain.m {
		s += fmt.Sprintf("\t%q: %q\n", k, v)
	}
	return s + "}"
}

func (chain MarkovChain) Equals(other MarkovChain) bool {
	if len(chain.m) != len(other.m) {
		return false
	}
	for k, v := range chain.m {
		if len(v) != len(other.m[k]) {
			return false
		}
		sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
		sort.Slice(other.m[k], func(i, j int) bool { return other.m[k][i] < other.m[k][j] })
		for i, c := range v {
			if c != other.m[k][i] {
				return false
			}
		}
	}
	return true
}
