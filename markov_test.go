package main

import "testing"

func TestChain(t *testing.T) {
	Ss := func(a ...string) []string { return a }
	testCases := []struct {
		words  []string
		output MarkovChain
	}{
		{Ss("aa"), MarkovChain{
			"aa": Ss(EOW),
		}},
		{Ss("aaa", "aab"), MarkovChain{
			"aa": Ss("a", "b", EOW),
			"ab": Ss(EOW),
		}},
		{Ss("aaa", "aa"), MarkovChain{
			"aa": Ss("a", EOW, EOW),
		}},
		{Ss("aaa", "aab", "aac"), MarkovChain{
			"aa": Ss("a", "b", "c", EOW),
			"ab": Ss(EOW),
			"ac": Ss(EOW),
		}},
	}
	for _, tc := range testCases {
		chain := buildChain(tc.words, 2)
		if !chain.Equals(tc.output) {
			t.Errorf("Expected %s, got %s", tc.output, chain)
		}
	}
}

// func TestChain(t *testing.T) {
// 	Ss := func(a ...string) []string { return a }
// 	testCases := []struct {
// 		words  []string
// 		output map[string][]string
// 	}{
// 		{Ss("aa"), map[string][]string{
// 			"aa": Ss(EOW),
// 		}},
// 		{Ss("aaa", "aab"), map[string][]string{
// 			"aa": Ss("a", "b", EOW),
// 			"ab": Ss(EOW),
// 		}},
// 		// {[]string{"aaa", "aa"}, map[string][]string{
// 		// 	"aa": []string{"a", EOW},
// 		// }},
// 	}
// 	for _, tc := range testCases {
// 		chain := buildChain(tc.words, 2)
// 		if chain != tc.output {
// 			t.Errorf("Expected %q, got %q", tc.output, chain)
// 		}
// 		// if len(chain) != len(tc.output) {
// 		// 	t.Error("Expected", len(tc.output), "chains, got", len(chain))
// 		// }
// 		for k, v := range chain {

// 			// if len(v) != len(tc.output[k]) {
// 			// 	t.Errorf("Expected %d values for %s, got %d: %q", len(tc.output[k]), k, len(v), v)

// 			// 	continue
// 			// }
// 			// for i, c := range v {
// 			// 	if c != tc.output[k][i] {
// 			// 		t.Errorf("Expected '%s' for %s[%d], got '%s'", tc.output[k][i], k, i, c)
// 			// 	}
// 			// }
// 		}
// 	}
// }
