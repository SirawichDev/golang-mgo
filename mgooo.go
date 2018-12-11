package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

type myself struct {
	Name string `bson:"name"`
	Age int `bson:"age"`
	Money int `bson:"money"`
	Job string `bson:"job"`
}

func main()  {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer  session.Close()

	myselfCollection := session.DB("myself").C("")
}
