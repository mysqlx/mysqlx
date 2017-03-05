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

	myColl := myDb.GetCollection("test")

	res := myColl.Add("{\"_id\":\"custom_id\",\"a\":1}").Execute()

	println(res.GetLastInsertId())

}
