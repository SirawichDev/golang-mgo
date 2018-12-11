package main

import "gomongo/mongodb"

func main() {
	if err := mongodb.Runner(); err != nil {
		panic(err)
	}
}