package main

import (
	"github.com/mysqlx/mysqlx"
	"log"
)

//import "github.com/mysqlx/mysqlx"

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

	// New method chaining used for executing an SQL SELECT statement
	// Recommended way for executing queries

	db := mysqlx.Database{}

	employees := db.GetSchema().GetTable("test")

	var rows chan *mysqlx.RowResult = make(chan *mysqlx.RowResult)

	//execute the query asynchronously
	employees.Select("name", "age").
		Where("name like :param").
		OrderBy("name").Bind(map[string]string{"name": "m%"}).
		ExecuteAsync(rows)

	<-rows

}
