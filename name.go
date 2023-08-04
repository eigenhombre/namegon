package main

import (
	"math/rand"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func generateName(chain MarkovChain, n int) string {
	// Randomly choose first ngram from chain:
	startIdx := rand.Intn(len(chain.initVals))
	name := chain.initVals[startIdx]
	// Randomly choose next character from chain,
	// or terminate when EOW is chosen.
	for {
		choices := chain.m[name[len(name)-n:]]
		if len(choices) == 0 {
			break
		}
		idx := rand.Intn(len(choices))
		choice := choices[idx]
		if choice == EOW {
			break
		}
		name += choices[idx]
	}
	return cases.Title(language.AmericanEnglish).String(name)
}

func Namer(names []string, n int) func() string {
	ch := buildChain(names, n)
	return func() string {
		return generateName(ch, n)
	}
}
