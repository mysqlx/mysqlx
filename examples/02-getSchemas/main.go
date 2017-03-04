package main

import (
	"github.com/mysqlx/mysqlx"
	"log"
)

func main()  {

	// Connecting to MySQL and working with a Session
	mySession,err := mysqlx.GetSession(mysqlx.SessionSettings{
		Host:"127.0.0.1",
		Port:33060,
		DbUser:"wsc",
		DbPassword:"ooxx",
	})

	if err!=nil {
		log.Fatal(err)
	}

	//close session
	defer mySession.Close()

	//Get a list of all available schemas

	schemaList := mySession.GetSchemas()

	// Loop over all available schemas and print their name
	for _,schema := range schemaList {
		println(&schema)
	}





}
