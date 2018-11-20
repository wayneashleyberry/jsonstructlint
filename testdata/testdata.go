package testdata

// NoTags has no json tags
type NoTags struct {
	F1 string
	F2 string
}

// Thing holds other things
type Thing struct {
	F1  string `json:"x"`
	F2  []int
	F3  map[byte]float64
	F4  bool        `json:"x_y"`
	F5  bool        `json:"xY"`
	F6  bool        `json:"foo bar"`
	F7  *int        `json:"TitleCase"`
	F8  []string    `json:"a b,omitempty"`
	F9  interface{} `json:"ignore me"`      // nolint: jsonstructlint
	F10 interface{} `json:"also ignore me"` // nolint: foo,jsonstructlint
}

func hasInlineStruct() {
	type inline struct {
		Z bool `json:"Inline Struct"`
	}

	fn := func() {
		type foobar struct {
			FX string `json:"Super Inline"`
		}
	}

	fn()
}

func nested() {
	type RequestBody struct {
		File struct {
			Name             string `json:"FileName"`
			MissingStructTag string
		} `json:"file"`
	}
}
