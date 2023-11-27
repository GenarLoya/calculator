package main

import (
	"fmt"
	. "genarold/calculator/src/httpserver"
)

func main() {

	srv := New(":8080")

	err := srv.ListenAndServe()

	if err != nil {
		fmt.Println("Errowferbervr: ", err)
		panic(err)
	}

}
