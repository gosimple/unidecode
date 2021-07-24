package unidecode

import (
	"testing"
)

func TestUnidecode(t *testing.T) {
	tests := map[string]struct {
		s    string
		want string
	}{
		"ASCII": {
			"ABCDEF",
			"ABCDEF",
		},
		"Knosos": {
			"Κνωσός",
			"Knosos",
		},
		"BeiJing": {
			"\u5317\u4EB0",
			"Bei Jing ",
		},
		"Emoji": {
			"Hey Luna t belle 😵😂",
			"Hey Luna t belle ",
		},
		"U+10000 plain string": {
			"𐀀",
			"",
		},
		"U+10000 ASCII string": {
			"\U00010000",
			"",
		},
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Unidecode(tt.s); got != tt.want {
				t.Errorf("Unidecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkUnidecode(b *testing.B) {
	cases := []string{
		"ABCDEF",
		"Κνωσός",
		"\u5317\u4EB0",
	}
	for ii := 0; ii < b.N; ii++ {
		for _, v := range cases {
			_ = Unidecode(v)
		}
	}
}

func BenchmarkDecodeTable(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		decodeTransliterations()
	}
}

func init() {
	decodeTransliterations()
}
