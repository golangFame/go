package fix_2_separated_files

import "go/build/testdata/issue60817/fix_2_separated_files"

/*type T2[T any] struct {
	e *innerT[T, *T2[T]]
}*/

type H2 struct {
	h *H2
}
type T2[T any] struct {
	e *fix_2_separated_files.innerT[T, *T2[T]]
}
