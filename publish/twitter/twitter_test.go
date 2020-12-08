package twitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkContent(t *testing.T) {
	twitClient := twitterClient{}
	testcases := map[string]struct {
		input  string
		output []string
	}{
		"Less than 140 chars": {
			input:  "One two three",
			output: []string{"One two three"},
		},
		"Two groups split on space": {
			input:  "One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three",
			output: []string{"One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two", "threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three"},
		},
		"Six groups split on space": {
			input: "One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three",
			output: []string{
				"One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two",
				"threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two",
				"threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two",
				"threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two",
				"threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two",
				"threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two",
				"threeOne two threeOne two three"},
		},
		"Two groups split on new line": {
			input:  "One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two\nthreeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three",
			output: []string{"One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two", "threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three"},
		},
		"Two groups split on tab": {
			input:  "One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two\tthreeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three",
			output: []string{"One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two", "threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three"},
		},
		"breaking in prod": {
			input: `He aha koa. Hai te tokorima a Māui.

			It does not matter. I have the five of Māui.
			(If the host apologises for the lack of cutlery available, the guest replies that he has his fingers - the five of Māui.)`,
			output: []string{`He aha koa. Hai te tokorima a Māui.

			It does not matter. I have the five of Māui.
			(If the host apologises for the lack of cutlery`, `available, the guest replies that he has his fingers - the five of Māui.)`},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			output := twitClient.chunkContent(tc.input)
			assert.Equal(t, len(tc.output), len(output), "Outputs have different lengths")
			for i := range tc.output {
				assert.Equalf(t, tc.output[i], output[i], "output strings do not match at position %d", i)
			}
		})
	}
}
