package main

import (
	"github.com/mysqlx/mysqlx"
	"log"
)

func main()  {
	mySession,err := mysqlx.GetSession(mysqlx.SessionSettings{
		Host:"127.0.0.1",
		Port:33060,
		DbUser:"wsc",
		DbPassword:"ooxx",
	})

	if err!=nil {
		log.Fatal(err)
	}

	defer mySession.Close()

	myDb := mySession.GetSchema("test")

	myColl := myDb.GetCollection("my_collection")

	myDocs := myColl.
		Find("name like :param").
		Limit(1).
		Bind("param","S%").
		Execute()


	println(myDocs.FetchOne())

}
