package main

import "fmt"


/*type c[P any] interface {
	f0[P,*c[P]]
}

func f0[P any, P1 c[P]] (x P1){

	fmt.Println(x)

	return
}*/


func main(){

	// var s = T1[any]
	//
	// _ = T1[any,*s]{}

	// var x X
	// fmt.Println(x)

	// f := f0(nil)
	// fmt.Println(f)
	// var x T1[int]
	// fmt.Println(x)

	f22 := FP(nil)
	fmt.Println(f22)

	fmt.Println("end")
}
