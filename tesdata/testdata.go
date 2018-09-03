package testdata

// Thing holds other things
type Thing struct {
	F1 string `json:"x"`
	F2 []int
	F3 map[byte]float64
	F4 bool     `json:"x_y"`
	F5 bool     `json:"xY"`
	F6 bool     `json:"foo bar"`
	F7 *int     `json:"TitleCase"`
	F8 []string `json:"a b,omitempty"`
}

func hasInlineStruct() {
	type inline struct {
		Z bool `json:"Inline Struct"`
	}
}
