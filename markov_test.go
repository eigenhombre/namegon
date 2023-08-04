package namegon

import "testing"

func TestChain(t *testing.T) {
	input := func(a ...string) []string { return a }
	strings := input // input is a slice of strings, output keys are as well
	yields := func(ss ...any) MarkovChain {
		m := make(map[string][]string)
		for i := 0; i < len(ss)-1; i += 2 {
			key := ss[i].(string)
			val := ss[i+1].([]string)
			m[key] = val
		}
		return MarkovChain{m: m}
	}
	testCases := []struct {
		words  []string
		output MarkovChain
	}{
		{input("aa"), yields("aa", strings(EOW))},
		{input("aaa", "aab"), yields("aa", strings("a", "b", EOW), "ab", strings(EOW))},
		{input("aaa", "aa"), yields("aa", strings("a", EOW, EOW))},
		{input("aaa", "aab", "aac"), yields("aa", strings("a", "b", "c", EOW), "ab", strings(EOW), "ac", strings(EOW))},
	}
	for _, tc := range testCases {
		chain := buildChain(tc.words, 2)
		if !chain.Equals(tc.output) {
			t.Errorf("Expected %s, got %s", tc.output, chain)
		}
	}
}
