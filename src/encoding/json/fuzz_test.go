// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"bytes"
	"io"
	"testing"
)

func FuzzUnmarshalJSON(f *testing.F) {
	f.Add([]byte(`{
"object": {
	"slice": [
		1,
		2.0,
		"3",
		[4],
		{5: {}}
	]
},
"slice": [[]],
"string": ":)",
"int": 1e5,
"float": 3e-9"
}`))

	f.Fuzz(func(t *testing.T, b []byte) {
		for _, typ := range []func() any{
			func() any { return new(any) },
			func() any { return new(map[string]any) },
			func() any { return new([]any) },
		} {
			i := typ()
			if err := Unmarshal(b, i); err != nil {
				return
			}

			encoded, err := Marshal(i)
			if err != nil {
				t.Fatalf("failed to marshal: %s", err)
			}

			if err := Unmarshal(encoded, i); err != nil {
				t.Fatalf("failed to roundtrip: %s", err)
			}
		}
	})
}

func FuzzDecoderToken(f *testing.F) {
	f.Add([]byte(`{
"object": {
	"slice": [
		1,
		2.0,
		"3",
		[4],
		{5: {}}
	]
},
"slice": [[]],
"string": ":)",
"int": 1e5,
"float": 3e-9"
}`))

	f.Fuzz(func(t *testing.T, b []byte) {
		r := bytes.NewReader(b)
		d := NewDecoder(r)
		for {
			_, err := d.Token()
			if err != nil {
				if err == io.EOF {
					break
				}
				return
			}
		}
	})
}

func FuzzEqual(f *testing.F) {

	s1 := `{"name":"Hiro","email":"laciferin@gmail.com","age":19}`
	s2 := `{"name":"Hiro","age":19,"email":"laciferin@gmail.com"}`

	f.Add([]byte(s1))
	f.Add([]byte(s2))

	var validJSON = func(sN ...[]byte) (err error) {

		ptr := new(struct{})

		for _, s := range sN {
			if err = Unmarshal(s, ptr); err != nil {
				return
			}
		}
		return
	}

	f.Fuzz(func(t *testing.T, s1 []byte) {

		if err := validJSON(s1); err != nil {
			return
		}

		tt := testType[[]byte]{
			s1: s1,
			s2: s1,
			sN: [][]byte{
				s1,
			},
			want: true,
		}

		testEqual(t, tt)
	})
}
