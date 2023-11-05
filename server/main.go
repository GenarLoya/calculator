package main

import (
	. "genarold/calculator/src/httpserver"
)

func main() {

	srv := New(":8080")

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
