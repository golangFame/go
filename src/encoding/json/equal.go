package json

import (
	"reflect"
)

// Eq Determine whether two JSON strings are equivalent
// (with the same key value, in different order)
func Eq(s1, s2 string) (bool, error) {
	var i, i2 interface{}
	err := Unmarshal([]byte(s1), &i)
	if err != nil {
		return false, err
	}
	err = Unmarshal([]byte(s2), &i2)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(i, i2), nil
}

// Equal Determine whether two or more JSON strings are equivalent
// (with the same key value and different order)
func Equal(s1, s2 string, sN ...string) (bool, error) {
	if len(sN) == 0 {
		return Eq(s1, s2)
	}
	equal, err := Eq(s1, s2)
	if err != nil {
		return false, err
	}
	if !equal {
		return false, nil
	}
	for _, js := range sN {
		eq, err := Eq(s1, js)
		if err != nil {
			return false, err
		}
		if !eq {
			return false, nil
		}
	}
	return true, nil
}
