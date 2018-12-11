package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type MySelf struct {
	Name string `bson:"name"`
	Money float32 `bson:"money"`
	Age int `bson:"age"`
	Job string `bson:"job"`
}

func Runner() error{
	db, err := Config()
	if err != nil{
		return err
	}
	connect := db.DB("admin").C("gomongo")
	if err := connect.Insert(&MySelf{"sirawich", 3000.0,23,"Developer"},
	&MySelf{"Exe", 1000.0,20,"Programmer"}, &MySelf{"kingkong", 3000.0,21,"Killer"}); err != nil {
		return err
	}
	var agent []MySelf
	if err := connect.Find(bson.M{"money": 3000.0}).All(&agent); err != nil {
		return err
	}
	//if err := connect.DropCollection(); err != nil {
	//	return err
	//}
	fmt.Printf("State: %v ", agent)
	return nil
}