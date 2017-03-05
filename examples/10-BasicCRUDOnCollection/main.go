package main

import (
	"github.com/mysqlx/mysqlx"
	"log"
)

func main() {
	mySession, err := mysqlx.GetNodeSession(mysqlx.SessionSettings{
		Host:       "127.0.0.1",
		Port:       33060,
		DbUser:     "wsc",
		DbPassword: "ooxx",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer mySession.Close()

	myDb := mySession.GetSchema("test")

	// Create a new collection 'my_collection'

	myColl := myDb.CreateCollection("my_collection")

	// Insert documents
	myColl.Add("{\"name\":\"Sakila\", \"age\":15}").Execute()
	myColl.Add("{\"name\":\"Susanne\", \"age\":24}").Execute()
	myColl.Add("{\"name\":\"Mike\", \"age\":39}").Execute()

}
