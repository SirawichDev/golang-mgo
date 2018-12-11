package mongodb

import "gopkg.in/mgo.v2"

func Config() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil{
		return nil,err
	}
	return session,nil
}