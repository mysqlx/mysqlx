package main

import (
	"github.com/mysqlx/mysqlx"
	"log"
)

func main()  {
	//connect to server on localhost
	mySession,err := mysqlx.GetNodeSession(mysqlx.SessionSettings{
		Host:"127.0.0.1",
		Port:33060,
		DbUser:"wsc",
		DbPassword:"ooxx",
	})

	if err!=nil {
		log.Fatal(err)
	}

	defer mySession.Close()

	mySession.SetCurrentSchema("test")
}
