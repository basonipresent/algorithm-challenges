package main

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{
			name:      "example from doc comment",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		{
			name:      "endWord not in wordList, no valid sequence",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      0,
		},
		{
			name:      "direct one-letter transformation",
			beginWord: "hot",
			endWord:   "dot",
			wordList:  []string{"hot", "dot"},
			want:      2,
		},
		{
			name:      "transformation requires the letter z",
			beginWord: "a",
			endWord:   "z",
			wordList:  []string{"z"},
			want:      2,
		},
		{
			name:      "endWord in wordList but unreachable, no bridging word",
			beginWord: "hot",
			endWord:   "dog",
			wordList:  []string{"hot", "dog"},
			want:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			if got != tt.want {
				t.Errorf("ladderLength(%q, %q, %v) = %d, want %d", tt.beginWord, tt.endWord, tt.wordList, got, tt.want)
			}
		})
	}
}
