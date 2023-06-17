package failing_test

import "go/build/testdata/issue60817/failing_test"

type T2[T any] struct {
	e *failing_test.innerT[T, *T2[T]]
}
